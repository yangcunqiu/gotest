package main

type GroupResult struct {
	Age   int
	Count int
}

type APIUserInfo struct {
	Name string
}

func main() {
	// 连接数据库
	//db, err := gorm.Open(mysql.Open("test:pwd1234@tcp(172.31.54.123)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	//迁移 schema (保持结构体和表结构一致, gorm不会删除列)
	//db.AutoMigrate(&model.UserInfo{})
	// 创建
	//db.Debug().Create(&model.UserInfo{Name: "test1", Age: 12, PhoneNumber: "123"})

	// 查询
	//u := model.UserInfo{}
	//us := make([]model.UserInfo, 20)
	//var tx *gorm.DB
	// 查第一个
	//tx = db.Debug().First(&u)

	// 根据id查询
	// First会加limit 1
	//tx = db.Debug().First(&u, 15)

	// in
	//tx = db.Debug().Find(&us, []int{1, 2, 3})

	// where
	// 查一个
	//tx = db.Debug().Where("name = ?", "ycq2").First(&u)
	// 查多个
	//tx = db.Debug().Where("name = ?", "ycq2").Find(&us)
	// like
	//tx = db.Debug().Where("name like ?", "%zml%").Find(&us)
	// 多条件
	//tx = db.Debug().Where("name = ? and age = ?", "ycq2", 10).Find(&us)
	// struct条件 (零值不会作为条件)
	//queryUser := model.UserInfo{Name: "ycq2", Age: 0}
	//tx = db.Debug().Where(&queryUser).Find(&us)
	// map条件
	//queryMap := make(map[string]interface{})
	//queryMap["name"] = "ycq2"
	//queryMap["age"] = 0
	//tx = db.Debug().Where(queryMap).Find(&us)
	// not
	//tx = db.Debug().Not("name = ?", "ycq2").Find(&us)
	//tx = db.Debug().Where("name != ?", "ycq2").Find(&us)
	// or
	//tx = db.Debug().Where("name = ?", "zml1").Or("name = ?", "zml2").Find(&us)
	// select
	//tx = db.Debug().Select("name, age").Where("name = ?", "ycq2").Find(&us)
	// order
	//tx = db.Debug().Order("age desc").Where("name = ?", "ycq2").Find(&us)
	// limit
	//tx = db.Debug().Limit(2).Find(&us)
	// offset limit = pageSize; offset = (pageNum-1)*pageSize
	//tx = db.Debug().Limit(2).Offset(2).Find(&us)
	// group
	//groupResult := make([]GroupResult, 20)
	//tx = db.Debug().Model(&u).Select("age, count(age) as count").Group("age").Having("age > ?", 0).Find(&groupResult)
	// 智能选择字段
	//apiUserInfo := APIUserInfo{}
	//tx = db.Debug().Model(&u).Where("name = ?", "zml1").Find(&apiUserInfo)
	//
	//fmt.Println("---------------------result-----------------------")
	//fmt.Printf("user: \n%v\ncount=%v\n", u, tx.RowsAffected)
	//fmt.Println("")
	//fmt.Printf("userList: \n%v\ncount=%v\n", us, tx.RowsAffected)
	//fmt.Println("")
	//fmt.Printf("groupResult: \n%v\n", groupResult)

	// update
	// 更新所有字段
	//db.First(&u)
	//u.Name = "改"
	//db.Debug().Save(&u)
	// 更新单个
	//db.Debug().Model(&u).Where("name = ?", "改").Update("name", "gai")
	// 根据id更新
	//u.ID = 1
	//db.Debug().Model(&u).Update("name", "gai by id")
	// 更新多个字段
	//u.ID = 1
	//db.Debug().Model(&u).Updates(&model.UserInfo{Name: "update", Age: 30})
	// 指定字段
	//u.ID = 1
	//u.Name = "改"
	//u.Age = 99
	//db.Debug().Model(&u).Select("age").Updates(&u)
	// 忽略字段
	//u.ID = 2
	//u.Name = "改"
	//u.Age = 99
	//db.Debug().Model(&u).Omit("age").Updates(&u)
	// 批量更新
	//db.Debug().Model(&u).Where("age = ?", 0).Updates(&model.UserInfo{Name: "---"})

	// delete
	// 主键删除
	//db.Debug().Delete(&u, 2)
	// 指定条件
	//u.ID = 1
	//db.Debug().Delete(&u)
	// 额外条件删除
	//db.Debug().Where("name = ?", "---").Delete(&u)
}
