package Models

import "backend/Database"

type Cast struct {
	Cast_id          uint `gorm:"primary_key"`
	Cast_name        string
	Cast_description string
	Cast_image       string
}

func AddCast(cast Cast) int64 {
	result := Database.DB.Create(&cast) // 通过数据的指针来创建
	affected := result.RowsAffected     // 返回插入记录的条数
	return affected
}

//update review
func UpdateCast(cast Cast) int64 {
	result := Database.DB.Model(&cast).Updates(&cast)
	//result := Database.DB.Save(&user)
	return result.RowsAffected
}

func DeleteCast(Cast_id int) int64 {
	result := Database.DB.Delete(&Cast{}, Cast_id)
	return result.RowsAffected
}

func SearchCast(castName string) []Cast {
	var cast []Cast
	//Database.DB.First(&cast, castName)
	Database.DB.Where("cast_name LIKE ?", "%"+castName+"%").Find(&cast)
	return cast
}

//
//func SearchMovieByName(name string) []Movie {
//	var movies []Movie
//	Database.DB.Where("movie_name LIKE ?", "%"+name+"%").Find(&movies)
//	return movies
//}
