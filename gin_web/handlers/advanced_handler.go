package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PathParams(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func WildcardParams(c *gin.Context) {
	path := c.Param("filepath")
	c.JSON(http.StatusOK, gin.H{"filepath": path})
}

func QueryArray(c *gin.Context) {
	ids := c.QueryArray("id")
	tags := c.QueryArray("tags")
	c.JSON(http.StatusOK, gin.H{
		"ids":  ids,
		"tags": tags,
	})
}

func QueryMap(c *gin.Context) {
	params := c.QueryMap("params")
	c.JSON(http.StatusOK, gin.H{"params": params})
}

func PostArray(c *gin.Context) {
	items := c.PostFormArray("items")
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func PostMap(c *gin.Context) {
	params := c.PostFormMap("params")
	c.JSON(http.StatusOK, gin.H{"params": params})
}

func ClientIP(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"clientIP":      c.ClientIP(),
		"remoteAddr":    c.Request.RemoteAddr,
		"xForwardedFor": c.GetHeader("X-Forwarded-For"),
	})
}

func RequestInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method":        c.Request.Method,
		"path":          c.Request.URL.Path,
		"query":         c.Request.URL.RawQuery,
		"userAgent":     c.Request.UserAgent(),
		"contentType":   c.ContentType(),
		"contentLength": c.Request.ContentLength,
	})
}

func BindQuery(c *gin.Context) {
	type QueryParams struct {
		Page  int    `form:"page" binding:"required"`
		Limit int    `form:"limit"`
		Sort  string `form:"sort"`
	}

	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, params)
}

func BindForm(c *gin.Context) {
	type FormData struct {
		Name  string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required,email"`
		Age   int    `form:"age"`
	}

	var data FormData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func BindURI(c *gin.Context) {
	type URIParams struct {
		ID   int    `uri:"id" binding:"required"`
		Name string `uri:"name" binding:"required"`
	}

	var params URIParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, params)
}

func BindHeader(c *gin.Context) {
	type HeaderData struct {
		Authorization string `header:"Authorization" binding:"required"`
		ContentType   string `header:"Content-Type"`
	}

	var headers HeaderData
	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, headers)
}
