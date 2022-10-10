package model

import "time"

type FileProviderInfo struct {
	Name string
	At   time.Time
	Size int64
	Uid  string
}
