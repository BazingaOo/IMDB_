package Models

type Genre struct {
	Genre_id   int `gorm:"primary_key"`
	Genre_name string
}
