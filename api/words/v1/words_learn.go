package v1

import "github.com/gogf/gf/v2/frame/g"

type RandListReq struct {
	g.Meta `path:"words/rand" method:"get" summary:"随机获取单词"`
	Limit uint `json:"limit" dc:"数量限制默认10"`
}

type RandListRes struct {
	List []List `json:"list" dc:"单词列表"`	
}

type SetLevelReq struct {
	g.Meta `path:"words/{id}/level" method:"patch" summary:"设置单词熟练度"`
	Id uint `json:"id" dc:"单词ID" v:"required#请输入单词ID"`
	Level ProficiencyLevel `json:"level" dc:"熟练度" v:"required|between:1,5#熟练度必须在1-5之间"`	
}

type SetLevelRes struct {
	
}