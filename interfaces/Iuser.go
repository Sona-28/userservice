package interfaces

import "sample/models"

type IUser interface{
	AddUser(user *models.User) (string, error)
	UpdateRole(role *models.UpdateRole) (string, error)
	ListFeatures(list *models.Role) (*models.Role, error)
	DisableUser(list *models.User) (string, error)
	EnableUser(list *models.User) (string, error)
	AssociateRole(list *models.RoleRequest) (string, error)
}