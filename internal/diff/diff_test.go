package diff

import (
	"reflect"
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		desc     string
		approved file
		received file
		ok       bool
		cmd      []string
	}{
		{
			desc: "different content",
			approved: file{
				name: "approved.go",
				val: strings.NewReader(`foo bar baz bat
bat baz bar foo`),
			},
			received: file{
				name: "received.go",
				val:  strings.NewReader(`foo bar baz bat`),
			},
			ok:  false,
			cmd: []string{"/usr/bin/code", "-w", "-d", "received.go", "approved.go"},
		},
		{
			desc: "same content",
			approved: file{
				name: "approved.go",
				val: strings.NewReader(`foo bar baz bat
bat baz bar foo`),
			},
			received: file{
				name: "received.go",
				val: strings.NewReader(`foo bar baz bat
bat baz bar foo`),
			},
			ok:  true,
			cmd: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := check(tt.received, tt.approved)
			if err != nil {
				t.Fatal(err)
			}
			if got.ok != tt.ok {
				t.Errorf("got %v, want %v", got.ok, tt.ok)
			}
			if !reflect.DeepEqual(got.cmd, tt.cmd) {
				t.Errorf("got %v, want %v", got.cmd, tt.cmd)
			}
		})
	}
}
