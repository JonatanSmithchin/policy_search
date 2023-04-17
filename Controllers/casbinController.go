package Controllers

import (
	Mysql "police_search/Databases"
	"police_search/Models"
)

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}
type CasbinCreateRequest struct {
	RoleId      string       `form:"role_id" json:"role_id" form:"role_id" description:"角色ID"`
	CasbinInfos []CasbinInfo `form:"casbin_infos" json:"casbin_infos" description:"权限模型列表"`
}

type CasbinListResponse struct {
	List []CasbinInfo `json:"list" form:"list"`
}

type CasbinListRequest struct {
	RoleID string `json:"role_id" form:"role_id"`
}

func CreateCasbin(param *CasbinCreateRequest) error {

	for _, info := range param.CasbinInfos {
		c := &Models.CasbinModel{
			PType:  "p",
			RoleId: param.RoleId,
			Path:   info.Path,
			Method: info.Method,
		}
		err := c.Create(Mysql.DB)
		if err != nil {
			return err
		}
	}
	return nil
}

func CasbinList(roleID string) [][]string {
	c := &Models.CasbinModel{RoleId: roleID}
	return c.List(Mysql.DB)
}

//func (s Service) CasbinCreate(param *CasbinCreateRequest) error {
//	for _, v := range param.CasbinInfos {
//		err := s.dao.CasbinCreate(param.RoleId, v.Path, v.Method)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (s Service) CasbinList(param *CasbinListRequest) [][]string {
//	return s.dao.CasbinList(param.RoleID)
//}
