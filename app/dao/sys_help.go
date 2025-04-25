package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
	
)

type SysHelpDao struct {
	db    *gorm.DB
	models *models.SysHelp
}

func NewSysHelpDao(db *gorm.DB) *SysHelpDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysHelpDao{db: db}
}

// Create
func (dao *SysHelpDao) Create(sysHelp models.SysHelp) error {
	return dao.db.Create(&sysHelp).Error
}

// DeleteById
func (dao *SysHelpDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysHelpDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysHelpDao) Update(sysHelp models.SysHelp) error {
	return dao.db.Updates(&sysHelp).Error
}

// GetList
func (dao *SysHelpDao) GetList() (list []models.SysHelp) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysHelpDao) GetPage(page page.PageParam) (list []models.SysHelp, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("type", "=", page.FilterMap["type"]),asql.Where("title", "LIKE", page.FilterMap["title"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysHelpDao) GetById(id int) (row models.SysHelp) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
