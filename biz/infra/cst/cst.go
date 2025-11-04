package cst

// 数据库相关
const (
	ID             = "_id"
	Status         = "status"
	Phone          = "phone"
	Name           = "name"
	Address        = "address"
	Contact        = "contact"
	CreateTime     = "createTime"
	UpdateTime     = "updateTime"
	DeleteTime     = "deleteTime"
	Password       = "password"
	NotEqual       = "$ne"
	Account        = "account"
	VerifyPassword = "verifyPassword"
	VerifyType     = "verifyType"
	Form           = "form"
)

// password
const (
	DefaultPassword = "123456"
	UpdateByOldPwd  = 0
	UpdateByCode    = 1
)

// status
const (
	Active  = 0
	Deleted = 1
)

// gender
const (
	Unknown = 0
	Male    = 1
	Female  = 2
)

// code type
const (
	CodeTypePhone = 0
	COdeTypeCode  = 1
)

// config type
const (
	Chain   = 0
	End2End = 1
)
