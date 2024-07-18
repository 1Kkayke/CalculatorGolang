package main

import (
	"CalculatorRest/packages/operations"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calculator", initCalculator)
	r.Run()
}

func initCalculator(c *gin.Context) {

	if c.Query("operation") == "" {
		c.JSON(400, gin.H{"error": "Invalid Operation"})
		return
	}

	switch c.Query("operation") {
	case "add":
		result := operations.Add(c)
		c.JSON(200, gin.H{"result": result})
	case "sub":
		result := operations.Sub(c)
		c.JSON(200, gin.H{"result": result})
	case "div":
		result := operations.Div(c)
		c.JSON(200, gin.H{"result": result})
	case "multi":
		result := operations.Multi(c)
		c.JSON(200, gin.H{"result": result})

	}
}
