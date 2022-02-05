package Services

import (
	Database "backend/Database" // 引入数据库全局变量所在的包
	Models "backend/Models"
	"fmt"
)

//add data
func AddUser(user Models.User) {
	result := Database.Db.Create(&user) // pass pointer of data to Create
	fmt.Println(user.User_id)           // returns inserted data's primary key
	fmt.Println(result.Error)           // returns error
	fmt.Println(result.RowsAffected)    // returns inserted records count

}

//delete
func DeleteUser() {
	Database.Db.Delete(&Models.User{}, 111)
}

//update

func UpdateUser() {
	Database.Db.Model(&Models.User{}).Update("user_name", "hello")
}

// 查询单条数据
func Signin(iUser Models.User) int {
	var user = new(Models.User)
	Database.Db.Where("username = ? AND password = ?", iUser.Username, iUser.Password).Find(&user)
	var a = 0
	if user.Username == "" {
		a = 0
	} else {
		a = 1
	}
	fmt.Println(a)
	return a
}

// 查询列表数据
func GetList(page int, pageSize int) []Models.User {
	var user []Models.User
	Database.Db.Where("status = ? AND user_id = ?", "1", "1").Preload("User").Limit(pageSize).Offset((page - 1) * pageSize).Order("create_time desc").Find(&user)
	return user
}
