package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type InstitutionInfo struct {
	Id                   int                  `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                                   //标识
	InstitutionName      string               `gorm:"column:institution_name" json:"institution_name" form:"institution_name" comment:"机构名称"`                          //机构名称
	Weight               *int                 `gorm:"column:weight" json:"weight" form:"weight" comment:"权重"`                                                          //权重
	ChannelId            asql.Json            `gorm:"column:channel_id" json:"channel_id" form:"channel_id" comment:"渠道标识"`                                            //权重
	Status               *int                 `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用"`                                                //状态:1=禁用;2=启用
	Type                 *int                 `gorm:"column:type" json:"type" form:"type" comment:"类型:1=债务;2=助贷;3=芝麻"`                                                 //类型:1=债务;2=助贷;3=芝麻
	IsPush               *int                 `gorm:"column:is_push" json:"is_push" form:"is_push" comment:"是否需要推送;1=不需要;2=需要"`                                        //是否需要推送;1=不需要;2=需要
	Logo                 asql.Json            `gorm:"column:logo" json:"logo" form:"logo" comment:"机构图片"`                                                              //机构图片
	BusinessLicense      asql.Json            `gorm:"column:business_license" json:"business_license" form:"business_license" comment:"营业执照"`                          //营业执照
	ComplianceCommitment asql.Json            `gorm:"column:compliance_commitment" json:"compliance_commitment" form:"compliance_commitment" comment:"业务合规承诺"`         //业务合规承诺
	PushAgreement        asql.Json            `gorm:"column:push_agreement" json:"push_agreement" form:"push_agreement" comment:"推送协议"`                                //推送协议
	CreatedAt            time.Time            `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                            //创建时间
	UpdatedAt            time.Time            `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                            //更新时间
	DeletedAt            gorm.DeletedAt       `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                     //删除时间
	InstitutionInfoGeo   []InstitutionInfoGeo `gorm:"foreignKey:InstitutionInfoId;references:Id" json:"institution_info_geo" form:"institution_info_geo" comment:"城市"` //城市
}

func (InstitutionInfo) TableName() string {
	return "loan.institution_info"
}
