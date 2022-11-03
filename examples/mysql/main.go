package main

import (
	"github.com/cyj19/gowalk"
	"github.com/cyj19/gowalk/component/mysql"
	"github.com/cyj19/gowalk/core/logx"
)

/**
组件MySQL使用例子
*/

type Student struct {
	ID   int64 `gorm:"primary_key;auto_increment"`
	Name string
}

func (s *Student) TableName() string {
	return "student"
}

func main() {
	gowalk.Run(&mysql.Instance{})

	// 根据结构体创建表
	err := mysql.Main().AutoMigrate(&Student{})
	if err != nil {
		logx.Log().Error(err)
	}

	//mysql.Main().Create(&Student{Name: "cyj19"})

	std := &Student{}
	mysql.Main().Where("name = ?", "cyj19").First(std)

	logx.Log().Infof("student: %#v \n", std)

}
