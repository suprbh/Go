package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// authorizeDevice is a middleware handler for making sure the request includes
// a valid device and an API key.
func (c connections) authorizeDevice() gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.Request.Header.Get("X-API-Key") == "" {
			g.IndentedJSON(http.StatusUnauthorized,
				ErrorResponse{
					ErrorType:    "APIKeyMissing",
					ErrorMessage: "X-API-Key request header missing"})
			g.Abort()
			return
		}

		deviceID := g.Param("deviceID")
		key := g.Request.Header.Get("X-API-Key")
		ok := c.authorize("device", deviceID, key)
		if !ok {
			g.IndentedJSON(http.StatusUnauthorized,
				ErrorResponse{
					ErrorType:    "APIKeyInvalid",
					ErrorMessage: "Device ID and/or API key are not valid"})
			g.Abort()
		}
	}
}

// authorizeDevice is a middleware handler for making sure the request includes
// a valid device and an API key.
func (c connections) authorizeUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.Request.Header.Get("X-API-Key") == "" {
			g.IndentedJSON(http.StatusUnauthorized,
				ErrorResponse{
					ErrorType:    "APIKeyMissing",
					ErrorMessage: "X-API-Key request header missing"})
			g.Abort()
			return
		}

		deviceID := g.Param("deviceID")
		key := g.Request.Header.Get("X-API-Key")
		ok := c.authorize("device", deviceID, key)
		if !ok {
			g.IndentedJSON(http.StatusUnauthorized,
				ErrorResponse{
					ErrorType:    "APIKeyInvalid",
					ErrorMessage: "Device ID and/or API key are not valid"})
			g.Abort()
		}
	}
}
