package archive

import (
	"bytes"
	"compress/zlib"
	"io"
)

// CompressDeflate ...
func CompressDeflate(data []byte) (compressed []byte, err error) {
	deflated := bytes.NewBuffer(nil)
	deflator := zlib.NewWriter(deflated)
	_, err = deflator.Write(data)
	if err != nil {
		return
	}
	if err = deflator.Flush(); err != nil {
		return
	}
	if err = deflator.Close(); err != nil {
		return
	}
	compressed = deflated.Bytes()
	return
}

// UncompressInflate ...
func UncompressInflate(data []byte) (uncompressed []byte, err error) {
	inflator, err := zlib.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer inflator.Close()

	out := bytes.NewBuffer(nil)

	_, err = io.Copy(out, inflator)
	if err != nil {
		return
	}

	uncompressed = out.Bytes()
	return
}
