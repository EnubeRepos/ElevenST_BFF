package controller

import (
	"fmt"
	"net/http"

	"github.com/EnubeRepos/ElevenST_BFF/config"
	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
	"github.com/EnubeRepos/ElevenST_BFF/model/api"
	"github.com/EnubeRepos/ElevenST_BFF/web/service"
	"github.com/gin-gonic/gin"
)

func GeneratePixController(c *gin.Context, cfg *config.Config) {
	auth := c.GetHeader("Authorization")
	if auth != fmt.Sprintf("Basic %s", cfg.APIAuth) {
		err := apiError.NewUnauthorizedError("This Authorization don't have access to use this resource")
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	req := new(api.PixGenerate)
	err := req.Bind(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = req.ValidateBody()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := service.GeneratePix(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
