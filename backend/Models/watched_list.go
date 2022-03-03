package Models

import "backend/Database"

type Watched_list struct {
	User_id  int
	Movie_id int
	User     User  `gorm:"ForeignKey:User_id"`  // 设置关联模型并指定UserId作为外键
	Movie    Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键

}

func AddWatched(watched_list Watched_list) int64 {
	result := Database.DB.Create(&watched_list) // 通过数据的指针来创建
	affected := result.RowsAffected             // 返回插入记录的条数
	return affected
}

func DeleteWatched(user_id int, movie_id int) int64 {
	result := Database.DB.Delete(&Watched_list{}, user_id, movie_id)
	return result.RowsAffected
}

func ReadWatched(user_id int) []Watched_list {
	var watchList []Watched_list
	//Database.DB.Where("Movie_name Like ? , "%jin%").Find(&movies)
	Database.DB.Where("user_id = ?", user_id).Find(&watchList)
	return watchList
}
