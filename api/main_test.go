package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Justinecav/simplebank.git/db/sqlc"
	"github.com/Justinecav/simplebank.git/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)

	require.NoError(t, err)

	return server

}

//this function is used to see more clear logs
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
