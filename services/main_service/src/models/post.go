package models

type Post struct {
	PostTitle string `json:"post_title" bindings:"required,max:512"`
	PostText  string `json:"post_text" bindings:"required"`
}

type Pagination struct {
	Offset uint64 `form:"offset"`
	Limit  uint64 `form:"limit"`
}

type PostId struct {
	PostId string `uri:"post_id" binding:"required"`
}

type UserId struct {
	PostId string `uri:"user_id" binding:"required"`
}
