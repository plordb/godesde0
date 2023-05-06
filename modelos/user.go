package modelos

import (
	"time"
)

type User struct {
	Id       int
	name     string
	CreateAt time.Time
	Status   bool
}

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.name = name
	usuario.CreateAt = createdAt
	usuario.Status = status
}
