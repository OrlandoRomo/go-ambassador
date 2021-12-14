package registry

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Register interface {
	NewAppController() controller.AppController
}

func Newegister(db *gorm.DB, redis *redis.Client) Register {
	return &register{db, redis}
}

type register struct {
	db    *gorm.DB
	redis *redis.Client
}

func (r *register) NewAppController() controller.AppController {
	return controller.AppController{
		Auth:    r.NewAuthController(),
		User:    r.NewUserController(),
		Product: r.NewProductController(),
	}
}
