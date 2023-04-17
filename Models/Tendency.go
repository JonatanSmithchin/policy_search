package Models

import (
	"gorm.io/gorm"
	Mysql "police_search/Databases"
)

type Tendency struct {
	Id          int    `gorm:"primaryKey;column:TendencyID"`
	UserId      int    `gorm:"foreignKey:UserID_FK;column:UserID"`
	Description string `gorm:"column:tendency_description"`
}

func (t *Tendency) TableName() string {
	return "tendency"
}

func (t *Tendency) AddTendency(tx *gorm.DB) error {
	result := tx.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Tendency) DeleteTendency(tx *gorm.DB) error {
	result := tx.Where("tendency_description=?", t.Description).
		Where("UserID=?", t.UserId).
		Delete(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindTendency(UserID int) ([]string, error) {
	var tendencies []string
	result := Mysql.DB.Table("tendency").
		Select("tendency_description").
		Where("UserID=?", UserID).
		Find(&tendencies)
	if result.Error != nil {
		return nil, result.Error
	}
	return tendencies, nil
}
