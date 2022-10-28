package structs

// _ "github.com/jinzhu/gorm"

type Orders struct {
	// gorm.Model
	OrderId       int    `gorm:"primary_key;auto_increment;not_null"`
	Customer_name string `json:"customerName"`
	Items         []Items
	OrderedAt     string
}
