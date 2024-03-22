package main

import (
	"math/rand"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type Data struct {
	WaterCon    int16  `json:"water_con"`
	WaterStatus string `json:"water_status"`
	WindCon     int16  `json:"win_con"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	g := gin.Default()

	data := Data{}

	go func() {
		for {
			data.updateStatus()
			time.Sleep(15 * time.Second)
		}
	}()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, data)
	})

	g.Run(":3000")
}

func (d *Data) updateStatus() {
	max := 100
	min := 1

	d.WaterCon = int16(rand.Intn(max-min) + min)
	d.WindCon = int16(rand.Intn(max-min) + min)

	waterStatus := "Aman"
	if d.WaterCon < 5 {
		waterStatus = "Aman"
	} else if d.WaterCon <= 8 {
		waterStatus = "Siaga"
	} else if d.WaterCon > 8 {
		waterStatus = "Bahaya"
	}

	d.WaterStatus = waterStatus

	windStatus := "Aman"
	if d.WindCon < 6 {
		windStatus = "Aman"
	} else if d.WindCon < 15 {
		windStatus = "Siaga"
	} else if d.WindCon >= 15 {
		windStatus = "Bahaya"
	}

	d.WindStatus = windStatus
}