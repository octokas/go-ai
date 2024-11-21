package models

type User struct {
	BaseModel `bson:",inline"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"-" bson:"password"`
	Name      string `json:"name" bson:"name"`
	Role      string `json:"role" bson:"role"`
}
