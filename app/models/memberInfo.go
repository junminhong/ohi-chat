package models

type MemberInfo struct {
	ID       int `gorm:"primaryKey"`
	MemberID int
	Name     string
	Age      int8
}
