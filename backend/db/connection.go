package db

import (
	"log"

	"github.com/hirakan045/feelog/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 接続に失敗した場合はMySQLを直接操作
// docker exec -i mysql mysql -uroot -p
func Connection() *gorm.DB {
	// データベース接続
	dsn := "root:rootpass@tcp(127.0.0.1:3306)/feelog?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	log.Println("DB接続成功")

	// 全てのテーブル削除
	err = database.Migrator().DropTable(&models.User{}, &models.Diary{}, &models.Chat{})
	if err != nil {
		log.Fatalf("テーブル削除失敗: %v", err)
	}

	log.Println("全テーブル削除成功")

	// 再マイグレーション
	err = database.AutoMigrate(&models.User{}, &models.Diary{}, &models.Chat{})
	if err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	log.Println("マイグレーション成功")

	return database
}
