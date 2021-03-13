package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpData struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Age             uint   `json:"age" binding:"required"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UpdatedData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   uint   `json:"age" binding:"required"`
}
