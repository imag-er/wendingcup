package model
import (
)

type Team struct {
    UUID      string `gorm:"primarykey"`
    TeamName  string
    Players   []Player `gorm:"foreignKey:TeamUUID"`
}

type Player struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string
    PhoneNumber string
    StudentId  string
    TeamUUID   string `gorm:"index;not null"` // 外键，指向Team的UUID
}


