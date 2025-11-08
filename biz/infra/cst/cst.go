package cst

// 数据库相关
const (
	ID         = "_id"
	Status     = "status"
	Phone      = "phone"
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
	NotEqual   = "$ne"
	Account    = "account"
)

// password
const (
	DefaultPassword = "123456"
)

// 前端字段相关
const (
	AuthTypePhonePassword     = "phone-password"
	AuthTypePhoneCode         = "phone-code"
	AuthTypeStudentIDPassword = "studentId-password"
	AuthTypeWeakAuth          = "weak"
	AuthTypeOldPassword       = "oldPassword"
	AuthTypeCode              = "code"
)
