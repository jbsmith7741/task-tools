package nop

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"strings"
	"time"

	"github.com/pcelvng/task-tools/file/stat"
)

// Stat can be used as a mock stat for testing.
// The following are supported mocks
// error, err, stat_error or stat_err - return an error when called
// stat_dir - identifies the path as a directory
func Stat(pth string) (stat.Stats, error) {
	u, _ := url.ParseRequestURI(pth)

	switch strings.ToLower(u.Host) {
	case "error", "err", "stat_error", "stat_err":
		return stat.Stats{}, errors.New("nop stat error")
	}
	return stat.Stats{
		Path:     pth,
		LineCnt:  10,
		Size:     123,
		Checksum: fmt.Sprintf("%x", md5.Sum([]byte(pth))),
		IsDir:    strings.Contains(strings.ToLower(pth), "stat_dir"),
		Created:  time.Now().UTC().Truncate(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

// NewFS can be used to mock a fs.FS for testing.
// the following are supported mocks
// error, err, init_err - return an error on init
func NewFS(root string) (fs.FS, error) {
	switch strings.ToLower(root) {
	case "error", "err", "init_err":
		return nil, errors.New("nop init fs error")
	}

	return &filesystem{}, nil
}

type filesystem struct{}

// Open can be used for mocking
// the following are supported mocks
// error or err will return an error value
// embed fileInfo data in path as a url ex: path?
func (fs *filesystem) Open(path string) (fs.File, error) {
	switch strings.ToLower(path) {
	case "error", "err":
		return nil, errors.New("nop fs.Open error")
	}
	return &file{}, nil
}

type file struct{}

func (f *file) Stat() (fs.FileInfo, error) {
	return nil, nil
}

func (f *file) Read([]byte) (int, error) {
	return 0, nil
}
func (f *file) Close() error {
	return nil
}
