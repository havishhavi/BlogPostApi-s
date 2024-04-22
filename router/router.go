package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"www.blog.com/controller"
	"www.blog.com/middleware"
)

// what the code in it??
func SetUpRouter() *gin.Engine {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("POST", "GET", "PUT", "DELETE", "UPDATE", "OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))
	/**Allow origin CORS setting end:**/

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	/*-------------Routeing started---------------*/
	user := r.Group("api/blog/user")
	{
		var UserController = new(controller.UserController)
		user.POST("/register", UserController.Register)
		user.POST("/login", UserController.Login)

	}
	//how you login user and authenticate : JWT authenntication using user id using signingmethodhs256  inmiddleware
	// how to handle if token is expired : relogin or by using refresh token
	//normally tokens life is for specified time but refrest token can be used to give time to login , we send two tokens
	// when we directly create a post controller  without a middle ware we will be
	// post := r.Group("api/blog/post")
	// {
	// 	var PostController = new(controller.PostController)
	// 	post.POST("/create_post", PostController.CreatePost)

	// }
	//middleware???  why defination to authorixe jwt?? to check if the jwt token the user sending is correct or not
	post := r.Group("api/blog/post")
	//middleware calling
	post.Use(middleware.AuthorizeJWT())
	{
		var PostController = new(controller.PostController)
		post.POST("/create_post", PostController.CreatePost)
		//sir why dint we use path param to get the value?
		post.GET("/view_post/:Postid", PostController.ViewPost)
	}
	return r
}

//shouldBindJson convert json type file to a struct
