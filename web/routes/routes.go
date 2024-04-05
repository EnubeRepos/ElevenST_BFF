package routes

import (
	"github.com/EnubeRepos/ElevenST_BFF/config"
	"github.com/EnubeRepos/ElevenST_BFF/web/controller"
	"github.com/gin-gonic/gin"
)

func Set(cfg *config.Config) {

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	api := r.Group("/api/v1/")
	api.POST("pix/generate", func(ctx *gin.Context) { controller.GeneratePixController(ctx, cfg) })
	api.GET("contact/cpf/:cpf", func(ctx *gin.Context) { controller.GetContactByCPF(ctx, cfg) })

	// r.GET("/contact/:cpf", AppConfig.RetrieveContacts)
	err := r.Run(":8090")
	if err != nil {
		cfg.Log.Error(err.Error())
	}
}
