package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, gin.H{"message": "Cookie set"})
}

func GetCookie(c *gin.Context) {
	cookie, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"cookie": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cookie": cookie})
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Cookie deleted"})
}

func MultipleCookies(c *gin.Context) {
	c.SetCookie("user", "alice", 3600, "/", "", false, true)
	c.SetCookie("theme", "dark", 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Multiple cookies set"})
}
