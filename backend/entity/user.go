package entity

type User struct {
	ID           int            `gorm:"PrimaryKey" json:"id"`
	FullName     string         `json:"full_name"`
	Address      string         `json:"address"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	Passwordlist []PasswordList `gorm:"ForeignKey:UserID"`
}

type UserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
