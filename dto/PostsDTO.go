package dto

type CreatePost struct {
	Title string `binding:"required,min=5,max=100"`
	Post  string `binding:"required,min=5,max=5000"`
}

type ViewPost struct {
	Postid uint `uri:"Postid" binding:"required"`
}

type EditPost struct {
	Id    uint   `binding:"required,numeric"`
	Title string `binding:"required,min=5,max=100"`
	Post  string `binding:"required,min=5,max=5000"`
}
