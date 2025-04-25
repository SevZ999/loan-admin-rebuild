package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type ChannelNewUserReportDao struct {
	db     *gorm.DB
	models *models.ChannelNewUserReport
}

func NewChannelNewUserReportDao(db *gorm.DB) *ChannelNewUserReportDao {
	if db == nil {
		db = ant.Db()
	}
	return &ChannelNewUserReportDao{db: db}
}

// Create
func (dao *ChannelNewUserReportDao) Create(channelNewUserReport models.ChannelNewUserReport) error {
	return dao.db.Create(&channelNewUserReport).Error
}

// DeleteById
func (dao *ChannelNewUserReportDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *ChannelNewUserReportDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *ChannelNewUserReportDao) Update(channelNewUserReport models.ChannelNewUserReport) error {
	return dao.db.Updates(&channelNewUserReport).Error
}

// GetList
func (dao *ChannelNewUserReportDao) GetList() (list []models.ChannelNewUserReport) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *ChannelNewUserReportDao) GetPage(page page.PageParam) (list []models.ChannelNewUserReport, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(

		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *ChannelNewUserReportDao) GetById(id int) (row models.ChannelNewUserReport) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetById
func (dao *ChannelNewUserReportDao) GetByDateAndChannelKey(date, channel_key string) (row models.ChannelNewUserReport) {
	dao.db.Model(&dao.models).Where("date=? AND channel_key=?", date, channel_key).Limit(1).Find(&row)
	return row
}
