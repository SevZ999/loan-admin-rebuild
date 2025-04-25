package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type Channel struct {
	Id               int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"主键，自增"`                                                //主键，自增
	SysAdminUsersId  int            `gorm:"column:sys_admin_users_id" json:"sys_admin_users_id" form:"sys_admin_users_id" comment:"所属合作渠道账号标识"`                              //所属合作渠道账号标识
	ChannelName      string         `gorm:"column:channel_name" json:"channel_name" form:"channel_name" comment:"渠道名称"`                                                      //上游渠道名称
	ChannelType      *int           `gorm:"column:channel_type" json:"channel_type" form:"channel_type" comment:"渠道类型:1=落地页;2=应用市场"`                                         //上游渠道类型:1=落地页;2=应用市场
	LandingPageType  *int           `gorm:"column:landing_page_type" json:"landing_page_type" form:"landing_page_type" comment:"跳转方式1=注册后下载APP;2=注册填写后下载;3=立即下载APP;4=注册后填写"` //跳转方式1=注册后下载APP;2=注册填写后下载;3=立即下载APP
	ChannelCode      string         `gorm:"column:channel_code" json:"channel_code" form:"channel_code" comment:"渠道唯一标识"`                                                    //上游渠道唯一标识
	Platform         string         `gorm:"column:platform" json:"platform" form:"platform" comment:"平台类型（H5 或 APP）"`                                                        //上游渠道唯一标识
	DeductionRate    float64        `gorm:"column:deduction_rate" json:"deduction_rate" form:"deduction_rate" comment:"扣量比例（百分比）"`                                           //扣量比例（百分比）
	PhoneBlacklist   asql.Json      `gorm:"column:phone_blacklist" json:"phone_blacklist" form:"phone_blacklist" comment:"手机号黑名单"`                                           //手机号黑名单
	ProcessPhone     asql.Json      `gorm:"column:process_phone" json:"process_phone" form:"process_phone" comment:"流程手机号"`                                                  //流程手机号
	LoanFailurePhone asql.Json      `gorm:"column:loan_failure_phone" json:"loan_failure_phone" form:"loan_failure_phone" comment:"失败手机号"`                                   //失败手机号
	LoanSuccessPhone asql.Json      `gorm:"column:loan_success_phone" json:"loan_success_phone" form:"loan_success_phone" comment:"成功手机号"`                                   //成功手机号
	GeoIds           asql.Json      `gorm:"column:geo_ids" json:"geo_ids" form:"geo_ids" comment:"不允许访问城市标识"`                                                                //不允许访问城市标识
	LandingPageUrl   string         `gorm:"column:landing_page_url;default:null" json:"landing_page_url" form:"landing_page_url" comment:"落地页链接"`                            //落地页链接
	CustomerPrice    float64        `gorm:"column:customer_price" json:"customer_price" form:"customer_price" comment:"获客单价"`                                                //获客单价
	ClickCount       *int           `gorm:"column:click_count" json:"click_count" form:"click_count" comment:"点击总数"`                                                         //点击总数
	Status           *int           `gorm:"column:status" json:"status" form:"status" comment:"是否开启1=关闭，2=开启"`                                                               //是否开启1=关闭，2=开启
	FormType         *int           `gorm:"column:form_type" json:"form_type" form:"form_type" comment:"填写类型1=贷款，2=表单"`                                                      //填写类型1=贷款，2=表单
	IsIdCard         *int           `gorm:"column:is_id_card" json:"is_id_card" form:"is_id_card" comment:"是否显示身份证:1=不显示;2=显示"`                                              //是否显示身份证:1=不显示;2=显示
	IsLoanExcess     *int           `gorm:"column:is_loan_excess" json:"is_loan_excess" form:"is_loan_excess" comment:"是否有贷超:1=否:2=是"`                                       //是否有贷超:1=否:2=是
	IsLoginRequired  *int           `gorm:"column:is_login_required" json:"is_login_required" form:"is_login_required" comment:"是否强制登录:1=否;2=是"`                             //是否强制登录:1=否;2=是
	SettlementType   string         `gorm:"column:settlement_type;default:null" json:"settlement_type" form:"settlement_type" comment:"结算类型"`                                //结算类型
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                                            //创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                                            //更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                                     //删除时间
	SysAdminUsers    SysAdminUsers  `gorm:"foreignKey:Id;references:SysAdminUsersId" json:"sys_admin_users" form:"sys_admin_users" comment:"系统账号"`                           //系统账号
}

func (Channel) TableName() string {
	return "loan.channel"
}
