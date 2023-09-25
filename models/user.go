package models

type User struct{
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Phone int64 `json:"phone" bson:"phone"`
	DOB string `json:"dob" bson:"dob"`
	Role []string `json:"role" bson:"role"`
	Status string `json:"status" bson:"status"`
}

type UpdateRole struct{
	Name string `json:"name"`
	Role []string `json:"role"`
}

type RoleRequest struct{
	Name string `json:"name"`
	Role string `json:"role"`
}

type Role struct{
	Role string `json:"role"`
	Responsibility string `json:"responsibility"`
	Access string `json:"access"`
}