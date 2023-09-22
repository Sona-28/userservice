package interfaces

import "sample/models"

type IUser interface{
	AddUser(user *models.User) (string, error)
	AddRole(role *models.Role) (string, error)
}