package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type ChannelDao struct {
	db     *gorm.DB
	models *models.Channel
}

func NewChannelDao(db *gorm.DB) *ChannelDao {
	if db == nil {
		db = ant.Db()
	}
	return &ChannelDao{db: db}
}

// Create
func (dao *ChannelDao) Create(channel models.Channel) error {
	return dao.db.Create(&channel).Error
}

// DeleteById
func (dao *ChannelDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *ChannelDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *ChannelDao) Update(channel models.Channel) error {
	return dao.db.Updates(&channel).Error
}

// GetList
func (dao *ChannelDao) GetList() (list []models.Channel) {
	dao.db.Model(&dao.models).Where("status=2 AND channel_type=2 AND settlement_type!=''").Find(&list)
	return list
}

// GetPage
func (dao *ChannelDao) GetPage(page page.PageParam) (list []models.Channel, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("channel_code", "LIKE", page.FilterMap["channel_code"]),
		asql.Where("channel_name", "LIKE", page.FilterMap["channel_name"]),
		asql.Where("channel_type", "=", page.FilterMap["channel_type"]),
		asql.Where("landing_page_type", "=", page.FilterMap["landing_page_type"]),
		asql.Where("status", "=", page.FilterMap["status"]),
		asql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *ChannelDao) GetById(id int) (row models.Channel) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByChannelCode
func (dao *ChannelDao) GetByChannelCode(channel_code string) (row models.Channel) {
	dao.db.Model(&dao.models).Where("channel_code=?", channel_code).Limit(1).Find(&row)
	return row
}
