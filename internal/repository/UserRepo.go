package repository

import (
	"log"

	"github.com/jmoiron/sqlx" // Зручна обгортка над database/sql
	_ "github.com/lib/pq"     // PostgreSQL driver
)
type User struct {
	user_id   string
	nickname  string
	username  string
	date      string
	status    string
	stars     int64
	comets    int64
}
func AddUser(user_id string, nickname string, username string, date string, status string, stars int64, comets int64) error {
	db, err := sqlx.Connect("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.NamedExec(`INSERT INTO users (user_id, nickname, username, date, status, stars, comets)
		VALUES (:user_id, :nickname, :username, :date, :status, :stars, :comets) ON CONFLICT (user_id) DO NOTHING`,
		&User{
			user_id:  user_id,
			nickname: nickname,
			username: username,
			date:     date,
			status:   status,
			stars:    stars,
			comets:  comets,
		})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}