package repository

import (
	"github.com/lashifrin/go-sql-class/internal/model"
)

type Repository interface {
	Create(*model.User) (*model.User, error)
	Select() ([]model.User, error)
	SelectById(int)
	Delete()
	TxInsert([]model.User)
}
