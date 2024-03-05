package Utility

import (
	"fmt"
)

type UserControllerInterface interface {
	GetUser(id int) string
}

type UserController struct {
}

func (uc *UserController) GetUser(id int) string {
	return fmt.Sprintf("User with ID %d", id)
}
