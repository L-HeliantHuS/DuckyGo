package serializer

import "DuckyGo/model"

// Base 主要序列化器
type Base struct {
	Base string `json:"data"`
}

// BaseAll 多个数据序列化器
type BaseAll struct {
	Results []Base `json:"list"`
	Count   int    `json:"count"`
}

// BaseResponse 主要序列化响应
func BaseResponse(db model.User) Base {
	return Base{}
}

// BaseAllResponse 多个数据序列化响应
func BaseAllResponse(db []model.User, count int) BaseAll {
	var result []Base
	for _, i := range db {
		result = append(result, BaseResponse(i))
	}
	return BaseAll{}
}
