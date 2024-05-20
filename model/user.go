package model

import "gorm.io/gorm"

type User struct {
	// gorm会将字段名自动在数据库中映射为蛇形命名
	// 大写字母开头说明是可导出的，即publish，小写则是私有的
	gorm.Model
	UserId uint
	Email  string
	Salt   string
}
