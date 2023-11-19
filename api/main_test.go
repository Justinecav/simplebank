package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

//this function is used to see more clear logs
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
