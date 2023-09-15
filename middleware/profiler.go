package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

func Profiler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Continue handling the request
		err := c.Next()

		// Calculate the time taken for the request
		duration := time.Since(start)

		// Log request details
		log := logrus.New()
		log.WithFields(logrus.Fields{
			"Method":   c.Method(),
			"Path":     c.Path(),
			"Query":    c.Queries(),
			"Duration": duration,
			"Headers":  &c.Request().Header,
			"Body":     string(c.Request().Body()),
		}).Info("Request details")

		// Log SQL queries
		if queries, ok := c.Locals("queries").([]string); ok {
			for _, query := range queries {
				log.WithField("SQL Query", query).Info("SQL Query")
			}
		}

		// Log execution time
		log.WithField("Duration", duration).Info("Request took")

		// Log errors, if any
		if err != nil {
			log.WithError(tracerr.Wrap(err)).Error("Error during request handling")
		}

		data := map[string]interface{}{
			"Method":   c.Method(),
			"Path":     c.Path(),
			"Query":    c.Queries(),
			"Duration": duration,
			"Headers":  &c.Request().Header,
			"Body":     string(c.Request().Body()),
		}

		logsDir := "./storage/logs"
		if _, err := os.Stat(logsDir); os.IsNotExist(err) {
			os.Mkdir(logsDir, 0755)
		}

		filename := fmt.Sprintf("request_%s.json", time.Now().Format("2006-01-02T15-04-05"))

		// Combine logsDir and filename to get the full path
		fullPath := filepath.Join(logsDir, filename)

		logError := SaveRequestDataAsJSON(data, fullPath)
		if logError != nil {
			log.WithError(logError).Error("Error saving request data as JSON")
		}

		return err
	}
}

func SaveRequestDataAsJSON(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}
