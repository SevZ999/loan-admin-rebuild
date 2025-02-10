package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id            int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`      //标识
	Phone         string         `gorm:"column:phone" json:"phone" form:"phone" comment:"手机"`                                //手机
	Password      string         `gorm:"column:password" json:"password" form:"password" comment:"密码"`                       //密码
	NickName      string         `gorm:"column:nick_name" json:"nick_name" form:"nick_name" comment:"昵称"`                    //昵称
	AppCode       int            `gorm:"column:app_code" json:"app_code" form:"app_code" comment:"app标识"`                    //app标识
	ChannelCode   string         `gorm:"column:channel_code" json:"channel_code" form:"channel_code" comment:"渠道标识"`         //渠道标识
	VerifiedCount int            `gorm:"column:verified_count" json:"verified_count" form:"verified_count" comment:"认证用户次数"` //认证用户次数
	MatchCount    int            `gorm:"column:match_count" json:"match_count" form:"match_count" comment:"匹配次数"`            //匹配次数
	Type          int            `gorm:"column:type" json:"type" form:"type" comment:"注册类型:1=短信;2=一键登录"`                     //注册类型:1=短信;2=一键登录
	Status        int            `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用;"`                  //状态:1=禁用;2=启用;
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`               //创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`               //修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                        //删除时间
}

func (Users) TableName() string {
	return "loan.users"
}
