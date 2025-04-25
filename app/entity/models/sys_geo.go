package models 

import (
)

type SysGeo struct { 
    Id int `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"编码"`//编码
    Pid *int `gorm:"column:pid" json:"pid" form:"pid" comment:"上级编码"`//上级编码
    Deep *int `gorm:"column:deep" json:"deep" form:"deep" comment:"深度"`//深度
    Name string `gorm:"column:name" json:"name" form:"name" comment:"名称"`//名称
    PinyinPrefix string `gorm:"column:pinyin_prefix" json:"pinyin_prefix" form:"pinyin_prefix" comment:"拼音首字母"`//拼音首字母
    Pinyin string `gorm:"column:pinyin" json:"pinyin" form:"pinyin" comment:"拼音"`//拼音
    ExtId *int `gorm:"column:ext_id" json:"ext_id" form:"ext_id" comment:"编码"`//编码
    ExtName string `gorm:"column:ext_name" json:"ext_name" form:"ext_name" comment:"全称"`//全称
}

func (SysGeo) TableName() string {
	return "loan.sys_geo"
}