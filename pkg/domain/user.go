package domain

type Users struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Email   string `json:"email" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
