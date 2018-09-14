package main

import "github.com/gin-gonic/gin"

type connections struct {
	dbUser   map[string]string
	dbDevice map[string]string
	// cass *cassandra.cctx  // TODO
}

// DB : Temporarily using a map instead of actual Redis/Cassandra/MySql/MongoDB database
// dummy credentials
var dbUserCreds = map[string]string{
	"foo": "bar",
	"sup": "test",
}

var dbDeviceCreds = map[string]string{
	"apiKey1": "device1",
	"apiKey2": " device2",
}

// TODO: This function authorizes without using gin BasicAuth,
// It takes in name: device or user,
// Corresponding key and value pair for authorization
func (c *connections) authorize(name, key, value string) bool {
	return true
}

// putDevice updates the information about the device TODO
func (c *connections) putDevice(g *gin.Context) {

}

// getDeviceByUser gets the information for the device owned by a user TODO
func (c *connections) getDeviceByUser() {

}
