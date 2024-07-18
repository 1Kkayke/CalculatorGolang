package operations

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(server *gin.Context) string {
	num1, err := strconv.Atoi(server.Query("num1"))

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(server.Query("num2"))

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 + num2

	return strconv.Itoa(realResult)
}

func Sub(server *gin.Context) string {
	num1, err := strconv.Atoi(server.Query("num1"))

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(server.Query("num2"))

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 - num2

	return strconv.Itoa(realResult)

}

func Div(server *gin.Context) string {

	num1, err := strconv.Atoi(server.Query("num1"))

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(server.Query("num2"))

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 / num2

	return strconv.Itoa(realResult)

}

func Multi(server *gin.Context) string {

	num1, err := strconv.Atoi(server.Query("num1"))

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(server.Query("num2"))

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 * num2

	return strconv.Itoa(realResult)

}
