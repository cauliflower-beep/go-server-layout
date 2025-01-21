package setting

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf = new(SvrConfig)

const (
	ConfigEnv         = "WEB_APP_CONFIG"             // 环境变量
	ConfigDefaultFile = "./conf/config.yaml"         // 默认配置文件
	ConfigTestFile    = "./conf/config.test.yaml"    // 测试环境默认配置文件
	ConfigDebugFile   = "./conf/config.debug.yaml"   // 开发环境默认配置文件
	ConfigReleaseFile = "./conf/config.release.yaml" // 线上默认配置文件
)

type SvrConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Addr      string `mapstructure:"addr"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*MongoConfig `mapstructure:"mongodb"`
}

type MongoConfig struct {
	Uri           string     `mapstructure:"uri"`
	ConnectTimout int        `mapstructure:"connect-timeout"` // 连接超时时间 s
	Database      string     `mapstructure:"database"`
	Credential    Credential `mapstructure:"credential"`
}

type Credential struct {
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	AuthMechanism string `mapstructure:"auth-mechanism"`
	AuthSource    string `mapstructure:"auth-source"`
	PasswordSet   bool   `mapstructure:"password-set"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	ParseTime    bool   `mapstructure:"parse_time"`
	Charset      string `mapstructure:"charset"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	TablePrefix  string `mapstructure:"table_prefix"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	Filepath   string `mapstructure:"filepath"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init(filePath ...string) (err error) {
	var path string
	if len(filePath) == 0 {
		// 启动时可添加命令行参数，指定配置文件 -c xxxx.yaml
		flag.StringVar(&path, "c", "", "choose config file.")
		flag.Parse()
		if path == "" {
			// 优先级：命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				path = ConfigDefaultFile
				fmt.Printf("get config path by default.path is %s\n", path)
			} else {
				path = configEnv
				fmt.Println("get config path by env. path is ", path)
			}
		} else {
			fmt.Println("get config path by args. path is ", path)
		}
	} else {
		path = filePath[0]
	}

	viper.SetConfigFile(path)

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 conf 变量中
	if err = viper.Unmarshal(conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	// 配置热重载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err = viper.Unmarshal(conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
		fmt.Printf("config is changed|%+v\n", *conf)
	})
	return
}

func GetConf() *SvrConfig {
	return conf
}
