package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hekimapro/utils/helpers"
	"github.com/hekimapro/utils/log"
)

func SetSession(response http.ResponseWriter, userID uuid.UUID) {
	cookie := &http.Cookie{
		Name:     "token",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Value:    userID.String(),
		MaxAge:   86400,
	}
	http.SetCookie(response, cookie)
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "token",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetToken(request *http.Request) (*uuid.UUID, error) {
	cookie, err := request.Cookie("token")
	if err != nil {
		log.Error(err.Error())
		return nil, helpers.CreateError("failed to get token")
	}

	// convert string to uuid
	parsedUUID, err := uuid.Parse(cookie.Value)
	if err != nil {
		log.Error(err.Error())
		return nil, helpers.CreateError("failed to parse token to uuid")
	}

	return &parsedUUID, nil
}
