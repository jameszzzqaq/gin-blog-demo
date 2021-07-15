package request

type TagAdd struct {
	Name      string `json:"name" binding:"required,max=100"`
	State     *int   `json:"state" binding:"required,oneof=0 1"`
	CreatedBy string `json:"created_by" binding:"required,max=100"`
}

type TagUpdate struct {
	Name       string `json:"name" binding:"required,max=100"`
	State      *int   `json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `json:"modified_by" binding:"required,max=100"`
}

type TagListGet struct {
	Page
	Name string `json:"name" binding:"max=10"`
}
