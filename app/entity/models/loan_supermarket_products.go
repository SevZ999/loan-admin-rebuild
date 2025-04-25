package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type LoanSupermarketProducts struct {
	Id                 int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                          //标识
	ProductName        string         `gorm:"column:product_name" json:"product_name" form:"product_name" comment:"产品名称"`                             //产品名称
	ProductImage       asql.Json      `gorm:"column:product_image" json:"product_image" form:"product_image" comment:"产品图片"`                          //产品图片
	Description        string         `gorm:"column:description" json:"description" form:"description" comment:"描述"`                                  //描述
	TagType            asql.Json      `gorm:"column:tag_type" json:"tag_type" form:"tag_type" comment:"标签类型:1=挂牌机构;2=高通过率"`                           //标签类型:1=挂牌机构;2=高通过率
	MaxLoanAmount      *int           `gorm:"column:max_loan_amount" json:"max_loan_amount" form:"max_loan_amount" comment:"最高可借，单位为万元"`              //最高可借，单位为万元
	FundProvider       string         `gorm:"column:fund_provider" json:"fund_provider" form:"fund_provider" comment:"资金方"`                           //资金方
	LoanTerm           string         `gorm:"column:loan_term" json:"loan_term" form:"loan_term" comment:"贷款期限"`                                      //贷款期限
	AnnualInterestRate string         `gorm:"column:annual_interest_rate" json:"annual_interest_rate" form:"annual_interest_rate" comment:"年利率（百分比）"` //年利率（百分比）
	UnitPrice          float64        `gorm:"column:unit_price" json:"unit_price" form:"unit_price" comment:"单价"`                                     //单价
	RedirectUrl        string         `gorm:"column:redirect_url" json:"redirect_url" form:"redirect_url" comment:"跳转链接"`                             //跳转链接
	IsAlipay           *int           `gorm:"column:is_alipay" json:"is_alipay" form:"is_alipay" comment:"是否为支付宝"`                                    //是否为支付宝
	StartTime          string         `gorm:"column:start_time" json:"start_time" form:"start_time" comment:"开始时间"`                                   //开始时间
	EndTime            string         `gorm:"column:end_time" json:"end_time" form:"end_time" comment:"结束时间"`                                         //结束时间
	Type               *int           `gorm:"column:type" json:"type" form:"type" comment:"类型:1=跳转链接;2=下载链接"`                                         //类型:1=跳转链接;2=下载链接
	Status             *int           `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用"`                                       //状态:1=禁用;2=启用
	Sort               int            `gorm:"column:sort" json:"sort" form:"sort" comment:"排序"`                                                       //排序
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                   //创建时间
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                   //更新时间
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                            //删除时间
}

func (LoanSupermarketProducts) TableName() string {
	return "loan.loan_supermarket_products"
}
