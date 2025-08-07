package user

type User struct {
	ID       uint   `gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
