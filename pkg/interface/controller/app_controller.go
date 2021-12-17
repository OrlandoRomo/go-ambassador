package controller

type AppController struct {
	Auth       AuthController
	User       UserController
	Product    ProductController
	Ambassador AmbassadorController
	Link       LinkController
	Ranking    RankingController
	Order      OrderController
}
