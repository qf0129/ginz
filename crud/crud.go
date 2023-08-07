package crud

import (
	"gorm.io/gorm"
)

const (
	CrudDefaultPageSize   = 10
	CrudDefaultPrimaryKey = "id"
)

type GormModel any

const FixedKeyPage = "Page"
const FixedKeyPageSize = "PageSize"
const FixedKeyPreload = "Preload"
const FixedKeyClosePaging = "ClosePaging"

var FIXED_KEYS = []string{FixedKeyPage, FixedKeyPageSize, FixedKeyPreload, FixedKeyClosePaging}

// 固定查询选项
type FixedOption struct {
	Page        int    // 页数，默认1
	PageSize    int    // 每页数量，默认10
	Preload     string // 预加载关联表名，若多个以英文逗号分隔
	ClosePaging bool   // 关闭分页，默认false
}

type PageBody[T any] struct {
	List     []T
	Page     int
	PageSize int
	Total    int64
}

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}