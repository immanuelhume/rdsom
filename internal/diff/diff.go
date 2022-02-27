package diff

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type diffChecker struct {
	got    *os.File
	want   *os.File
	editor editor
	errors []error
}

type viewer interface {
	// View opens an editor.
	View() error
	// Close closes the editor.
	Close() error
}

func NewDiff(got io.Reader, approvedF string) *diffChecker {
	d := &diffChecker{}
	want, err := os.Open(approvedF)
	if err != nil {
		d.errors = append(d.errors, err)
		return d
	}
	gotX, err := ioutil.TempFile("", "received-*.go")
	if err != nil {
		d.errors = append(d.errors, fmt.Errorf("diff.NewDiff error creating temp file: %w", err))
		return d
	}

	// Copy the received contents to the tmp file.
	if _, err := io.Copy(gotX, got); err != nil {
		d.errors = append(d.errors, err)
		return d
	}

	// Reset the file pointer.
	if _, err := gotX.Seek(0, io.SeekStart); err != nil {
		d.errors = append(d.errors, err)
		return d
	}
	d.got = gotX
	d.want = want
	return d
}

// Do performs the diff check, and closes all open files.
func (d *diffChecker) Do() (viewer, error) {
	if len(d.errors) != 0 {
		return nil, d.errors[0]
	}
	// Conduct the actual byte-by-byte check.
	ok, err := check(d.got, d.want)
	if err != nil {
		return nil, err
	}
	err = d.got.Close()
	if err != nil {
		return nil, err
	}
	err = d.want.Close()
	if err != nil {
		return nil, err
	}
	if !ok {
		switch d.editor {
		case VSCode:
			return newVSCodeViewer(d.got.Name(), d.want.Name()), nil
		}
	}
	return nil, nil
}

func (d *diffChecker) With(e editor) *diffChecker {
	d.editor = e
	return d
}

type editor int

const (
	VSCode editor = iota
)

type vsCodeViewer struct {
	x   string    // first filename
	y   string    // second filename
	cmd *exec.Cmd // command executed to open VSCode
}

func newVSCodeViewer(x, y string) *vsCodeViewer {
	// The -w flag tells VSCode not to detach from terminal.
	args := []string{"-w", "-d", x, y}
	cmd := exec.Command("code", args...)
	return &vsCodeViewer{x, y, cmd}
}

func (v *vsCodeViewer) View() error {
	err := v.cmd.Start()
	if err != nil {
		return err
	}
	defer v.Close() // Unhandled error.
	err = v.cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (v *vsCodeViewer) Close() error {
	return v.cmd.Process.Kill()
}

// check simply checks if the data from two readers
// are the same.
func check(x, y io.Reader) (bool, error) {
	xs := bufio.NewReader(x)
	ys := bufio.NewReader(y)
	// Check the content byte by byte.
	for {
		x, errx := xs.ReadByte()
		y, erry := ys.ReadByte()
		if errx == io.EOF && erry == io.EOF {
			return true, nil
		}
		if errx == io.EOF || erry == io.EOF {
			break
		}
		if errx != nil {
			return false, errx
		}
		if erry != nil {
			return false, erry
		}
		if x != y {
			break
		}
	}
	return false, nil
}
