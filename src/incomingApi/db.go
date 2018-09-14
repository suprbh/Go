package main

import (
	"cassandra"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type connections struct {
	// dbData map[string][]DeviceData
	csndra *cassandra.CassCtx
	// mydb *mysqldb.mctx
	// red *redis.ctx
}

// DeviceProfile has all the information regarding the specific device
type DeviceProfile struct {
	name        string
	deviceID    string
	deviceType  string
	modelNumber string
	lat         float64
	lon         float64
}

// UserProfile contains profile information of the user of the product
type UserProfile struct {
	name        string
	address     string
	memberSince time.Time
	devices     *[]DeviceProfile
}

// DeviceData contains the values of individual data points sent by the device.
type DeviceData struct {
}

// putDevice updates the information about the device TODO
func (c *connections) PutDevice(g *gin.Context) {

}

func (c *connections) PutData(g *gin.Context) {

	deviceID := g.Param("deviceID")

	// Unmarshal and validate the payload.
	var dd DeviceData
	err := g.Bind(&dd)
	if err != nil {
		g.IndentedJSON(http.StatusBadRequest,
			ErrorResponse{
				ErrorType:    "PayloadDecodeError",
				ErrorMessage: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"Data":    dd,
		"Station": deviceID,
	}).Debug("Put Device data")

	// TODO : add this data to the database
	c.csndra.AddDeviceData(dd)
	// c.dbData[deviceID] = append(c.dbData[deviceID], dd)

}

// getDeviceByUser gets the information for the device owned by a user TODO
func (c *connections) GetDevicesByUser(g *gin.Context) {

}

func (c *connections) CreateDevice(g *gin.Context) {

}

func (c *connections) GetDeviceByCreds(g *gin.Context) {

}
