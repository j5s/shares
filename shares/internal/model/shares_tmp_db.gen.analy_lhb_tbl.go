package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AnalyLhbTblMgr struct {
	*_BaseMgr
}

// AnalyLhbTblMgr open func
func AnalyLhbTblMgr(db *gorm.DB) *_AnalyLhbTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyLhbTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyLhbTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_lhb_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyLhbTblMgr) GetTableName() string {
	return "analy_lhb_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyLhbTblMgr) Reset() *_AnalyLhbTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyLhbTblMgr) Get() (result AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyLhbTblMgr) Gets() (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyLhbTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyLhbTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyLhbTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyLhbTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_AnalyLhbTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyLhbTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithJlr jlr获取 价格
func (obj *_AnalyLhbTblMgr) WithJlr(jlr float64) Option {
	return optionFunc(func(o *options) { o.query["jlr"] = jlr })
}

// WithJgJlr jg_jlr获取 机构净流入(万元)
func (obj *_AnalyLhbTblMgr) WithJgJlr(jgJlr float64) Option {
	return optionFunc(func(o *options) { o.query["jg_jlr"] = jgJlr })
}

// WithPrice price获取 龙虎榜净流入(万元)
func (obj *_AnalyLhbTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyLhbTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyLhbTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyLhbTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyLhbTblMgr) GetByOption(opts ...Option) (result AnalyLhbTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyLhbTblMgr) GetByOptions(opts ...Option) (results []*AnalyLhbTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyLhbTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyLhbTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AnalyLhbTblMgr) GetFromID(id int) (result AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyLhbTblMgr) GetBatchFromID(ids []int) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyLhbTblMgr) GetFromCode(code string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyLhbTblMgr) GetBatchFromCode(codes []string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyLhbTblMgr) GetFromName(name string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyLhbTblMgr) GetBatchFromName(names []string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_AnalyLhbTblMgr) GetFromDay0(day0 int64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_AnalyLhbTblMgr) GetBatchFromDay0(day0s []int64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyLhbTblMgr) GetFromDayStr(dayStr string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyLhbTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromJlr 通过jlr获取内容 价格
func (obj *_AnalyLhbTblMgr) GetFromJlr(jlr float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`jlr` = ?", jlr).Find(&results).Error

	return
}

// GetBatchFromJlr 批量查找 价格
func (obj *_AnalyLhbTblMgr) GetBatchFromJlr(jlrs []float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`jlr` IN (?)", jlrs).Find(&results).Error

	return
}

// GetFromJgJlr 通过jg_jlr获取内容 机构净流入(万元)
func (obj *_AnalyLhbTblMgr) GetFromJgJlr(jgJlr float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`jg_jlr` = ?", jgJlr).Find(&results).Error

	return
}

// GetBatchFromJgJlr 批量查找 机构净流入(万元)
func (obj *_AnalyLhbTblMgr) GetBatchFromJgJlr(jgJlrs []float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`jg_jlr` IN (?)", jgJlrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 龙虎榜净流入(万元)
func (obj *_AnalyLhbTblMgr) GetFromPrice(price float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 龙虎榜净流入(万元)
func (obj *_AnalyLhbTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyLhbTblMgr) GetFromNewPrice(newPrice float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyLhbTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyLhbTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyLhbTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyLhbTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyLhbTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyLhbTblMgr) FetchByPrimaryKey(id int) (result AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_AnalyLhbTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result AnalyLhbTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyLhbTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
