package diff

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"

	"github.com/google/uuid"
)

// Diff reads from an io.Reader and compares it against the contents of an approved file.
// If they differ, it creates two temp files and opens them in VS Code.
func Diff(got io.Reader, approvedFileName string) (ok bool, err error) {
	r, err := os.Open(approvedFileName)
	if err != nil {
		return false, err
	}
	defer func() {
		err = r.Close()
	}()
	buf := bytes.NewBuffer(nil)
	rcv := file{
		name: fmt.Sprintf("testdata/received-%s.go", uuid.New().String()),
		val:  io.TeeReader(got, buf),
	}
	apv := file{
		name: approvedFileName,
		val:  r,
	}

	res, err := check(rcv, apv)
	if err != nil {
		return false, err
	}
	if res.ok {
		return true, nil
	}

	// Create the temp files.
	os.WriteFile(rcv.name, buf.Bytes(), 0664)
	defer os.Remove(rcv.name)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	cmd := exec.Command(res.cmd[0], res.cmd[1:]...)
	cmd.Start()
	go func(cmd *exec.Cmd) {
		<-c
		os.Remove(rcv.name)
	}(cmd)
	cmd.Wait()

	return false, nil
}

func check(received, approved file) (result, error) {
	var (
		xs  = bufio.NewReader(received.val)
		ys  = bufio.NewReader(approved.val)
		res = result{}
	)

	// Check the contents byte by byte.
	for {
		x, errx := xs.ReadByte()
		y, erry := ys.ReadByte()
		if errx == io.EOF && erry == io.EOF {
			res.ok = true
			return res, nil
		}
		if errx == io.EOF || erry == io.EOF {
			break
		}
		if errx != nil {
			return res, errx
		}
		if erry != nil {
			return res, erry
		}
		if x != y {
			break
		}
	}
	res.cmd = []string{"/usr/bin/code", "-w", "-d", received.name, approved.name}
	return res, nil
}

type result struct {
	ok  bool
	cmd []string
}

type file struct {
	name string
	val  io.Reader
}
