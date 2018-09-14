package main

import (
	"cassandra"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// test server alive status
func heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func setupRouter() *gin.Engine {

	// HTTP server settings
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	cass, err := cassandra.Open(cfg.Cassandra)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Fatal("Error connecting to Cassandra")
	}
	defer cass.Close()
	c := connections{csndra: &cass}

	// Ping test for heartbeat
	r.GET("/ping", heartbeat)

	// TODO implement each route handler
	v1 := r.Group("/v1")
	{
		deviceAuth := v1.Group("/devices")
		{
			deviceAuth.Use(c.authorizeDevice())
			deviceAuth.PUT("/devices/:deviceID", c.PutDevice)
			dataAuth := deviceAuth.Group("/data")
			{
				dataAuth.PUT("/data/:deviceID", c.PutData)
			}
		}

		userAuth := v1.Group("/users")
		{
			userAuth.Use(c.authorizeUser())
			userAuth.GET("/users/devices/list", c.GetDevicesByUser)
			userAuth.POST("/users/create", c.CreateDevice)
			userAuth.DELETE("/devices/:deviceID", c.DeleteDevice)

		}

	}

	return r

}
