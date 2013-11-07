package graphite

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type GraphiteSender struct {
	conn *net.Conn
	addr string
	net  string // "tcp" "udp" etc.

	Verbose bool
}

func NewWithConnection(net, addr string) (*GraphiteSender, error) {
	g := GraphiteSender{addr: addr, net: net}

	err := g.reconnect()
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (g *GraphiteSender) reconnect() error {
	c, err := net.Dial(g.net, g.addr)
	if err != nil {
		return err
	}

	g.conn = &c

	return nil
}

func (g *GraphiteSender) Send(key []string, time int64, value float32) {
	k := strings.Join(key, ".")
	if g.Verbose {
		log.Printf("sending %s %f %d", k, value, time)
	}

	_, err := fmt.Fprintf(*(g.conn), "%s %f %d\n", k, value, time)
	if err != nil {
		// retry connection?
		log.Print(err)
		return
	}
}
