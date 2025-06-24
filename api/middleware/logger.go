package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// bodyLogWriter is a custom response writer that captures the response body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// AccessLogger middleware logs non-200 requests and responses
func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Read and store the request body
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// Restore the request body for later use
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Create a custom response writer to capture the response
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:          bytes.NewBufferString(""),
		}
		c.Writer = blw

		// Process request
		c.Next()

		// Only log if status code is not 2xx
		if c.Writer.Status() < 200 || c.Writer.Status() > 299 {
			// Format request data
			requestData := map[string]interface{}{
				"method":     c.Request.Method,
				"path":       c.Request.URL.Path,
				"query":      c.Request.URL.RawQuery,
				"ip":         c.ClientIP(),
				"user_agent": c.Request.UserAgent(),
				"headers":    c.Request.Header,
			}

			// Add request body if it exists and is JSON
			if len(requestBody) > 0 {
				var jsonBody interface{}
				if err := json.Unmarshal(requestBody, &jsonBody); err == nil {
					requestData["body"] = jsonBody
				}
			}

			// Format response data
			responseData := map[string]interface{}{
				"status": c.Writer.Status(),
				"size":   c.Writer.Size(),
			}

			// Add response body if it exists and is JSON
			if blw.body.Len() > 0 {
				var jsonBody interface{}
				if err := json.Unmarshal(blw.body.Bytes(), &jsonBody); err == nil {
					responseData["body"] = jsonBody
				}
			}

			// Create the log entry
			logEntry := map[string]interface{}{
				"timestamp":  time.Now().Format(time.RFC3339),
				"duration":   time.Since(start).String(),
				"request":    requestData,
				"response":   responseData,
				"errors":     c.Errors.Errors(),
			}

			// Convert to JSON and log
			logJSON, _ := json.MarshalIndent(logEntry, "", "  ")
			log.Printf("Access Log:\n%s\n", string(logJSON))
		}
	}
}
