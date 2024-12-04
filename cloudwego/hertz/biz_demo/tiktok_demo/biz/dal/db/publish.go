package db

import (
	"time"

	"github.com/chhz0/go-hello/cloudwego/hertz/biz_demo/tiktok_demo/pkg/constants"
)

type Video struct {
	ID         int64
	AuthorID   int64
	PlayURL    string
	CoverURL   string
	PulishTime time.Time
	Title      string
}

func (Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(video *Video) (Video_id int64, err error) {
	err = DB.Create(video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, err
}
