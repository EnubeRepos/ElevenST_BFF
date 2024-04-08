package main

import (
	"fmt"

	"github.com/EnubeRepos/ElevenST_BFF/config"
	"github.com/EnubeRepos/ElevenST_BFF/web"
	"github.com/EnubeRepos/ElevenST_BFF/web/routes"
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
	routes.Set(cfg)
}

// func (App *App) RetrieveContacts(ctx *gin.Context) {
// 	var contactList model.ResponseContactList
// 	var rsp []model.ResponseContact
// 	var wg sync.WaitGroup
// 	wg.Add(len(App.Config.Clinics))

// 	for _, clinic := range App.Config.Clinics {
// 		go func(clinic model.ClinicReference) {
// 			defer wg.Done()

// 			list, err := App.Web.GetContactByDocumentNumber(clinic.URL, ctx.Param("cpf"))
// 			if err != nil {
// 				App.Config.Log.Error(err.Error())
// 				return
// 			}

// 			if list.Total == 0 {
// 				App.Config.Log.Info(fmt.Sprintf("Clinic: %s | Contact not found with doc: %s", clinic.Name, ctx.Param("cpf")))
// 				return
// 			}

// 			rsp = append(rsp, model.ResponseContact{
// 				ClinicName: clinic.Name,
// 				URL:        clinic.URL,
// 				Entity:     "Contact",
// 				Content:    list.List[0],
// 			})
// 		}(clinic)
// 	}

// 	wg.Wait()

// 	contactList.Total = len(rsp)
// 	contactList.List = rsp

// 	ctx.JSON(http.StatusOK, contactList)
// }
