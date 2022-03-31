package Models

import "backend/Database"

type Watch_list struct {
	User_id  int
	Movie_id int
	User     User  `gorm:"ForeignKey:User_id"`  // 设置关联模型并指定UserId作为外键
	Movie    Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键

}

func AddWatched(Watch_list Watch_list) int64 {
	result := Database.DB.Create(&Watch_list) // 通过数据的指针来创建
	affected := result.RowsAffected           // 返回插入记录的条数
	return affected
}

//func DeleteWatched(user_id int, movie_id int) int64 {
//	var watched = Models.Watch_list{User_id: user_id, Movie_id: movie_id}
//	result:=Database.DB.Delete(&watched)
//	//result := Database.DB.Delete(&Watch_list{}, user_id, movie_id)
//	//result := Database.DB.Where("user_id = ? and movie_id=?", 2, 1).Delete(&User{})
//
//	return result.RowsAffected
//}

func DeleteWatched(watched Watch_list) int64 {

	//user = User{Username: user.Username, Password: user.Password, User_type: 0, Gender: user.Gender, Age: user.Age, Email: user.Email}
	watched = Watch_list{User_id: watched.User_id, Movie_id: watched.Movie_id}
	result := Database.DB.Delete(&watched)
	//result := Database.DB.Delete(&Watch_list{}, user_id, movie_id)
	//result := Database.DB.Where("user_id = ? and movie_id=?", 2, 1).Delete(&User{})

	return result.RowsAffected
}

func ReadWatched(user_id int) []Watch_list {
	var watchList []Watch_list
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Where("user_id = ?", user_id).Find(&watchList)
	return watchList
}
