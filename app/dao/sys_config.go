package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
	
)

type SysConfigDao struct {
	db    *gorm.DB
	models *models.SysConfig
}

func NewSysConfigDao(db *gorm.DB) *SysConfigDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysConfigDao{db: db}
}

// Create
func (dao *SysConfigDao) Create(sysConfig models.SysConfig) error {
	return dao.db.Create(&sysConfig).Error
}

// DeleteById
func (dao *SysConfigDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysConfigDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysConfigDao) Update(sysConfig models.SysConfig) error {
	return dao.db.Updates(&sysConfig).Error
}

// GetList
func (dao *SysConfigDao) GetList() (list []models.SysConfig) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysConfigDao) GetPage(page page.PageParam) (list []models.SysConfig, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("config_key", "LIKE", page.FilterMap["config_key"]),asql.Where("config_type", "=", page.FilterMap["config_type"]),asql.Where("description", "LIKE", page.FilterMap["description"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysConfigDao) GetById(id int) (row models.SysConfig) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
