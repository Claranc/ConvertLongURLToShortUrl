package model

type Shorttolong struct {
	Shorturl     	string `gorm:"type:varchar(128);not null;"`
	Longurl         string `gorm:"type:varchar(256);not null;primary key"`
}
