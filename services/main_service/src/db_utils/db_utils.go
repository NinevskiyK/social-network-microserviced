package db_utils

import (
	"database/sql"
	"fmt"
	"log"
	"main_service/models"
	"os"

	uuid "github.com/google/uuid"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func StartUpDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	DB = db
	if err != nil {
		return err
	}
	return nil
}

func IsLoginAlreadyInUse(login string) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM users_creds WHERE login = $1", login).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddNewUser(creds models.Credentionals) {
	_, err := DB.Exec("INSERT INTO users_creds (id, login, password) VALUES ($1, $2, $3)", uuid.NewString(), creds.UserName, creds.UserPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPassword(creds models.Credentionals) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users_creds WHERE login = $1 AND password = $2", creds.UserName, creds.UserPassword).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, err
}

func GetId(login string) (string, error) {
	var id string
	err := DB.QueryRow("SELECT id FROM users_creds WHERE login = $1", login).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetLogin(id string) (string, error) {
	var login string

	err := DB.QueryRow("SELECT login FROM users_creds WHERE id = $1", id).Scan(&login)
	if err != nil {
		return "", err
	}
	return login, nil
}

func UpdateInfo(id string, user_info models.UserInfo) error {
	res, err := DB.Exec("UPDATE users_info SET (first_name, last_name, user_birthday, user_email, user_phone) = ($2, $3, $4, $5, $6) WHERE id = $1", id,
		user_info.FirstName, user_info.LastName, user_info.UserBirthday, user_info.UserEmail, user_info.UserPhone)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 0 {
		return nil
	}

	_, err = DB.Exec("INSERT INTO users_info (id, first_name, last_name, user_birthday, user_email, user_phone) VALUES ($1, $2, $3, $4, $5, $6)", id,
		user_info.FirstName, user_info.LastName, user_info.UserBirthday, user_info.UserEmail, user_info.UserPhone)

	if err != nil {
		return err
	}

	return nil
}
