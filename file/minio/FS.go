package minio

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/pcelvng/task-tools/file/stat"
)

var _ fs.FS = (*fileSystem)(nil)

type fileSystem struct {
	root string
	opt  Option
}

func NewFS(root string, opt Option) fs.FS {
	return &fileSystem{
		root: strings.TrimLeft(root, "/"),
		opt:  opt,
	}
}

func (fs *fileSystem) Open(name string) (fs.File, error) {
	p := fs.root + "/" + name
	r, err := NewReader(p, fs.opt)
	if err != nil {
		return nil, fmt.Errorf("error opening %s:%w", name, err)
	}
	sts, err := Stat(p, fs.opt)
	return &file{reader: r, stats: sts}, err
}

type file struct {
	reader *Reader
	stats  stat.Stats
}

func (f file) Stat() (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (f file) Read(bytes []byte) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (f file) Close() error {
	//TODO implement me
	panic("implement me")
}
