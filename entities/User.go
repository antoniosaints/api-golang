package entities

import "github.com/google/uuid"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Date string `json:"date"`
}

func NewUser(name string, age int, date string) User {
	return User{
		Id:   int(uuid.New().ID()),
		Name: name,
		Age:  age,
		Date: date,
	}
}

func GetUser(Name string) User {
	return NewUser(Name, 20, "2022-01-01")
}
