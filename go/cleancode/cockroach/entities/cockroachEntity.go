package entities

import "time"

type (
	InsertCockroachDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Cockroach struct {
		Id        uint32    `josn:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}
	CockroachPushNotificationDto struct {
		Title        string    `json:"title"`
		Amount       uint32    `json:"amount"`
		ReportedTime time.Time `json:"createdAt"`
	}
)
