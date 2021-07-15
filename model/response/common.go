package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}
