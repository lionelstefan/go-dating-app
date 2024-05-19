package repository

import (
	"database/sql"

	"github.com/lionelstefan/go-dating-app/model"
	"github.com/lionelstefan/go-dating-app/shared"
	"golang.org/x/crypto/bcrypt"
)

func FindByCredentials(email string) (user model.User, err error) {
	db := shared.CreateDBConnection()

	err = db.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return
	}

	return
}

func RegisterNewUser(user model.UserRegister) (sql.Result, error) {
	db := shared.CreateDBConnection()

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return nil, err
	}

	res := db.MustExec("INSERT INTO users (email, password, full_name, age, gender, is_verified) VALUES ($1, $2, $3, $4, $5)", user.Email, hashed, user.FullName, user.Age, user.Gender)

	return res, nil
}