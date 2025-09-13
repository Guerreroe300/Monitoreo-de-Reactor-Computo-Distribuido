package model

import "time"

type Temperature struct{
	Temperature float32 `json:"temperature"`
	Date time.Time `json:"date"`
}
