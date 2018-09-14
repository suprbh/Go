package main

import (
	"net/http"

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

	c := connections{
		dbUser:   dbUserCreds,
		dbDevice: dbDeviceCreds,
	}

	// Ping test for heartbeat
	r.GET("/ping", heartbeat)

	// TODO implement each route handler
	v1 := r.Group("/v1")
	{
		// deviceAuth := v1.Group("/devices", gin.BasicAuth(dbDeviceCreds))
		// deviceAuth.PUT("/devices/:deviceID", c.putDevice)
		deviceAuth := v1.Group("/devices")
		{
			deviceAuth.Use(c.authorizeDevice())
			deviceAuth.PUT("/devices/:deviceID", c.PutDevice)
			// deviceAuth.PATCH("/devices/:deviceID", c.UpdateDevice)
			// deviceAuth.DELETE("/devices/:deviceID", c.DeleteDevice)

			dataAuth := deviceAuth.Group("/data")
			{
				dataAuth.PUT("/data/:deviceID", c.PutData)
			}
		}

		// userAuth := v1.Group("/", gin.BasicAuth(dbUserCreds))
		// userAuth.GET("/users/devices/list", c.getDevicesByUser)
		userAuth := v1.Group("/users")
		{
			userAuth.Use(c.authorizeUser())
			userAuth.GET("/users/devices/list", c.GetDevicesByUser)
			userAuth.POST("/users/create", c.CreateDevice)
			userAuth.GET("/list", c.GetDeviceByCreds)

		}

	}

	return r

}
