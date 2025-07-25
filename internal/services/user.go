package services

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/hekimapro/task-vault/internal/models"
	"github.com/hekimapro/task-vault/internal/repositories"
	"github.com/hekimapro/utils/encryption"
	"github.com/hekimapro/utils/helpers"
)

// service blue print
type User interface {
	Create(user *models.User) (*uuid.UUID, error)
	Authenticate(emailAddress, password string) (*models.User, error)
}

type service struct {
	Repo repositories.User
}

func NewUser(repo repositories.User) User {
	return &service{
		Repo: repo,
	}
}

func (service *service) Create(user *models.User) (*uuid.UUID, error) {

	// name validation
	nameLength := len(user.Name)
	if user.Name == "" {
		return nil, helpers.CreateError("name is required")
	} else if nameLength < 3 || nameLength > 100 {
		return nil, helpers.CreateError("name must have characters between 3 and 100 inclusive")
	}

	// email validation
	if user.Email == "" {
		return nil, helpers.CreateError("email address is required")
	} else if len(user.Email) > 256 {
		return nil, helpers.CreateError("email address must have characters less or equal to 256")
	} else {
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			return nil, helpers.CreateError("invalid email address")
		}
	}

	// password validation
	if user.Password == "" {
		return nil, helpers.CreateError("password is required")
	} else if len(user.Password) < 8 {
		return nil, helpers.CreateError("password must have atleast 8 characters")
	}

	// hashing user password
	hashedPassword, err := encryption.CreateHash(user.Password)
	if err != nil {
		return nil, helpers.CreateError("failed to hash user password")
	}

	// set user password with hashed password
	user.Password = hashedPassword

	// create user
	return service.Repo.Create(user)
}

func (service service) Authenticate(emailAddress, password string) (*models.User, error) {

	// validate email address
	if emailAddress == "" {
		return nil, helpers.CreateError("email address is required")
	} else if len(emailAddress) > 256 {
		return nil, helpers.CreateError("email address must characters less or equal to 256")
	} else {
		_, err := mail.ParseAddress(emailAddress)
		if err != nil {
			return nil, helpers.CreateError("invalid email address")
		}
	}

	// validate password
	if password == "" {
		return nil, helpers.CreateError("password is required")
	} else if len(password) < 8 {
		return nil, helpers.CreateError("password must have atleast 8 characters")
	}

	// get user
	user, err := service.Repo.FindByEmail(emailAddress)
	if err != nil {
		return nil, err
	}

	// compare provided plain password with user hashed password
	passwordMatched := encryption.CompareWithHash(user.Password, password)
	if !passwordMatched {
		return nil, helpers.CreateError("password is not correct")
	}

	// hide user hashed password
	user.Password = ""

	return user, nil
}
