package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var turnOffTime = time.Now()

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		turnOffTime = time.Now().Add(time.Duration(10 * time.Second))
		c.JSON(http.StatusOK, gin.H{
			"message":     "ok",
			"turnOffTime": turnOffTime,
		})
	})
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "ok",
			"turnOffTime": turnOffTime,
			"status":      turnOffTime.After(time.Now()),
		})
	})

	r.Run("0.0.0.0:8009")
}

func toggleScreen() {
	for range time.Tick(time.Millisecond * 500) {
		if turnOffTime.After(time.Now()) {
			setScreenPower("1")
		} else {
			setScreenPower("0")
		}
	}
}

func setScreenPower(screenPower string) {
	bl_power, err := os.Create("/sys/class/backlight/rpi_backlight/bl_power")
	if err != nil {
		fmt.Print("Could not set backlight: ")
		fmt.Println(err)
	}
	fmt.Fprintf(bl_power, screenPower)
	defer bl_power.Close()
}
