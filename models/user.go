package models

type User struct{
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Phone int64 `json:"phone" bson:"phone"`
	DOB string `json:"dob" bson:"dob"`
}

type Role struct{
	Name string `json:"name"`
	Role string `json:"role"`
}