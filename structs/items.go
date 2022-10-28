package structs

type Items struct {
	// gorm.Model
	ItemId      int    `gorm:"primary_key;auto_increment;not_null"`
	OrderId     int    `json:"order_id"`
	ItemCode    int    `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
