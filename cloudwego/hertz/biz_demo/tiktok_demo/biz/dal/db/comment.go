package db

import (
	"time"

	"github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/pkg/constants"
	"gorm.io/gorm"
)

type Comment struct {
	ID          int64          `json:"id,omitempty"`
	UserId      int64          `json:"user_id,omitempty"`
	VideoId     int64          `json:"video_id,omitempty"`
	CommentText string         `json:"comment_text,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}
