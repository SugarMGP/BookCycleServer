package models

import "gorm.io/gorm"

// Book 书籍
type Book struct {
	gorm.Model
	UserID       uint
	Name         string // 书名
	Author       string // 作者
	Course       string // 适用课程
	Edition      string // 版次
	Publisher    string // 出版社
	Completeness string // 完好程度
	Img          string // 图片
	Price        string // 价格
	Note         string // 备注
	Status       uint   // 状态 1已上架 2已售出 3待审核 4审核不通过
	Reason       string // 审核不通过原因
}
