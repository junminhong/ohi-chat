package models

type Member struct {
	ID            int `gorm:"primaryKey"`
	Account       string
	Password      string
	Access_token  string
	Refresh_token string
}
