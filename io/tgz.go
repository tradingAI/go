package io

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

// ReadTGZContent ...
func ReadTGZContent(r io.Reader, names ...string) (data map[string][]byte, err error) {
	gzf, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	tarReader := tar.NewReader(gzf)

	data = make(map[string][]byte)

	tmp := make(map[string]io.Reader)
	for {
		header, e := tarReader.Next()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		bf, e := ioutil.ReadAll(tarReader)
		if e != nil {
			return nil, e
		}

		tmp[header.Name] = bytes.NewReader(bf)
	}

	for _, name := range names {
		r, has := tmp[name]
		if !has {
			err = fmt.Errorf("missing %s", name)
			return
		}
		content, e := ioutil.ReadAll(r)
		if e != nil {
			err = e
			return
		}
		data[name] = content
	}

	return
}

// ReadAllTGZContent ...
func ReadAllTGZContent(r io.Reader) (data map[string][]byte, err error) {
	gzf, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	tarReader := tar.NewReader(gzf)

	data = make(map[string][]byte)

	for {
		header, e := tarReader.Next()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		bf, e := ioutil.ReadAll(tarReader)
		if e != nil {
			return nil, e
		}

		data[header.Name] = bf
	}

	return
}

// ExtractTgz ...
// refer to https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func ExtractTgz(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer gzr.Close()

	if err = ExtractTar(dst, gzr); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// IsTGZ ...
func IsTGZ(r io.ReadSeeker) bool {
	defer r.Seek(0, io.SeekStart)

	gzr, err := gzip.NewReader(r)
	if err != nil {
		glog.Error(err)
		return false
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	header, err := tr.Next()

	switch {
	case err == io.EOF:
		// empty tgz is fine
		return true

	case err != nil:
		// failed to read
		glog.Error(err)
		return false

	case header == nil:
		// not sure how it happens
		glog.Warningf("empty header")
		return false
	}

	return true
}

// ExtractTar ...
func ExtractTar(dst string, r io.Reader) error {
	tr := tar.NewReader(r)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			glog.Error(err)
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					glog.Error(err)
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			dir := filepath.Dir(target)
			if err := os.MkdirAll(dir, 0755); err != nil {
				glog.Error(err)
				return err
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				glog.Error(err)
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				glog.Error(err)
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
