package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"ksbeasle.net/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (um *UserModel) Insert(name string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
			 VALUES (?, ?, ?, UTC_TIMESTAMP())`

	_, err = um.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var mysqlerr *mysql.MySQLError
		if errors.As(err, &mysqlerr) {
			if mysqlerr.Number == 1062 && strings.Contains(mysqlerr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil

}

func (um *UserModel) AuthenticateUser(email string, password string) (int, error) {
	return 0, nil
}

func (um *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
