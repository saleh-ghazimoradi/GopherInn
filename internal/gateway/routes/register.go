package routes

import "github.com/gin-gonic/gin"

type Register struct {
	HealthRoutes *HealthRoutes
	UserRoutes   *UserRoutes
}

type Options func(*Register)

func WithHealthRoutes(healthRoutes *HealthRoutes) Options {
	return func(r *Register) {
		r.HealthRoutes = healthRoutes
	}
}

func WithUserRoutes(userRoutes *UserRoutes) Options {
	return func(r *Register) {
		r.UserRoutes = userRoutes
	}
}

func (r *Register) RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	r.HealthRoutes.HealthRoute(router)
	r.UserRoutes.UserRoute(router)

	return router
}

func NewRegister(opts ...Options) *Register {
	register := &Register{}
	for _, opt := range opts {
		opt(register)
	}
	return register
}
