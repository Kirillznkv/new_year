package store

import (
	"database/sql"
	"errors"
	"github.com/Kirillznkv/new_year/api/internal/model"
	"log"
)

type UsersRepository struct {
	store *Store
}

func (r *UsersRepository) Create(u *model.User) error {
	if err := r.store.db.QueryRow("INSERT INTO users (first_name, second_name) VALUES ($1, $2) RETURNING id",
		u.FirstName,
		u.SecondName,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, first_name, second_name FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.SecondName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return u, nil
}

func (r *UsersRepository) GetUsers() []*model.User {
	var users []*model.User

	rows, err := r.store.db.Query("SELECT id, first_name, second_name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		usr := &model.User{}

		if err := rows.Scan(&usr.ID, &usr.FirstName, &usr.SecondName); err != nil {
			log.Fatal(err)
		}

		users = append(users, usr)
	}

	return users
}
