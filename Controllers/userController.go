package Controllers

import (
	Mysql "police_search/Databases"
	"police_search/Models"
	"police_search/Redis"
	"strconv"
)

func GetUser(name string) *Models.User {
	return Models.GetUserByName(name)
}

func GetAllUsers() *[]*Models.User {
	return Models.GetAllUsers()
}

func CreateNewUser(NewUser *Models.User) (int, error) {
	err := NewUser.CreateUser()
	if err != nil {
		return -1, err
	}
	u := Models.GetUserByName(NewUser.UserName)
	err = Redis.Record(strconv.Itoa(u.UserID), 12)
	if err != nil {
		return -1, err
	}
	return u.UserID, nil
}

func ChangePassword(User *Models.User, password string) error {
	err := User.ChangePassword(password)
	if err != nil {
		return err
	}
	return nil
}

func GetTendency(UserID int) ([]string, error) {
	return Models.FindTendency(UserID)
}

func SetTendency(UserID int, tendencies []string) error {
	tx := Mysql.DB.Begin()
	for _, description := range tendencies {
		tendency := &Models.Tendency{
			UserId:      UserID,
			Description: description,
		}
		err := tendency.AddTendency(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func RemoveTendency(UserID int, tendencies []string) error {
	tx := Mysql.DB.Begin()
	for _, description := range tendencies {
		tendency := &Models.Tendency{
			UserId:      UserID,
			Description: description,
		}
		err := tendency.DeleteTendency(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func AddRecord(f *Models.Footprint) (int, error) {
	recId, err := f.AddFootprint()
	if err != nil {
		return -1, err
	}
	return recId, nil
}

func UpdateRecordDuration(RecId int, duration int) error {
	f := &Models.Footprint{
		RecordID: RecId,
	}
	err := f.UpdateDuration(duration)
	if err != nil {
		return err
	}
	return nil
}
