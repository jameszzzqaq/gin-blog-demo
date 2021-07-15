package request

type AriticleAdd struct {
	TagID int `json:"tag_id" binding:"required"`

	Title     string `json:"title" binding:"required,max=30"`
	Desc      string `json:"desc" binding:"required,max=150"`
	Content   string `json:"content" binding:"required,max=180"`
	CreatedBy string `json:"created_by" binding:"required"`
	State     *int   `json:"state" binding:"required,oneof=0 1"`
}

type ArticleUpdate struct {
	TagID int `json:"tag_id" binding:"required"`

	Title      string `json:"title" binding:"required,max=30"`
	Desc       string `json:"desc" binding:"required,max=150"`
	Content    string `json:"content" binding:"required,max=180"`
	ModifiedBy string `json:"modified_by" binding:"required"`
	State      *int   `json:"state" binding:"required,oneof=0 1"`
}

type ArticleListGet struct {
	Page
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	TagID   *int   `json:"tag_id"`
}
