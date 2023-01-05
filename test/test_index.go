package test

import (
	"fmt"
	"ginchat/models"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Test() {
	// ------- 迁移 schema -------
	db.AutoMigrate(&models.UserBasic{})

	// ------- Create -------
	// user := &models.UserBasic{}
	// user.Name = "LDL"
	// db.Create(user)

	// user2 := &models.UserBasic{}
	// user2.Name = "asdfasdf"
	// db.Create(user2)

	// ------- Read -------
	// user := models.UserBasic{}
	// db.First(&user, 1) // 根据整型主键查找
	// db.Where("Name = ?", "LDL").First(&user)
	// fmt.Println(user)

	users := []models.UserBasic{}
	db.Model(&models.UserBasic{}).Where("").Find(&users)
	fmt.Println(users)
	// db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// ------- Update -------
	// user.IsLogout = true
	// db.Save(&user)
	// Update - 更新多个字段
	// db.Model(&models.UserBasic{}).Where("ID").Updates(models.UserBasic{ClientIP: "200.0.0.1", Email: "F432@qq.com"}) // 仅更新非零值字段
	// db.Model(&models.UserBasic{}).Where("Name = ?", "LDL").Update("is_logout", false)
	// db.Model(&models.UserBasic{}).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// ------- Delete -------
	// db.Delete(&models.UserBasic{}, 3)
}
