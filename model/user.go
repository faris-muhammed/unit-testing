package model

type UserModel struct {
	Id       uint   `gorm:"primarykey"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}

// type OTP struct {
// 	Id        uint `gorm:"primaryKey"`
// 	Email     string
// 	OTP       string
// 	createdAt time.Time
// 	ExpiresAt time.Time
// 	Verified  bool `gorm:"default:false"`
// }
