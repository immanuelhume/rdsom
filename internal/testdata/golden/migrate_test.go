package rdsom_test

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestMigrateAll(t *testing.T) {
	idxs := []string{
		fmt.Sprintf("rdsom:rdsomgolden:idx:Bar:%d", _codegenAt),
	}
	sort.Strings(idxs)
	ctx := context.Background()
	err := _client.MigrateAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	got, err := _client.Redis().Do(ctx, "FT._LIST").StringSlice()
	if err != nil {
		t.Error(err)
	}
	if len(got) == 0 {
		t.Error("no indexes were updated by rdsom.MigrateAll")
	}
	got = sort.StringSlice(got)
	if !reflect.DeepEqual(got, idxs) {
		t.Errorf("got %#v want %#v", got, idxs)
	}
}
