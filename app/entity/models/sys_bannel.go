package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type SysBannel struct {
	Id          int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	Title       string         `gorm:"column:title" json:"title" form:"title" comment:"名称"`                           //名称
	ImageUrl    asql.Json      `gorm:"column:image_url" json:"image_url" form:"image_url" comment:"图片"`               //图片
	ChannelCode asql.Json      `gorm:"column:channel_code" json:"channel_code" form:"channel_code" comment:"渠道唯一标识"`  //渠道唯一标识
	GeoIds      asql.Json      `gorm:"column:geo_ids" json:"geo_ids" form:"geo_ids" comment:"不允许访问城市标识"`              //不允许访问城市标识
	Action      string         `gorm:"column:action;default:null" json:"action" form:"action" comment:"类型操作"`         //类型操作
	Param       string         `gorm:"column:param" json:"param" form:"param" comment:"动态参数"`                         //动态参数
	Link        string         `gorm:"column:link" json:"link" form:"link" comment:"外部链接(和类型操作互斥)"`                   //外部链接(和类型操作互斥)
	Sort        *int           `gorm:"column:sort" json:"sort" form:"sort" comment:"排序"`                              //排序
	ChanelIds   asql.Json      `gorm:"column:chanel_ids" json:"chanel_ids" form:"chanel_ids" comment:"渠道标识"`          //渠道标识
	Status      *int           `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用"`              //状态:1=禁用;2=启用
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`          //创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`          //修改时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
}

func (SysBannel) TableName() string {
	return "loan.sys_bannel"
}
