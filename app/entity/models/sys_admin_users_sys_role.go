package models

type SysAdminUsersSysRole struct {
	SysAdminUsersId int `gorm:"column:sys_admin_users_id;primaryKey" json:"sys_admin_users_id" form:"sys_admin_users_id" comment:"管理员用户标识"` //管理员用户标识
	SysRoleId       int `gorm:"column:sys_role_id;primaryKey" json:"sys_role_id" form:"sys_role_id" comment:"角色标识"`                         //角色标识
}

func (SysAdminUsersSysRole) TableName() string {
	return "loan.sys_admin_users_sys_role"
}
