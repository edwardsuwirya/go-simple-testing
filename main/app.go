package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosimpletest/services"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/calc", func(c *gin.Context) {
		sNum1 := c.DefaultQuery("num1", "0")
		sNum2 := c.DefaultQuery("num2", "0")
		opr := c.DefaultQuery("opr", "add")

		num1, err := strconv.Atoi(sNum1)
		num2, err := strconv.Atoi(sNum2)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unknown number",
			})
			return
		}
		var result = 0
		switch opr {
		case "add":
			myCalc := services.Calculator{
				num1, num2,
			}
			res, err := myCalc.Add()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": err,
				})
				return
			} else {
				result = res
			}
		case "sub":
			myCalc := services.Calculator{
				num1, num2,
			}
			res, err := myCalc.Sub()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": err,
				})
				return
			} else {
				result = res
			}
		default:
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Unknown operator",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("Can not run server")
	}

}
