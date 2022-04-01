package Models

import "backend/Database"

type Movie_cast struct {
	Cast_id  uint
	Movie_id uint
	Movie    Movie `gorm:"ForeignKey:Movie_id"` // 设置关联模型并指定UserId作为外键
	Cast     Cast  `gorm:"ForeignKey:Cast_id"`  // 设置关联模型并指定UserId作为外键
}

//// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
//func (Movie_cast) TableName() string {
//	return "movie_cast"
//}

func AddMovieCast(movieCast Movie_cast) int64 {
	result := Database.DB.Create(&movieCast) // 通过数据的指针来创建
	affected := result.RowsAffected          // 返回插入记录的条数
	return affected
}
