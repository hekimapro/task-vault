package handlers

import (
	"net/http"

	"github.com/hekimapro/task-vault/internal/middleware"
	"github.com/hekimapro/task-vault/internal/models"
	"github.com/hekimapro/task-vault/internal/services"
	views "github.com/hekimapro/task-vault/web/templates/views/user"
)

type handler struct {
	Service services.User
}

func NewUser(service services.User) *handler {
	return &handler{
		Service: service,
	}
}

func (handler *handler) LoginView(response http.ResponseWriter, request *http.Request) {
	views.Login().Render(request.Context(), response)
}

func (handler *handler) Create(response http.ResponseWriter, request *http.Request) {
	var user models.User
	user.Name = request.FormValue("name")
	user.Email = request.FormValue("email")
	user.Password = request.FormValue("password")
	user.PasswordConfirmation = request.FormValue("password_confirmation")

	userID, err := handler.Service.Create(&user)
	if err != nil {

	}

	middleware.SetSession(response, *userID)
	http.Redirect(response, request, "/dashboard", http.StatusSeeOther)
}

func (handler *handler) Login(response http.ResponseWriter, request *http.Request) {
	password := request.FormValue("password")
	emailAddress := request.FormValue("email")
	user, err := handler.Service.Authenticate(emailAddress, password)
	if err != nil {
		views.Login().Render(request.Context(), response)
		return 
	}
	middleware.SetSession(response, user.ID)
	http.Redirect(response, request, "/dashboard", http.StatusSeeOther)
}
