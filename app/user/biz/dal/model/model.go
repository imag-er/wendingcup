package model

import ()

type Team struct {
	UUID    string   `gorm:"type:varchar(36);primaryKey"`
	Name    string   `gorm:"type:varchar(64);uniqueIndex;not null"`  // 显式指定类型
	Players []Player `gorm:"foreignKey:TeamID"`
}

type Player struct {
    TeamID    string `gorm:"type:varchar(36);not null;"`
	Name      string `gorm:"type:varchar(64);not null"`
	Phone     string `gorm:"type:varchar(11);not null"`
	StudentId string `gorm:"type:varchar(9);not null"`
}
