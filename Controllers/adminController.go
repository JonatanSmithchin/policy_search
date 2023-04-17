package Controllers

import "police_search/Models"

func GetAllAdmins() *[]*Models.Admin {
	return Models.GetAdmins()
}

func CreateAdmin(a *Models.Admin) error {
	err := a.AddAdmin()
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(name string) error {
	a := &Models.Admin{
		Name: name,
	}
	err := a.DeleteAdmin()
	if err != nil {
		return err
	}
	return nil
}
