package Controllers

import "police_search/Models"

func GetAllLogs() *[]*Models.AdminLog {
	return Models.GetLogs()
}

func FindLogsByName(manager string) *[]*Models.AdminLog {
	return Models.FindLogsByName(manager)
}
