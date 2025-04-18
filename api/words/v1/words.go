package v1

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type CreateReq struct {
	g.Meta `path:"words" method:"post" tags:"单词" sm:"创建单词"`	
	Word string `json:"word" dc:"单词" v:"required#请输入单词"`
	Definition string `json:"definition" dc:"定义" v:"required#请输入定义"`
	ExampleSentence string `json:"example" dc:"例句" v:"required#请输入例句"`
	ChineseTranslation string `json:"chineseTranslation" dc:"中文翻译" v:"required#请输入中文翻译"`
	Pronunciation string `json:"pronunciation" dc:"音标" v:"required#请输入音标"`
	ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel" dc:"熟练度" v:"required#请输入熟练度"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta `path:"words/{id}" method:"put" tags:"单词" sm:"更新单词"`
	Id uint `json:"id" dc:"单词ID" v:"required#请输入单词ID"`
	Word string `json:"word" dc:"单词" v:"required#请输入单词"`
	Definition string `json:"definition" dc:"定义" v:"required#请输入定义"`
	ExampleSentence string `json:"example" dc:"例句" v:"required#请输入例句"`
	ChineseTranslation string `json:"chineseTranslation" dc:"中文翻译" v:"required#请输入中文翻译"`
	Pronunciation string `json:"pronunciation" dc:"音标" v:"required#请输入音标"`
	ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel" dc:"熟练度" v:"required|between:1,5#请输入熟练度|熟练度只能是1-5"`	
}

type UpdateRes struct {	
}


type List struct {
	Id uint `json:"id" dc:"单词ID"`
	Word string `json:"word" dc:"单词"`
	Definition string `json:"definition" dc:"定义"`
	ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel" dc:"熟练度"`
}

type ListReq struct {
	g.Meta `path:"words" method:"get" tags:"单词" sm:"获取单词列表"`
	Word string `json:"word" dc:"单词"`
	Page int `json:"page" dc:"页码"`
	Size int `json:"size" dc:"每页数量"`
}

type ListRes struct {
	List []List `json:"list" dc:"单词列表"`
	Total int `json:"total" dc:"单词总数"`
}

type DetailReq struct {
	g.Meta `path:"words/{id}" method:"get" tags:"单词" sm:"获取单词详情"`
	Id uint `json:"id" dc:"单词ID" v:"required#请输入单词ID"`
}

type DetailRes struct {
	Id	uint `json:"id" dc:"单词ID"`
	Word string `json:"word" dc:"单词"`
	Definition string `json:"definition" dc:"定义"`
	ExampleSentence string `json:"example" dc:"例句"`
	ChineseTranslation string `json:"chineseTranslation" dc:"中文翻译"`
	Pronunciation string `json:"pronunciation" dc:"音标"`
	ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel" dc:"熟练度"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

type DeleteReq struct {
	g.Meta `path:"words/{id}" method:"delete" tags:"单词" sm:"删除单词"`
	Id uint `json:"id" dc:"单词ID" v:"required#请输入单词ID"`
}

type DeleteRes struct {
	
}

