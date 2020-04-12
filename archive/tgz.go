package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"

	"github.com/golang/glog"
)

type GzipTar struct {
	buffer     *bytes.Buffer
	tarWriter  *tar.Writer
	gzipWriter *gzip.Writer
}

func NewGzipTar() *GzipTar {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)
	return &GzipTar{
		buffer:     &buf,
		tarWriter:  tw,
		gzipWriter: gw,
	}
}

func (t *GzipTar) AddFile(name string, content []byte, mode int64) (err error) {
	hdr := &tar.Header{
		Name: name,
		Mode: mode,
		Size: int64(len(content)),
	}
	if err = t.tarWriter.WriteHeader(hdr); err != nil {
		glog.Error(err)
		return
	}
	if _, err = t.tarWriter.Write(content); err != nil {
		glog.Error(err)
		return
	}
	return
}

func (t *GzipTar) Close() (content []byte, err error) {
	if err = t.tarWriter.Close(); err != nil {
		glog.Error(err)
		return
	}
	if err = t.gzipWriter.Close(); err != nil {
		glog.Error(err)
		return
	}
	content = t.buffer.Bytes()
	return
}
