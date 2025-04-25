package models

import (
	"gorm.io/gorm"
	"time"
)

type ChannelNewUserReport struct {
	Id              int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"主键"`                    //主键
	Date            string         `gorm:"column:date" json:"date" form:"date" comment:"统计日期"`                                               //统计日期
	SysAdminUsersId int            `gorm:"column:sys_admin_users_id" json:"sys_admin_users_id" form:"sys_admin_users_id" comment:"系统用户"`     //系统用户
	ChannelName     string         `gorm:"column:channel_name" json:"channel_name" form:"channel_name" comment:"渠道名称"`                       //渠道名称
	ChannelKey      string         `gorm:"column:channel_key" json:"channel_key" form:"channel_key" comment:"渠道Key"`                         //渠道Key
	Platform        string         `gorm:"column:platform;default:null" json:"platform" form:"platform" comment:"平台类型（H5 或 APP）"`            //平台类型（H5 或 APP）
	UvClicks        int            `gorm:"column:uv_clicks" json:"uv_clicks" form:"uv_clicks" comment:"UV 点击量"`                              //UV 点击量
	Activation      int            `gorm:"column:activation" json:"activation" form:"activation" comment:"激活/登录数量"`                          //激活/登录数量
	Registrations   int            `gorm:"column:registrations" json:"registrations" form:"registrations" comment:"注册数量"`                    //注册数量
	Submissions     int            `gorm:"column:submissions" json:"submissions" form:"submissions" comment:"进件数量"`                          //进件数量
	Completions     int            `gorm:"column:completions" json:"completions" form:"completions" comment:"完件数量"`                          //完件数量
	AppDownloads    int            `gorm:"column:app_downloads" json:"app_downloads" form:"app_downloads" comment:"APP 下载量（仅适用于 APP 平台）"`    //APP 下载量（仅适用于 APP 平台）
	SettlementPrice float64        `gorm:"column:settlement_price" json:"settlement_price" form:"settlement_price" comment:"结算价"`            //结算价
	SettlementType  string         `gorm:"column:settlement_type;default:null" json:"settlement_type" form:"settlement_type" comment:"结算类型"` //结算类型
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"记录时间"`                             //记录时间
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                             //更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                      //删除时间
}

func (ChannelNewUserReport) TableName() string {
	return "loan.channel_new_user_report"
}
