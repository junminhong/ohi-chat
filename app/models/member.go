package models

type Member struct {
	ID       int `gorm:"primaryKey"`
	Account  string
	Password int8
	token    string
	MemberID int
}
