package controllers

import (
	"assignment_2/structs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrders(c *gin.Context) {
	var (
		orders structs.Orders
		result gin.H
	)

	c.ShouldBindJSON(&orders)

	idb.DB.Debug().Create(&orders)
	for i, item := range orders.Items {
		item.OrderId = orders.OrderId
		idb.DB.Debug().Create(&item)
		orders.Items[i] = item
	}

	// orders.OrderId = orders.OrderId
	result = gin.H{
		"result": orders,
	}

	fmt.Printf("order_id: %d", orders.OrderId)
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		items  []structs.Items
		orders []structs.Orders
		result gin.H
	)

	idb.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		for i, order := range orders {
			idb.DB.Find(&items).Where("order_id = ?", order.OrderId)
			orders[i].Items = items
			fmt.Println(items)
		}
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrders(c *gin.Context) {
	var (
		newOrders structs.Orders
		result    gin.H
	)

	c.ShouldBindJSON(&newOrders)

	id := c.Param("id")
	selectedID, _ := strconv.Atoi(id)
	// fmt.Println(id)
	err := idb.DB.Debug().Model(&newOrders).Where("order_id = ?", selectedID).Updates(&newOrders).Error

	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	} else {
		result = gin.H{
			"result": newOrders,
			// "count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)

}

func (idb *InDB) DeleteOrders(c *gin.Context) {
	var (
		order  structs.Orders
		result gin.H
	)

	id := c.Param("id")
	selectedID, _ := strconv.Atoi(id)
	// fmt.Println(id)
	idb.DB.Debug().Delete(&order, selectedID)
	result = gin.H{
		"result": "Deleted",
		// "count":  len(orders),
	}

	c.JSON(http.StatusOK, result)

}
