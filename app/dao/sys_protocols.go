package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
	
)

type SysProtocolsDao struct {
	db    *gorm.DB
	models *models.SysProtocols
}

func NewSysProtocolsDao(db *gorm.DB) *SysProtocolsDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysProtocolsDao{db: db}
}

// Create
func (dao *SysProtocolsDao) Create(sysProtocols models.SysProtocols) error {
	return dao.db.Create(&sysProtocols).Error
}

// DeleteById
func (dao *SysProtocolsDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysProtocolsDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysProtocolsDao) Update(sysProtocols models.SysProtocols) error {
	return dao.db.Updates(&sysProtocols).Error
}

// GetList
func (dao *SysProtocolsDao) GetList() (list []models.SysProtocols) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysProtocolsDao) GetPage(page page.PageParam) (list []models.SysProtocols, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysProtocolsDao) GetById(id int) (row models.SysProtocols) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
