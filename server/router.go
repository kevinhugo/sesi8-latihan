package server

import (
	_ "sesi8-latihan/docs"
	"sesi8-latihan/server/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	people *controllers.PersonController
}

func NewRouter(people *controllers.PersonController) *Router {
	return &Router{people: people}
}

func (r *Router) Start(port string) {
	router := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/people", r.people.GetPeople)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(port)
}
