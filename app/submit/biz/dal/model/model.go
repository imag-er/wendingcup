package model

import (
	"gorm.io/gorm"
)

// 定义提交状态的枚举类型
const (
	StatusUploaded  string = "uploaded"
	StatusPending   string = "pending"
	StatusCompleted string = "completed"
	StatusFailed    string = "failed"
)

type Submit struct {
	gorm.Model
	TeamId  string
	Status  string
	Message string
	Score float32
}
