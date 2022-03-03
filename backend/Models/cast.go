package Models

type Cast struct {
	Cast_id         uint `gorm:"primary_key"`
	Cast_name       string
	Cast_desciption string
}
