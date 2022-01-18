package rdsom

import "github.com/go-redis/redis/v8"

type Client struct {
	rdb             *redis.Client
	migrationScript string

	Bar *barClient
}

func NewClient(opts *redis.Options) *Client {
	rdb := redis.NewClient(opts)

	return &Client{
		rdb:             rdb,
		migrationScript: "lua/migrate.lua",
		Bar:             &barClient{rdb},
	}
}

func (c *Client) Redis() *redis.Client {
	return c.rdb
}

type barClient struct {
	rdb *redis.Client
}

func (c *barClient) Create() *BarCreate {
	return NewBarCreate(c.rdb)
}

func (c *barClient) Find() *BarFind {
	return NewBarFind(c.rdb)
}
