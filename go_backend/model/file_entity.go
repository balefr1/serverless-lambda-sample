package model

import (
	"fmt"
	"time"
)

type CustomTime time.Time

func (t CustomTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2 15:06"))
	return []byte(stamp), nil
}

type FileEntity struct {
	FileName  string     `json:"filename"`
	CreatedAt CustomTime `json:"at"`
	FromIP    string     `json:"ip"`
}
