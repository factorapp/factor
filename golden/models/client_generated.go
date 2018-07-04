package models

type Client struct {
	TodoClient
}

func NewClient() (*Client, error) {
	return &Client{

		TodoClient{},
	}, nil
}
