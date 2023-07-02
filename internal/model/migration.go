package model

import "time"

type Migration struct {
	ID        uint      `gorm:"->;primarykey"`
	Timestamp time.Time `gorm:"column:timestamp;not null"`
	Name      string    `gorm:"column:name;not null;unique"`
}
