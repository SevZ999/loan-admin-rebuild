package models

import (
	"gorm.io/gorm"
	"time"
)

type SysApi struct {
	Id        int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	Title     string         `gorm:"column:title" json:"title" form:"title" comment:"API名称"`                        //API名称
	Path      string         `gorm:"column:path" json:"path" form:"path" comment:"API路径"`                           //API路径
	Group     string         `gorm:"column:group" json:"group" form:"group" comment:"API分组"`                        //API分组
	Method    string         `gorm:"column:method;default:null" json:"method" form:"method" comment:"请求类型"`         //请求类型
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`          //创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`          //修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
}

func (SysApi) TableName() string {
	return "loan.sys_api"
}
