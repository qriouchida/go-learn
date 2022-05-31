package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	engine *gorm.DB
}

func NewService(engine *gorm.DB) Servicer {
	return &Service{engine}
}

type Servicer interface {
	NewUser() *Users
}

func (s *Service) NewUser() *Users {
	return NewUsers(s.engine)
}

const (
	ServiceKey = "service_factory"
)

func ServiceFactoryMiddleware(svc Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ServiceKey, svc)
		c.Next()
	}
}
