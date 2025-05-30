package models

import (
	"gorm.io/gorm"
	"time"
)

type SysMenu struct {
	Id        int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	Title     string         `gorm:"column:title" json:"title" form:"title" comment:"菜单名称"`                         //菜单名称
	Path      string         `gorm:"column:path" json:"path" form:"path" comment:"菜单跳转地址"`                          //菜单跳转地址
	Component string         `gorm:"column:component" json:"component" form:"component" comment:"前端跳转地址"`           //前端跳转地址
	Param     string         `gorm:"column:param" json:"param" form:"param" comment:"前端参数"`                         //前端参数
	Status    *int           `gorm:"column:status" json:"status" form:"status" comment:"是否隐藏0=禁用;1=启用"`             //是否隐藏1=禁用;2=启用
	ParentId  *int           `gorm:"column:parent_id" json:"parent_id" form:"parent_id" comment:"父级标识"`             //父级标识
	Sort      *int           `gorm:"column:sort" json:"sort" form:"sort" comment:"排序"`                              //排序
	Icon      string         `gorm:"column:icon" json:"icon" form:"icon" comment:"图标"`                              //图标
	IsTab     *int           `gorm:"column:is_tab" json:"is_tab" form:"is_tab" comment:"便签栏"`                       //便签栏
	IsCache   *int           `gorm:"column:is_cache" json:"is_cache" form:"is_cache" comment:"页面缓存"`                //页面缓存
	CreatedAt time.Time      `gorm:"column:created_at" json:"-" form:"created_at" comment:"创建时间"`                   //创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-" form:"updated_at" comment:"修改时间"`                   //修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
}

func (SysMenu) TableName() string {
	return "loan.sys_menu"
}
