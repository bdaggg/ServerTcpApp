package handlers

import (
	"bufio"
	"io"
	"mytcpapp/pkg/utils"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

/*
func HandleConnection(conn net.Conn) {
 defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				utils.LogMessage("Connection closed by client: " + conn.RemoteAddr().String())
			} else {
				utils.LogMessage("Error reading message: " + err.Error())
			}
			return
		}
		utils.LogMessage("Message received: " + string(message))
		_, err = conn.Write([]byte("Message received.\n"))
		if err != nil {
			utils.LogMessage("Error sending response: " + err.Error())
			return
		}
	}

	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				utils.LogMessage("Connection closed by client: " + conn.RemoteAddr().String())
			} else {
				utils.LogMessage("Error reading message: " + err.Error())
			}
			return
		}

		// Trim the newline character from the message
		message = strings.TrimSpace(message)
		utils.LogMessage("Message received: " + message)

		var response string
		switch message {
		case "ls":
			response, err = executeCommand("ls")
		case "pwd":
			response, err = executeCommand("pwd")
		default:
			response = "Command not recognized."
		}

		if err != nil {
			response = "Error executing command: " + err.Error()
		}

		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			utils.LogMessage("Error sending response: " + err.Error())
			return
		}
	}
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command(command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
*/

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				utils.LogMessage("Connection closed by client: " + conn.RemoteAddr().String())
			} else {
				utils.LogMessage("Error reading message: " + err.Error())
			}
			return
		}

		// Trim the newline character from the message
		message = strings.TrimSpace(message)
		utils.LogMessage("Message received: " + message)

		var response string
		switch message {
		case "ls":
			if runtime.GOOS == "windows" {
				response, err = executeCommand("dir")
			} else {
				response, err = executeCommand("ls")
			}
		case "pwd":
			if runtime.GOOS == "windows" {
				response, err = executeCommand("cd")
			} else {
				response, err = executeCommand("pwd")
			}
		default:
			response = "Command not recognized."
		}

		if err != nil {
			response = "Error executing command: " + err.Error()
		}

		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			utils.LogMessage("Error sending response: " + err.Error())
			return
		}
	}
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command(command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
