package graphite

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type GraphiteSender struct {
	conn *net.Conn
}

func NewWithConnection(addr string) (*GraphiteSender, error) {
	s := new(GraphiteSender)

	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	s.conn = &c

	return s, nil
}

func (g *GraphiteSender) Send(key []string, time int64, value float32) {
	k := strings.Join(key, ".")
	log.Printf("%v %s %f %d", key, k, value, time)
	_, err := fmt.Fprintf(*(g.conn), "%s %f %d\n", k, value, time)
	if err != nil {
		// retry connection?
		log.Print(err)
		return
	}
}
