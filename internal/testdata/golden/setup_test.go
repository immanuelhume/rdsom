package rdsom_test

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
	rdsom "github.com/immanuelhume/rdsomgolden"
)

// Create an rdsom.Client for use in all test cases.
var _client = rdsom.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// Time of codegen to emulate.
const _codegenAt uint64 = 1642524233

func TestMain(m *testing.M) {
	ctx := context.Background()
	// Flush everything before tests run.
	err := _client.Redis().FlushAll(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}
