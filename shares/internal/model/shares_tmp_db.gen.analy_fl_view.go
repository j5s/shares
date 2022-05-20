package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AnalyFlViewMgr struct {
	*_BaseMgr
}

// AnalyFlViewMgr open func
func AnalyFlViewMgr(db *gorm.DB) *_AnalyFlViewMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyFlViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyFlViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_fl_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyFlViewMgr) GetTableName() string {
	return "analy_fl_view"
}

// Reset 重置gorm会话
func (obj *_AnalyFlViewMgr) Reset() *_AnalyFlViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyFlViewMgr) Get() (result AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyFlViewMgr) Gets() (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyFlViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyFlViewMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyFlViewMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyFlViewMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUpName up_name获取 行业名
func (obj *_AnalyFlViewMgr) WithUpName(upName string) Option {
	return optionFunc(func(o *options) { o.query["up_name"] = upName })
}

// WithDay day获取 入池时间
func (obj *_AnalyFlViewMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyFlViewMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyFlViewMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyFlViewMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyFlViewMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyFlViewMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithScore score获取 得分
func (obj *_AnalyFlViewMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyFlViewMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithNum num获取 总家数
func (obj *_AnalyFlViewMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithUp up获取 上涨家数
func (obj *_AnalyFlViewMgr) WithUp(up int) Option {
	return optionFunc(func(o *options) { o.query["up"] = up })
}

// WithHyName hy_name获取 行业名
func (obj *_AnalyFlViewMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyFlViewMgr) GetByOption(opts ...Option) (result AnalyFlView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyFlViewMgr) GetByOptions(opts ...Option) (results []*AnalyFlView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyFlViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyFlView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where(options.query)
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
func (obj *_AnalyFlViewMgr) GetFromID(id int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyFlViewMgr) GetBatchFromID(ids []int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyFlViewMgr) GetFromCode(code string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyFlViewMgr) GetBatchFromCode(codes []string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyFlViewMgr) GetFromName(name string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyFlViewMgr) GetBatchFromName(names []string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUpName 通过up_name获取内容 行业名
func (obj *_AnalyFlViewMgr) GetFromUpName(upName string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`up_name` = ?", upName).Find(&results).Error

	return
}

// GetBatchFromUpName 批量查找 行业名
func (obj *_AnalyFlViewMgr) GetBatchFromUpName(upNames []string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`up_name` IN (?)", upNames).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_AnalyFlViewMgr) GetFromDay(day datatypes.Date) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_AnalyFlViewMgr) GetBatchFromDay(days []datatypes.Date) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyFlViewMgr) GetFromDayStr(dayStr string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyFlViewMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyFlViewMgr) GetFromPrice(price float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyFlViewMgr) GetBatchFromPrice(prices []float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyFlViewMgr) GetFromNewPrice(newPrice float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyFlViewMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyFlViewMgr) GetFromPercent(percent float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyFlViewMgr) GetBatchFromPercent(percents []float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyFlViewMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyFlViewMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_AnalyFlViewMgr) GetFromScore(score int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_AnalyFlViewMgr) GetBatchFromScore(scores []int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyFlViewMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyFlViewMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 总家数
func (obj *_AnalyFlViewMgr) GetFromNum(num int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 总家数
func (obj *_AnalyFlViewMgr) GetBatchFromNum(nums []int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromUp 通过up获取内容 上涨家数
func (obj *_AnalyFlViewMgr) GetFromUp(up int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`up` = ?", up).Find(&results).Error

	return
}

// GetBatchFromUp 批量查找 上涨家数
func (obj *_AnalyFlViewMgr) GetBatchFromUp(ups []int) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`up` IN (?)", ups).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_AnalyFlViewMgr) GetFromHyName(hyName string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_AnalyFlViewMgr) GetBatchFromHyName(hyNames []string) (results []*AnalyFlView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlView{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
