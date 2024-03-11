package api

import (
	"fmt"
	"strconv"
	"sub2sing-box/api/handler"

	"github.com/gin-gonic/gin"
)

func RunServer(port uint16) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/convert", handler.Convert)

	fmt.Println("Server is running on port ", port)
	err := r.Run(":" + strconv.Itoa(int(port)))
	if err != nil {
		fmt.Println("Run server failed: ", err)
	}
}
