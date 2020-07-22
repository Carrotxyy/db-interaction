package repository

import (
	"bytes"
	"db-interaction/common/db"
	"db-interaction/models"
	"fmt"
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

// 保存数据
func (b *BaseRepository) Save(value interface{})error{
	return b.DB.Conn.Create(value).Error
}

// 批量保存数据
func (b *BaseRepository) BatchSave(value []*models.Visitor)error{
	var buffer bytes.Buffer
	sql := "insert into `go_visitor` (`vis_name`,`vis_number`,`vis_uname`,`vis_unumber`,`vis_idtype`,`vis_idnum`,`vis_starttime`,`vis_endtime`,`vis_message`,`vis_isacc`,`vis_state`,`vis_filename`) values "

	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	// 构建批量添加的sql
	for i, e := range value {
		if i == len(value)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s',%t,%t,'%s');", e.Name, e.Number, e.UName ,e.UNumber,e.IdType,e.IdNum,e.StartTime,e.EndTime,e.Message,e.IsAcc,e.State,e.FileName))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s',%t,%t,'%s'),", e.Name, e.Number, e.UName ,e.UNumber,e.IdType,e.IdNum,e.StartTime,e.EndTime,e.Message,e.IsAcc,e.State,e.FileName))
		}
	}

	// 执行sql
	return b.DB.Conn.Exec(buffer.String()).Error

}

// 删除表中所有数据
func (b *BaseRepository) TruncateTable(tableName string)error{
	sql := "truncate table " + tableName
	// 执行sql
	return b.DB.Conn.Exec(sql).Error
}