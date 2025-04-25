package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type SysRole struct {
	Id            int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`         //标识
	Name          string         `gorm:"column:name" json:"name" form:"name" comment:"角色名称"`                                    //角色名称
	ParentMenuIds asql.Json      `gorm:"column:parent_menu_ids" json:"parent_menu_ids" form:"parent_menu_ids" comment:"一级菜单标识"` //一级菜单标识
	MenuIds       asql.Json      `gorm:"column:menu_ids" json:"menu_ids" form:"menu_ids" comment:"菜单标识"`                        //菜单标识
	ApiIds        asql.Json      `gorm:"column:api_ids" json:"api_ids" form:"api_ids" comment:"API标识"`                          //API标识
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                  //创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                  //修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                           //删除时间
}

func (SysRole) TableName() string {
	return "loan.sys_role"
}
