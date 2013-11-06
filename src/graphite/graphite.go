package graphite

type GraphiteSender struct {
	//connection, etc
}

func NewWithConnection() (*GraphiteSender, error) {
	return nil, nil
}

func (g *GraphiteSender) Send(key []string, time int64, value float32) {
}
