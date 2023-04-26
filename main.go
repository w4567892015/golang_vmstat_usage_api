package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		results, err := getMetrics()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.String(200, results)
	})
	router.Run(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	result := fmt.Sprintf(":%s", port)
	return result
}

func getMetrics() (string, error) {
	cmdName := "vmstat"
	cmdArgs := []string{"1", "1"}

	var perfCounters string
	out, err := exec.Command(cmdName, cmdArgs...).Output()

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	parts := strings.Fields(lines[2])
	perfCounters = fmt.Sprintf("%s,%s", time.Now().Format("15:04:05.000"), strings.Join(parts, ","))

	return perfCounters, err
}
