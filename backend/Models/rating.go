package Models

type Rating struct {
	user_id  int
	movie_id int
	score    int
}

func (Rating) TableName() string {
	return "rating"
}
