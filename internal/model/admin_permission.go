package model

type AdminPermission struct {
	Model
	// 权限名称
	Name string `json:"name"`

	//请求路径
	HttpPath string `json:"http_path"`

	//请求方法
	HttpMethod string `json:"http_method"`
}

type AdminPermissionList struct {
	//总条数
	Total int64 `json:"total"`

	//页数
	Page int `json:"page"`

	//每页条数
	PageSize int `json:"page_size"`

	//总页数
	LastPage int `json:"last_page"`

	//当前页数据
	Items []*AdminPermission `json:"items"`
}

func (model AdminPermission) TableName() string {
	return "admin_permission"
}
