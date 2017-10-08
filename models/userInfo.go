package models

type UserInfo struct {
}

func CurrentUser() *UserInfo {
	return &UserInfo{}
}

func (this *UserInfo) SetControllerData(data map[interface{}]interface{}) {
	data["UserName"] = "盛鹏超"
	data["Role"] = "系统管理员"
}

func (this *UserInfo) IsAdmin() bool {
	return true
}

func (this *UserInfo) IsLogin() bool {
	return true
}
