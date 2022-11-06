package k

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	EnvNameAMQPEndpoint = "REALTIME_AMQP_ENDPOINT"
	EnvNamePort         = "REALTIME_PORT"
)

const (
	DefaultPort = 80
)

func GetAMQPEndpoint() (string, error) {
	amqpEndpoint := strings.TrimSpace(os.Getenv(EnvNameAMQPEndpoint))
	if amqpEndpoint == "" {
		return "", errors.New("AMQP endpoint is not set")
	}
	return amqpEndpoint, nil
}

func GetPort() int {
	strPort := strings.TrimSpace(os.Getenv(EnvNamePort))
	if strPort == "" {
		return DefaultPort
	}

	intPort, err := strconv.Atoi(strPort)
	if err != nil {
		return DefaultPort
	}

	return intPort
}
