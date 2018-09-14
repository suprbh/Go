package main

import (
	"cassandra"
	"math/rand"
	"net/http"
	"redis"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type connections struct {
	// dbData map[string][]DeviceData
	csndra *cassandra.CassCtx
	red    *redis.RedisCtx
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
// This is too general. Need to be specific about the device
type DeviceData struct {
	tmpData1 string
	tmpData2 string
	tmpData3 int
	tmpData4 float32
	tmpData5 float32
}

// randomKey generates a random key of length n.
func randomKey(n int) string {
	rand.Seed(time.Now().UnixNano())

	// Note that the letters i, l, and o are missing.
	var letters = []rune("abcdefghjkmnpqrstuvwxyz0123456789")

	key := make([]rune, n)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}

	return string(key)
}

// putDevice inserts/updates the profile information about the device
func (c *connections) PutDevice(g *gin.Context) {
	// deviceID has already been validated in Auth section, so the deviceID exists and
	// is valid.
	deviceID := g.Param("deviceID")

	var dp DeviceProfile
	err := g.Bind(&dp)
	if err != nil {
		g.IndentedJSON(http.StatusBadRequest,
			ErrorResponse{
				ErrorType:    "PayloadDecodeError",
				ErrorMessage: err.Error()})
		return
	}

	log.WithFields(log.Fields{
		"Data":   dp,
		"Device": deviceID,
	}).Debug("Put Device profile information")

	c.csndra.InsertUpdateDeviceProfile(dp)

}

// PutData ingests incoming device data and saves it in the database
func (c *connections) PutData(g *gin.Context) {

	// deviceID has already been validated in Auth section, so the deviceID exists and
	// is valid.
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
		"Data":   dd,
		"Device": deviceID,
	}).Debug("Put Device data")

	c.csndra.AddDeviceData(dd)

}

// getDeviceByUser gets the information for the device owned by a user
func (c *connections) GetDevicesByUser(g *gin.Context) []string {
	//username is already validated by Authorization step
	username := g.Request.Header.Get("username")

	devices, err := c.csndra.FindDevicesByUsername(username)
	if err != nil {
		g.IndentedJSON(http.StatusNotFound,
			ErrorResponse{
				ErrorType:    "DatabaseError",
				ErrorMessage: err.Error()})
		return
	}

	// list of DeviceIDs
	return devices

}

// CreateDevice : add a new device to the list of devices in the database, for the user
// assuming 1-1 relation between a device and a User.
func (c *connections) CreateDevice(g *gin.Context) {
	// validate device
	var dp DeviceProfile
	err := g.Bind(&dp)
	if err != nil {
		g.IndentedJSON(http.StatusBadRequest,
			ErrorResponse{
				ErrorType:    "PayloadDecodeError",
				ErrorMessage: err.Error()})
		return
	}

	// Validate description data.
	if (dp.lat < -90.0) || (dp.lat > 90.0) || (dp.lon < -180.0) || (dp.lon > 180.0) {
		g.IndentedJSON(http.StatusBadRequest,
			ErrorResponse{
				ErrorType:    "LocationError",
				ErrorMessage: "Missing or invalid latitude/longitude"})
		return
	}

	cdp := cassandra.DeviceProfile{
		Name:        dp.name,
		DeviceID:    dp.deviceID,
		DeviceType:  dp.deviceType,
		ModelNumber: dp.modelNumber,
		Lat:         dp.lat,
		Lon:         dp.lon,
	}
	// Generate a random API Key for the device
	k := randomKey(8)

	err = c.csndra.AddDeviceData(&cdp)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Error adding Device")
		g.IndentedJSON(http.StatusInternalServerError,
			ErrorResponse{
				ErrorType:    "DeviceCreationError",
				ErrorMessage: err.Error()})
	} else {
		g.IndentedJSON(http.StatusCreated,
			CreateSuccessResponse{
				DeviceID: cdp.DeviceName,
				APIKey:   k})
	}

}

// DeleteDevice : remove a device from the database
func (c *connections) DeleteDevice(g *gin.Context) {
	deviceID := g.Param("deviceID")

	err := c.csndra.DeleteDevice(deviceID)

	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Error("Error deleting device")
		g.IndentedJSON(http.StatusInternalServerError,
			ErrorResponse{
				ErrorType:    "DeviceDeleteError",
				ErrorMessage: err.Error()})
	}
}
