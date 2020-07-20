package repository

import (
	"db-interaction/common/db"
	"log"
	"os"
)

type BaseRepository struct {
	DB *db.DB
}

// 初始化
func NewRepository()*BaseRepository{
	db := db.DB{}
	err := db.Connect()
	if err != nil {
		log.Panic("初始化数据库错误:",err)
		os.Exit(1)
	}
	return &BaseRepository{
		&db,
	}
}

// 根据条件，获取人员表数据
func (b *BaseRepository) Get(where , out interface{} ,sel string)error{
	// 获取数据库对象
	db := b.DB.Conn.Where(where)
	if sel != "" {
		// 检索的字段
		db = db.Select(sel)
	}
	return db.Find(out).Error
}