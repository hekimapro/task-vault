package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/hekimapro/task-vault/internal/models"
	"github.com/hekimapro/task-vault/internal/queries"
	"github.com/hekimapro/utils/database"
	"github.com/hekimapro/utils/helpers"
	"github.com/hekimapro/utils/log"
)

// user repository blue print
type User interface {
	Create(user *models.User) (*uuid.UUID, error)
	FindByEmail(emailAddress string) (*models.User, error)
}

type repo struct {
	Database *sql.DB
}

func NewUser(db *sql.DB) User {
	return &repo{
		Database: db,
	}
}

// create user
func (repo *repo) Create(user *models.User) (*uuid.UUID, error) {

	err := database.Transaction(repo.Database, func(transaction *sql.Tx) error {

		err := transaction.QueryRow(queries.UserCreation, user.Name, user.Email, user.Password).Scan(&user.ID)
		if err != nil {
			log.Error(err.Error())
			return helpers.CreateError("failed to create user")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}

// find user by email address
func (repo *repo) FindByEmail(emailAddress string) (*models.User, error) {
	var user models.User

	err := repo.Database.QueryRow(queries.ReadUserByEmail, emailAddress).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, helpers.CreateError("no user has been found")
	}

	return &user, nil
}
