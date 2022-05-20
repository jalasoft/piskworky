package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strings"
)

func indexHandler(c *gin.Context) {
	bytes, err := res.ReadFile("resources/index.html");
	
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.Data(http.StatusOK, "text/html", bytes)
	}
}

func resourceHandler(c *gin.Context) {
	rest := c.Param("rest")
	resourcePath := fmt.Sprintf("resources%s", rest)

	bytes, err := res.ReadFile(resourcePath)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		contentType := resolveContentType(resourcePath)

		if contentType == "" {
			c.String(http.StatusOK, fmt.Sprintf("Not existing resource '%s'", resourcePath))
		} else {
			c.Data(http.StatusOK, contentType, bytes)
		}
	}
}

func resolveContentType(path string) string {
	if strings.Contains(path, "/js/") {
		return "text/javascript"
	} else if strings.Contains(path, "/css/") {
		return "text/css"
	} 
	return "text/raw"
}