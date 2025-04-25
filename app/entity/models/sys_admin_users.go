package models

import (
	"github.com/small-ek/antgo/utils/password"
	"gorm.io/gorm"
	"time"
)

type SysAdminUsers struct {
	Id                int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                                                                                             //标识
	Username          string         `gorm:"column:username" json:"username" form:"username" comment:"用户名"`                                                                                                             //用户名
	Password          string         `gorm:"column:password" json:"password" form:"password" comment:"密码"`                                                                                                              //密码
	NickName          string         `gorm:"column:nick_name" json:"nick_name" form:"nick_name" comment:"昵称"`                                                                                                           //昵称
	Phone             string         `gorm:"column:phone" json:"phone" form:"phone" comment:"手机"`                                                                                                                       //手机
	Email             string         `gorm:"column:email" json:"email" form:"email" comment:"电子邮件"`                                                                                                                     //电子邮件
	Status            *int           `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用;"`                                                                                                         //状态:1=禁用;2=启用;
	RoleId            *int           `gorm:"column:role_id" json:"role_id" form:"role_id" comment:"当前角色标识"`                                                                                                             //当前角色标识
	InstitutionInfoId *int           `gorm:"column:institution_info_id" json:"institution_info_id" form:"institution_info_id" comment:"所属机构"`                                                                           //当前角色标识
	ParentId          *int           `gorm:"column:parent_id" json:"parent_id" form:"parent_id" comment:"父级标识"`                                                                                                         //当前角色标识
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                                                                                      //创建时间
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                                                                                                      //修改时间
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                                                                               //删除时间
	SysRole           []SysRole      `gorm:"many2many:sys_admin_users_sys_role;foreignKey:Id;References:Id;joinForeignKey:sys_admin_users_id;joinReferences:sys_role_id" json:"sys_role" form:"sys_role" comment:"角色" ` //角色
}

func (SysAdminUsers) TableName() string {
	return "loan.sys_admin_users"
}

func (m *SysAdminUsers) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		pwd, err2 := password.Generate(m.Password)
		if err2 != nil {
			return err2
		}
		m.Password = pwd
	}
	return nil
}

func (m *SysAdminUsers) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		pwd, err2 := password.Generate(m.Password)
		if err2 != nil {
			return err2
		}
		m.Password = pwd
	}
	return nil
}
