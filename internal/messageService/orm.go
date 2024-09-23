package messageService

type RequestBody struct {
	Id      uint   `gorm:"primary_key"`
	Message string `json:"message"`
}
