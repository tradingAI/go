package archive

import (
	"archive/tar"
	"bytes"

	"github.com/golang/glog"
)

type Tar struct {
	buffer *bytes.Buffer
	writer *tar.Writer
}

func NewTar() *Tar {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	return &Tar{
		buffer: &buf,
		writer: tw,
	}
}

func (t *Tar) AddFile(name string, content []byte, mode int64) (err error) {
	hdr := &tar.Header{
		Name: name,
		Mode: mode,
		Size: int64(len(content)),
	}
	if err = t.writer.WriteHeader(hdr); err != nil {
		glog.Error(err)
		return
	}
	if _, err = t.writer.Write(content); err != nil {
		glog.Error(err)
		return
	}
	return
}

func (t *Tar) Close() (content []byte, err error) {
	if err = t.writer.Close(); err != nil {
		glog.Error(err)
		return
	}
	content = t.buffer.Bytes()
	return
}
