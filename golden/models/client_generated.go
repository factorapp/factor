package models

type Client struct {
	TodoClient
}

func NewClient(port int) (*Client, error) {
	return &Client{

		TodoClient{},
	}, nil
}
