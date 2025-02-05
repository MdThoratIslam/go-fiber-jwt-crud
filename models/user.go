package models

// User model with auto-increment ID
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Phone    string `gorm:"type:varchar(255);unique;not null" json:"phone"`
	Address  string `gorm:"type:varchar(255);default:null" json:"address"`
	Gender   string `gorm:"type:varchar(255);default:null" json:"gender"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Age      int    `gorm:"type:int;default:null" json:"age"`
}
