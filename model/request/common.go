package request

type Page struct {
	PageNum  int `json:"page_num" binding:"required,gt=0"`
	PageSize int `json:"page_size" binding:"required,gt=0,lt=51"`
}
