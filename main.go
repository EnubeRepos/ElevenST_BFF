package main

import (
	"fmt"
	"sync"

	"github.com/EnubeRepos/ElevenST_BFF/config"
	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/EnubeRepos/ElevenST_BFF/web"
	"github.com/gin-gonic/gin"
)

func main() {
	var cfg, err = config.New()
	if err != nil {
		fmt.Println("error initing config")
		return
	}

	w := web.New(cfg.Log, cfg.Authorization)

	var wg sync.WaitGroup
	wg.Add(len(cfg.Clinics))

	for _, clinic := range cfg.Clinics {
		go func(clinic model.ClinicReference) {
			defer wg.Done()

		}(clinic)
	}
}

func Serve() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.GET("/")

	err := r.Run(":9909")
	if err != nil {
		fmt.Println(err.Error())
	}
}
