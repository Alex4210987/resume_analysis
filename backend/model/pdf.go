package model

import (
	"database/sql"
	"time"
)

type PDF struct {
	PDFId    int64 `gorm:"primarykey"`
	PDFUrl   string
	PDFName  string
	UserId   int64
	PDFSize  int64
	Content  string
	CreateAt time.Time
	DeleteAt sql.NullTime
	UpdateAt time.Time
	ExpireAt time.Time
}
