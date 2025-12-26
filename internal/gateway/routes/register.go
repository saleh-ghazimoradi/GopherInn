package routes

import "github.com/gin-gonic/gin"

type Register struct {
	HealthRoutes *HealthRoutes
}

type Options func(*Register)

func WithHealthRoutes(healthRoutes *HealthRoutes) Options {
	return func(r *Register) {
		r.HealthRoutes = healthRoutes
	}
}

func (r *Register) RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	r.HealthRoutes.HealthRoute(router)

	return router
}

func NewRegister(opts ...Options) *Register {
	register := &Register{}
	for _, opt := range opts {
		opt(register)
	}
	return register
}
