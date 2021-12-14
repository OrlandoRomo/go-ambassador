package controller

type AppController struct {
	Auth    AuthController
	User    UserController
	Product ProductController
}
