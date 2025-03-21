

package model

import (
)

type Result struct {
	TeamId          string  `gorm:"primaryKey;unique;index;column:team_id"`
	FileUploadTime  string  `gorm:"column:file_upload_time"`
	JudgeResultTime string  `gorm:"column:judge_result_time"`
	Score           float32 `gorm:"column:score"`
}
