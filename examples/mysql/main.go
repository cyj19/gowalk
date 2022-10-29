package main

import (
	"github.com/cyj19/gowalk"
	"github.com/cyj19/gowalk/component/mysql"
	"log"
)

/**
组件MySQL使用例子
*/

type Student struct {
	ID   int64 `gorm:"primary_key;auto_increment"`
	Name string
}

func main() {
	gowalk.Run(&mysql.Instance{})

	// 根据结构体创建表
	_ = mysql.Main().AutoMigrate(&Student{})

	mysql.Main().Create(&Student{Name: "cyj19"})

	std := &Student{}
	mysql.Main().Where("name = ?").First(std)

	log.Printf("Studen: %+v \n", std)

}
