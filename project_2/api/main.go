package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/TomChv/csc-847/project_2/api/pictures"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins:  false,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Referer", "User-Agent"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowOrigins: []string{"http://localhost:3000", "https://cdc-847-project-2.ue.r.appspot.com"},
	}

	fmt.Printf("Cors config: %#v\n", corsConfig)

	r.Use(cors.New(corsConfig))

	picEndpoints := r.Group("/pictures")
	{
		picEndpoints.GET("", pictures.List)
		picEndpoints.POST("", pictures.Add)
		picEndpoints.PUT("/:name", pictures.Update)
		picEndpoints.DELETE("/:name", pictures.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.Run()
}
