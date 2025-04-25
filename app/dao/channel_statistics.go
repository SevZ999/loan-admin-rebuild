package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type ChannelStatisticsDao struct {
	db     *gorm.DB
	models *models.ChannelStatistics
}

func NewChannelStatisticsDao(db *gorm.DB) *ChannelStatisticsDao {
	if db == nil {
		db = ant.Db()
	}
	return &ChannelStatisticsDao{db: db}
}

// Create
func (dao *ChannelStatisticsDao) Create(channelStatistics models.ChannelStatistics) error {
	return dao.db.Create(&channelStatistics).Error
}

// DeleteById
func (dao *ChannelStatisticsDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *ChannelStatisticsDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *ChannelStatisticsDao) Update(channelStatistics models.ChannelStatistics) error {
	return dao.db.Updates(&channelStatistics).Error
}

// GetList
func (dao *ChannelStatisticsDao) GetList() (list []models.ChannelStatistics) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *ChannelStatisticsDao) GetPage(page page.PageParam) (list []models.ChannelStatistics, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("channel_name", "LIKE", page.FilterMap["channel_name"]),
		asql.Where("channel_key", "LIKE", page.FilterMap["channel_key"]),
		asql.Where("type", "=", page.FilterMap["type"]),
		asql.Where("date", "BETWEEN", page.FilterMap["date"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *ChannelStatisticsDao) GetById(id int) (row models.ChannelStatistics) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
