package models

type Pagination struct {
	PageNum  int `form:"pagenum"`
	PageSize int `form:"pagesize"`
}
