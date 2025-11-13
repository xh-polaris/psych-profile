package cst

// 数据库相关
const (
	ID         = "_id"
	Status     = "status"
	Phone      = "phone"
	StudentID  = "studentId"
	Code       = "code"
	Name       = "name"
	UnitID     = "unitId"
	Gender     = "gender"
	Birth      = "birth"
	EnrollYear = "enrollYear"
	Grade      = "grade"
	Class      = "class"
	Address    = "address"
	Contact    = "contact"
	CreateTime = "createTime"
	UpdateTime = "updateTime"
	DeleteTime = "deleteTime"
	Password   = "password"
)

// 前端字段相关
const (
	// TODO: studentId和Phone都是Code
	AuthTypePhonePassword = "phone-password"
	AuthTypePhoneCode     = "phone-code"
	AuthTypePassword      = "password"
	AuthTypeCode          = "code"
)
