package handler

import (
	"encoding/json"
	"sub2sing-box/api/model"
	iutil "sub2sing-box/internal/util"
	putil "sub2sing-box/pkg/util"

	"github.com/gin-gonic/gin"
)

func Convert(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if c.Query("data") == "" {
		c.JSON(400, gin.H{
			"error": "Missing data parameter",
		})
		return
	}
	j, err := iutil.DecodeBase64(c.Query("data"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid data",
		})
		return
	}
	var data model.ConvertRequest
	err = json.Unmarshal([]byte(j), &data)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if data.Proxies == nil && data.Subscriptions == nil {
		c.JSON(400, gin.H{
			"error": "Must provide at least one subscription or proxy",
		})
		return
	}
	result, err := putil.Convert(data.Subscriptions, data.Proxies, data.Template, data.Delete, data.Rename)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(200, result)
}
