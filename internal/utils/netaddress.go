package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n *NetAddress) Set(flagValue string) error {
	var err error
	values := strings.Split(flagValue, ":")
	n.Host = values[0]
	n.Port, err = strconv.Atoi(values[1])
	if err != nil {
		return err
	}
	return nil
}
