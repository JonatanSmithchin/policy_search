package Models

import (
	Mysql "police_search/Databases"
	"police_search/api"
)

type Footprint struct {
	RecordID int        `gorm:"column:RecID;primaryKey"`
	UserID   int        `gorm:"column:UserID;foreignKey:record_UserID_FK" json:"UserID"`
	IP       string     `gorm:"column:IP"`
	Begin    api.MyTime `gorm:"column:Read_time;type:timestamp" json:"begin"`
	Content  string     `gorm:"column:content" json:"content"`
	Duration int        `gorm:"column:Read_duration"`
	Tag      string     `gorm:"column:tag" json:"tag"`
}

func (f Footprint) TableName() string {
	return "read_record"
}

func (f *Footprint) AddFootprint() (int, error) {
	result := Mysql.DB.Create(f)
	if result.Error != nil {
		return -1, result.Error
	}
	return f.RecordID, nil
}

func (f Footprint) UpdateDuration(duration int) error {
	result := Mysql.DB.Model(&f).Update("Read_duration", duration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
