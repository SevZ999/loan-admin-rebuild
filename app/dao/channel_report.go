package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type ChannelReportDao struct {
	db     *gorm.DB
	models *models.ChannelReport
}

func NewChannelReportDao(db *gorm.DB) *ChannelReportDao {
	if db == nil {
		db = ant.Db()
	}
	return &ChannelReportDao{db: db}
}

// Create
func (dao *ChannelReportDao) Create(channelReport models.ChannelReport) error {
	return dao.db.Create(&channelReport).Error
}

// DeleteById
func (dao *ChannelReportDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *ChannelReportDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *ChannelReportDao) Update(channelReport models.ChannelReport) error {
	return dao.db.Updates(&channelReport).Error
}

// GetList
func (dao *ChannelReportDao) GetList() (list []models.ChannelReport) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *ChannelReportDao) GetPage(page page.PageParam) (list []models.ChannelReport, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("date", "BETWEEN", page.FilterMap["date"]),
		asql.Where("channel_name", "LIKE", page.FilterMap["channel_name"]),
		asql.Where("channel_key", "LIKE", page.FilterMap["channel_key"]),
		asql.Where("platform", "=", page.FilterMap["platform"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *ChannelReportDao) GetById(id int) (row models.ChannelReport) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
