package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type ConversionEventDao struct {
	db     *gorm.DB
	models *models.ConversionEvent
}

func NewConversionEventDao(db *gorm.DB) *ConversionEventDao {
	if db == nil {
		db = ant.Db()
	}
	return &ConversionEventDao{db: db}
}

// Create
func (dao *ConversionEventDao) Create(conversionEvent models.ConversionEvent) error {
	return dao.db.Create(&conversionEvent).Error
}

// DeleteById
func (dao *ConversionEventDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *ConversionEventDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *ConversionEventDao) Update(conversionEvent models.ConversionEvent) error {
	return dao.db.Updates(&conversionEvent).Error
}

// GetList
func (dao *ConversionEventDao) GetList() (list []models.ConversionEvent) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *ConversionEventDao) GetPage(page page.PageParam) (list []models.ConversionEvent, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(

		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *ConversionEventDao) GetById(id int) (row models.ConversionEvent) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByChannelCount 获取结算类型的渠道统计数量
func (dao *ConversionEventDao) GetByChannelCount(event_type int, channel string, timeNow string) (count int64) {
	startTime := timeNow + " 00:00:00"
	endTime := timeNow + " 23:59:59"
	dao.db.Model(&dao.models).Where("channel=? AND event_type=? AND callback_url!='' AND callback_status='success' AND event_time BETWEEN ? AND ?", channel, event_type, startTime, endTime).Distinct("oaid").Count(&count)
	return count
}
