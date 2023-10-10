package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/EnubeRepos/ElevenST_BFF/config"
	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/EnubeRepos/ElevenST_BFF/web"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config *config.Config
	Web    *web.Web
}

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println("error initing config")
		return
	}

	var AppConfig = App{
		Config: cfg,
		Web:    web.New(cfg.Log, cfg.Authorization),
	}

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.GET("/contact/:cpf", AppConfig.RetrieveContacts)

	err = r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		cfg.Log.Error(err.Error())
	}
}

func (App *App) RetrieveContacts(ctx *gin.Context) {
	var contactList model.ResponseContactList
	var rsp []model.ResponseContact
	var wg sync.WaitGroup
	wg.Add(len(App.Config.Clinics))

	for _, clinic := range App.Config.Clinics {
		go func(clinic model.ClinicReference) {
			defer wg.Done()

			list, err := App.Web.GetContactByDocumentNumber(clinic.URL, ctx.Param("cpf"))
			if err != nil {
				App.Config.Log.Error(err.Error())
				return
			}

			if list.Total == 0 {
				App.Config.Log.Info(fmt.Sprintf("Clinic: %s | Contact not found with doc: %s", clinic.Name, ctx.Param("cpf")))
				return
			}

			rsp = append(rsp, model.ResponseContact{
				ClinicName: clinic.Name,
				URL:        clinic.URL,
				Entity:     "Contact",
				Content:    list.List[0],
			})
		}(clinic)
	}

	wg.Wait()

	contactList.Total = len(rsp)
	contactList.List = rsp

	ctx.JSON(http.StatusOK, contactList)
}
