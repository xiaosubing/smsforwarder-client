package models

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"smsforwarder/conf"
	"strings"
)

var DB *gorm.DB

func NewDB() {
	db, err := gorm.Open(sqlite.Open(conf.Smsforwarder.Db.DbName), &gorm.Config{

		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				//SlowThreshold: time.Second, // 慢查询阈值
				LogLevel: logger.Error, // 日志级别
				Colorful: true,         // 彩色日志
				//LogLevel: logger.Silent,
			},
		),
	})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
		return
	}

	// 自动键表
	err = db.AutoMigrate(&Message{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	DB = db
}

// InsertData 通用插入函数
func InsertData[T any](data T) error {
	result := DB.Create(&data)

	if result.Error != nil {
		panic("插入失败: " + result.Error.Error())
	}

	return nil
}

func QueryData(model interface{}, params QueryParams) error {
	query := DB.Model(model)

	if params.Keyword != "" {
		query = query.Where(params.Keyword)
	}

	query = query.Order("id DESC")

	if params.PageSize > 0 {
		query = query.Limit(params.PageSize)
	}

	if err := query.Find(model).Error; err != nil {
		return err
	}

	return nil
}

// UpdateData 通用更新函数（带类型推断）
func UpdateData[T any](model interface{}, condition interface{}, data interface{}, out *int64) error {
	result := DB.Model(model).Where(condition).Updates(data)

	if result.Error != nil {
		return fmt.Errorf("更新失败: %w", result.Error)
	}

	*out = result.RowsAffected
	if *out == 0 {
		return fmt.Errorf("未找到匹配数据")
	}

	//fmt.Printf("成功更新%d条记录\n", *out)
	return nil
}

// DeleteData 通用删除函数（带安全验证）
func DeleteData(model interface{}, condition interface{}, out *int64) error {
	result := DB.Where(condition).Delete(&model)

	if result.Error != nil {
		return fmt.Errorf("删除失败: %w", result.Error)
	}

	*out = result.RowsAffected
	if *out == 0 {
		return fmt.Errorf("未找到匹配数据")
	}

	if isFullTableDelete(condition) {
		return errors.New("禁止全表删除操作")
	}

	//fmt.Printf("成功删除%d条记录\n", *out)
	return nil
}

// 全表删除检测函数
func isFullTableDelete(condition interface{}) bool {
	switch cond := condition.(type) {
	case string:
		return strings.Contains(strings.ToLower(cond), "1=1")
	case []interface{}:
		return len(cond) == 0
	default:
		return false
	}
}
