package dto

type CreatePost struct {
	Title string `binding:"required,min=5,max=100"`
	Post  string `binding:"required,min=5,max=5000"`
}
