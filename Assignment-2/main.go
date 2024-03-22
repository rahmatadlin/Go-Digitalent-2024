package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Item struct {
	ID          uint      `gorm:"primaryKey" json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     uint      `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}

func main() {
	dsn := "host=localhost user=postgres dbname=assignment2 password=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Order{}, &Item{})

	r := gin.Default()

	r.POST("/orders", func(c *gin.Context) {
		var order Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&order)
		c.JSON(200, order)
	})

	r.GET("/orders", func(c *gin.Context) {
		var orders []Order
		db.Preload("Items").Find(&orders)
		c.JSON(200, orders)
	})

	r.PUT("/orders/:orderId", func(c *gin.Context) {
		orderID := c.Param("orderId")
		var order Order
		if err := db.Preload("Items").First(&order, orderID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}
		var updatedOrder Order
		if err := c.BindJSON(&updatedOrder); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Model(&order).Updates(updatedOrder)
		c.JSON(200, order)
	})

	r.DELETE("/orders/:orderId", func(c *gin.Context) {
		orderID := c.Param("orderId")
		var order Order
		if err := db.First(&order, orderID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		db.Delete(&order)
		c.JSON(200, gin.H{"message": "Order deleted successfully"})
	})

	r.Run(":8080")
}