package domain

import "gorm.io/gorm"

type Resume struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255)"`
	Email            string `gorm:"type:varchar(255)"`
	PhoneNumber      string `gorm:"type:varchar(20)"`
	Designation      string `gorm:"type:varchar(100)"`
	Experience       int    `gorm:"type:int"`
	HighestEducation string `gorm:"type:varchar(100)"`
	Location         string `gorm:"type:varchar(100)"`
	Skills           string `gorm:"type:text"`
}
