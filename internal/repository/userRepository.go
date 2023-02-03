package repository

import (
	"log"

	"github.com/lashifrin/go-sql-class/internal/model"
	"github.com/lashifrin/go-sql-class/internal/store"
)

type UserRepository struct {
	db *store.Store
}

func New(db *store.Store) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(u *model.User) error {
	if err := repo.db.Db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		u.Name,
		u.Email,
	).Scan(&u.ID); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Select() ([]model.User, error) {
	var users []model.User
	rows, err := repo.db.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (repo *UserRepository) SelectById(id int) (model.User, error) {

	var u model.User
	stmt, err := repo.db.Db.Prepare("SELECT * FROM users where id = $1")
	if err != nil {
		return u, err
	}
	defer stmt.Close()

	row, err := stmt.Query(id)
	if err != nil {
		return u, err
	}
	defer row.Close()

	for row.Next() {

		err := row.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return u, err
		}

	}
	if err = row.Err(); err != nil {
		return u, err
	}

	return u, nil
}

func (repo *UserRepository) Delete() error {
	_, err := repo.db.Db.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) TxInsert(users []model.User) error {
	tx, err := repo.db.Db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO users VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		_, err := stmt.Exec(user.ID, user.Name, user.Email)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return nil
}
