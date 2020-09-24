package entity

import (
	"time"
)

type Memo struct {
	ID    string `gorm:"primary_key"`
	Updated time.Time
	Title   string
	Text    string
}