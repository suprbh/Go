package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// cfg points to the configuration parameters from yaml file
var cfg config

// ErrorResponse is the payload for errored requests.
type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorType    string `json:"errorType,omitempty"`
}

// CreateSuccessResponse is the payload returned after creating
// a new Device.
type CreateSuccessResponse struct {
	DeviceID string `json:"deviceId"`
	APIKey   string `json:"apiKey"`
}

// TODO add testing framework
func main() {

	configFile := flag.String("config", "config/config.yaml", "configuration file")
	// debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	// Load global configuration from file
	err := cfg.loadConfig(*configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"ConfigFile": *configFile,
			"Error":      err,
		}).Fatal("Unable to read configuration file")
	}

	r := setupRouter()

	// Check for PORT environment variable config override.
	p := cfg.Server.Port
	if s := os.Getenv("PORT"); len(s) > 0 {
		p, _ = strconv.Atoi(s)
	}

	// TODO fix path for logrus. Using fmt temporarily
	// log.WithFields(log.Fields{
	// 	"Address": cfg.Server.Address,
	// 	"Port":    p,
	// }).Debug("Starting HTTP server")

	fmt.Println("Starting HTTP Server")

	// Listen and Serve at server address and port specified in config file
	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Address, p))
}
