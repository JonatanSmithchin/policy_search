package Models

import Mysql "police_search/Databases"

type AdminLog struct {
	ID        int    `json:"id"`
	AdminId   uint   `json:"admin_id"`
	AdminName string `json:"admin_name"`
	Method    string `json:"method"`
	Ip        string `json:"ip"`
	Url       string `json:"url"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func GetLogs() *[]*AdminLog {
	var logs []*AdminLog
	Mysql.DB.Find(&logs)
	return &logs
}

func FindLogsByName(manager string) *[]*AdminLog {
	var logs []*AdminLog
	Mysql.DB.Where("admin_name=?", manager).Find(&logs)
	return &logs
}

func (l *AdminLog) TableName() string {
	return "admin_log"
}

func (l *AdminLog) CreateLog(log AdminLog) {
	Mysql.DB.Create(&log)
}
