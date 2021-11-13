package client

type Client struct {
	port int
}

func MakeClient() *Client {
	newClient := Client{port: 1}
	return &newClient
}
func (c *Client) SetPort(port int) {
	c.port = port
}
