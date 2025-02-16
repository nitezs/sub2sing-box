package api

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nitezs/sub2sing-box/api/handler"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFiles embed.FS

func RunServer(bind string, port uint16) {
	tpl, err := template.ParseFS(staticFiles, "static/*")
	if err != nil {
		println(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.SetHTMLTemplate(tpl)

	r.GET(
		"/static/*path", func(c *gin.Context) {
			c.FileFromFS("static/"+c.Param("path"), http.FS(staticFiles))
		},
	)

	r.GET(
		"/", func(c *gin.Context) {
			c.HTML(
				200, "index.html", nil,
			)
		},
	)

	r.GET("/convert", handler.Convert)

	address := bind + ":" + strconv.Itoa(int(port))
	fmt.Println("Server is running on", address)
	err = r.Run(address)
	if err != nil {
		fmt.Println("Run server failed: ", err)
	}
}
