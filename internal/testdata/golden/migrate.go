package rdsom

import (
	"context"
	"io/ioutil"
)

const redisNilErr = "redis: nil"

func (c *Client) MigrateAll(ctx context.Context) error {
	return c.eval(ctx, c.migrationScript)
}

func (c *Client) eval(ctx context.Context, luaFile string) error {
	data, err := ioutil.ReadFile(luaFile)
	if err != nil {
		return err
	}
	err = c.rdb.Eval(ctx, string(data), nil).Err()
	if err.Error() != redisNilErr {
		return err
	}
	return nil
}
