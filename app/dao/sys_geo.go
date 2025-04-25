package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
	
)

type SysGeoDao struct {
	db    *gorm.DB
	models *models.SysGeo
}

func NewSysGeoDao(db *gorm.DB) *SysGeoDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysGeoDao{db: db}
}

// Create
func (dao *SysGeoDao) Create(sysGeo models.SysGeo) error {
	return dao.db.Create(&sysGeo).Error
}

// DeleteById
func (dao *SysGeoDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysGeoDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysGeoDao) Update(sysGeo models.SysGeo) error {
	return dao.db.Updates(&sysGeo).Error
}

// GetList
func (dao *SysGeoDao) GetList() (list []models.SysGeo) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysGeoDao) GetPage(page page.PageParam) (list []models.SysGeo, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("pid", "=", page.FilterMap["pid"]),asql.Where("deep", "=", page.FilterMap["deep"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysGeoDao) GetById(id int) (row models.SysGeo) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
