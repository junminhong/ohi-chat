package models

type Member struct {
	ID       int `gorm:"primaryKey"`
	Account  string
	Password string
	token    string
}
