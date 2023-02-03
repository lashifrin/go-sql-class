package main

import (
	"fmt"

	"github.com/lashifrin/go-sql-class/internal/model"
	"github.com/lashifrin/go-sql-class/internal/repository"
	"github.com/lashifrin/go-sql-class/internal/store"
)

func main() {
	s := store.New()
	if err := s.Open(); err != nil {
		panic(err.Error())
	}
	defer s.Close()

	user := model.New("Rob Pike", "alan@email.com")

	repository := repository.New(s)

	//Create a single user record
	error := repository.Create(user)
	if error != nil {
		fmt.Println(error.Error())
		panic(error.Error())
	}

	//Query what you have in the table
	results, err := repository.Select()
	if err != nil {
		panic(error.Error())
	}

	for _, user := range results {
		fmt.Printf("Result of simple query: Id: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	//Query what you have for an id in the table
	record, err := repository.SelectById(1)
	if err != nil {
		panic(error.Error())
	}
	fmt.Printf("Result of Prepared Statement: Id: %d, Name: %s, Email: %s\n", record.ID, record.Name, record.Email)

	//Delete all the rows in the table
	err = repository.Delete()
	if err != nil {
		panic(error.Error())
	}

	//Use transaction to populate a list(slises) of records
	err = repository.TxInsert(getSomeUsers())
	if err != nil {
		panic(error.Error())
	}

	//Display content of the table
	results, err = repository.Select()
	if err != nil {
		panic(error.Error())
	}

	for _, user := range results {
		fmt.Printf("Result of simple query: Id: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	//Finally clear the table
	err = repository.Delete()
	if err != nil {
		panic(error.Error())
	}

	fmt.Println("End of the Run... Db is empty")
}

func getSomeUsers() []model.User {
	users := []model.User{
		{ID: 1, Name: "Rob Pike", Email: "rob@google.com"},
		{ID: 2, Name: "Ken Thompson", Email: "ken@email.com"},
		{ID: 3, Name: "Robert Griesemer", Email: "robert@email.com"},
	}
	return users
}
