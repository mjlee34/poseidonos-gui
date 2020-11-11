// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../resources/events.yaml
package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5d\xdd\x52\x1c\x39\xb2\xbe\x9f\xa7\x50\x70\xb1\x3b\x13\x01\x3e\xa6\x7f\x6c\xe0\x0e\xf3\xe3\xc3\x86\x31\x5e\xda\xf6\x9c\x73\x35\xa1\xae\x52\x77\x6b\xa8\x92\x6a\x24\x15\xd0\xb3\xb1\xef\xc2\xb3\xf0\x64\x27\xf4\x53\x3f\x5d\x52\x55\x67\xab\xbd\x73\x26\x62\x22\x0c\xa8\xf4\x65\xa6\xa4\x54\x2a\x95\x99\x3a\x3a\x3a\xfa\xe9\x23\xc7\xd9\x19\xba\x7a\x24\x4c\x49\xf4\x8d\xd1\x05\x4d\xb0\xa2\x9c\xfd\xf4\x9d\x08\x49\x39\x3b\x43\x07\x8f\x6f\x4e\x47\x07\x3f\xe5\x3c\x2d\x33\x22\xcf\x7e\x42\xe8\x08\x31\x9c\x93\x33\x74\x70\x71\x77\x7b\x7b\xf7\xf9\xe0\x27\x84\x10\x4a\x78\xc9\xd4\x19\x3a\x3d\x3d\x35\x3f\xd2\x74\xa6\xb0\x50\x67\xe8\xad\xfb\xf1\x8a\xa5\x67\x08\x35\x7f\x67\x0b\x7e\x66\xfe\xa5\xfb\x4b\x78\x4a\xaa\xa6\xfa\xbf\x8c\x3c\x92\xec\x0c\x1d\xdc\x7c\xbe\xbe\x3b\xa8\x7f\x9b\x13\x29\xf1\x52\x03\xcf\xca\x24\x21\x52\x36\x7f\x2a\x04\x9f\x67\x24\x3f\x43\x07\xcd\xef\x24\xcf\x4a\x65\x59\x38\xd8\xa0\xfa\xd3\xcd\x06\xc9\xc7\x6f\xdf\x6e\x92\x7c\xfc\xf6\x6d\x87\xea\xe3\xb7\x03\x64\xd7\xcd\x5b\x94\x6b\xc2\x7d\xba\xe9\x9c\x2f\xb8\x44\x54\x22\xa9\x91\x48\xea\xd1\xef\x13\xef\x83\x1d\xef\x0a\xa6\x88\xc8\x29\xc3\xb1\x78\x23\x18\x5e\x92\x51\x24\x89\x78\x24\x42\x63\x52\x46\x15\xc5\x19\xfd\x33\x12\x74\x0c\x03\x65\xe4\x49\x03\x13\xa6\x34\x68\xc2\x19\x23\x49\x2c\x9f\x13\x30\x9f\x0e\x2e\xa5\x72\x3f\xc4\x29\x0c\x51\x90\x3f\x4a\x22\x15\xca\xe5\x52\xc3\x0a\x92\x10\xfa\x18\x09\xf9\x0e\x0a\x29\x0b\xce\x24\xa9\x30\x25\x61\x2a\x06\xef\x18\xb8\x32\x5a\x93\xa7\x10\xa4\xc0\x82\xb2\x25\x22\xcf\x34\x0e\x14\xb8\x42\x5a\xa0\x6a\x25\x08\x4e\xd1\xef\x9c\xb2\x38\xc1\x1e\x03\x57\x09\x61\x78\x9e\x11\x24\x48\x29\xc9\x11\x4e\x53\x11\x05\xe6\xad\x8e\x5f\xcf\xef\x3f\x03\xc0\xd0\x02\xd3\x2c\x92\x41\x6f\x79\x5c\xdd\xdf\xdf\xdd\xfb\xa0\x92\x27\x0f\x44\xa1\x44\x10\xb3\x91\xec\x03\xe9\xad\x8f\x61\xc8\x39\x65\xa9\x9e\x37\x7b\x20\x7a\xcb\x63\x18\x31\xa3\x52\x91\xbd\x58\x7c\x0f\x04\x24\x05\xcf\xb2\x1f\x22\xd4\x93\xdd\x58\xc4\x49\x42\x0a\xb5\x0f\xe0\x29\x10\x30\xc7\xcf\x95\x72\x25\x42\xf0\xa8\x95\x31\xf2\xf4\x4d\x1f\x98\xfd\x87\x56\x6c\x7b\xce\x99\x91\xa7\x6d\xb6\x40\x5a\xfd\xbd\x27\xa8\xa7\x6e\xc2\x1a\xe0\xeb\xcd\xed\xd5\x25\xba\xfb\xf6\x35\x0a\xc4\x53\x33\x3d\x9c\xdd\x7c\xfe\x7e\xfe\xe9\xe6\x12\x7d\x39\xbf\x3f\xbf\x8d\x41\x1a\x03\xb7\x89\x2f\x77\x33\x74\x33\x43\x1f\xbe\xcd\xfe\x17\x06\x53\x1b\x7d\x37\x1f\xf8\xf5\xdd\x0c\xcd\x14\x56\x04\xdd\x62\x86\x97\x44\x6c\x58\x81\x23\xcf\x0a\x3c\xf6\xac\xc0\xd1\x90\x15\x78\xec\x1b\x66\x97\x57\x1f\xbe\x7d\x0c\xac\x2c\x43\x44\xc2\x99\x22\xcf\x0a\xe1\x34\x8d\x9a\x03\xc7\x50\xc3\xcc\xc1\xad\x30\x5b\x46\x01\x8d\xc6\x9e\xea\xbf\xbf\xfa\x72\x77\xff\xf5\xb7\xaf\xf7\xe7\x17\x57\x01\xcb\xd3\xc9\x7a\x2d\x15\xc9\xd1\x3d\x49\xf8\x23\x11\x6b\x74\xc3\x0a\xc1\x97\x82\x48\xb9\xe3\xd8\x7d\xe7\x59\x99\x93\xd0\xa0\x8d\xbb\x83\x36\xf2\x4c\xf7\xd1\xd0\xa0\x8d\xc0\xa6\xbb\xa5\xc1\x2a\xe0\x18\x29\x8e\xc0\x76\xbb\x43\x4a\x49\x46\x62\x91\x80\x13\xc3\x21\xe5\x5a\x96\x91\x48\x40\x33\xdd\x21\x95\x6c\x1f\x2c\xa0\x7d\xee\xb0\xcc\xb2\x42\x8a\x23\xb5\x22\x66\x9f\x8e\xc2\x04\x5a\xe8\x16\xf3\xf5\x45\x90\x9c\x3f\x92\x14\x2d\x04\xcf\x35\xf0\xeb\x4b\x2c\xb2\x6f\x38\xf7\x28\xf8\x15\xa9\xce\x07\x24\x45\x8f\x6e\xee\x70\x22\x11\xe3\x4a\x5b\xd0\x01\x78\xf3\x91\x6b\xfa\x44\xd5\xca\x88\xc8\xeb\x44\x2f\x3e\xc4\x45\xf5\xe3\xcd\x65\x5f\xb7\xcd\x39\xfb\x8a\x29\x63\x4b\x6b\xe5\x26\x04\x49\x54\x7f\x5f\x78\xa1\x9b\x26\x2b\x92\x3c\xe8\x9d\x50\x35\x14\xb5\x24\xd6\x12\x87\xb7\x78\x7a\xc4\xc1\x39\xca\x31\x5b\xbb\xce\x7c\x4d\x53\x2f\x64\xcc\x0c\x2b\xf3\x7a\x4d\xa3\x39\x49\x70\x29\x89\xa1\x25\xc7\xcf\x34\x2f\x73\xc4\xca\x7c\x4e\x04\xe2\x8b\xaa\x43\xa4\x56\x58\x99\xaf\x5b\x5f\x52\x89\xc8\x73\x42\x48\x5b\x91\x37\x52\xb9\x27\x4a\xac\x1d\xc3\x66\x82\x68\x86\x4b\x7d\x6e\xd4\x54\x8b\x8a\x56\x84\x73\x6e\x0f\x3d\x52\xe9\x16\x55\xe7\x9b\x9c\xb4\x44\x02\x34\x01\xae\x0c\x65\xb6\x3b\x73\x1c\x70\x70\x92\xfe\x49\xc2\x53\xc3\x9b\x0a\xba\xa9\xe3\x50\xa2\x73\x21\xf0\x1a\x25\xb8\xc0\x09\x55\xeb\x00\xbf\x17\x7a\x50\x8d\x14\xa5\xdd\x01\xaa\xb6\x08\xb3\x14\x19\x59\x2c\x31\x65\x1e\x43\xbe\xed\x16\x66\xe8\x7b\x6b\x4e\x51\x89\x14\xe7\x48\xae\xb8\x18\x9e\xe7\xa6\x35\xd1\xf3\xd3\x8e\x97\xea\x7e\xd4\x9d\xc5\x78\xe3\xcb\x39\x51\x4f\x84\x30\x34\x32\x3c\x8c\xa6\x53\xbd\x9f\x0a\x9c\x28\x22\xfc\x91\xf1\x2d\xc2\xad\x8c\xbc\xbe\x54\xac\x64\x9c\x2d\x7b\x67\xad\xcf\x45\xe7\x83\x61\x2e\xcc\xdc\x6d\xad\x62\x33\x31\xb6\x31\x03\x9c\x66\x9d\x51\x49\xcb\x22\xa3\x49\x70\xaf\x44\xe7\x1b\xca\x07\x37\x6d\xed\xd7\x38\xd3\x67\xf1\xb5\x5d\x09\x72\x80\xb5\x94\x2e\x16\x44\xe8\xf3\x42\x8b\x49\x9f\x01\xe0\x61\xf9\x1b\xb3\x47\x9d\xf6\x1a\x69\x75\xb8\x65\x30\xb4\x35\x87\x29\x93\x08\x23\xa9\x84\x55\x69\xd8\x78\x89\xb4\xa8\x71\x96\xf1\xa7\xa0\x72\x68\x54\xa6\x37\x50\x39\x21\x4a\x7a\x7f\x12\x65\x16\x50\x06\xbe\x01\x1d\x66\xf2\xbe\x5e\xda\x66\x4d\x9b\xf9\x83\xc5\x72\x27\x45\x50\xcd\xbb\x8d\xef\xc2\xea\x6e\x43\xbf\x17\x5c\x4a\x1a\x56\x40\x2d\x46\x80\x6b\x27\xc0\x88\xcc\x71\x96\xed\xce\xc8\xeb\xcb\xe6\x87\x21\x45\x96\x53\x66\xb6\x03\x3d\x8e\x89\xaf\x46\x8d\x52\x10\x9a\x6f\x9f\x1f\xe0\xf2\xe9\xf0\x63\x27\x0d\x5d\x86\xbc\x51\x41\x8e\x5e\x5f\xcc\x77\xf5\xe2\x36\x1f\xdb\x05\x36\xcf\x78\xf2\xb0\xa9\xee\x1b\x1e\x6f\x58\x51\xaa\x0d\x5e\x14\xd7\x5b\x5b\x5e\x66\x8a\x16\x19\xd1\x5b\x9f\xd7\x41\xc3\xde\x64\x37\x9d\x5d\x2d\xed\x3e\x33\x10\x9d\x2b\x45\xf2\x42\x69\x22\x4c\x9b\x46\x81\x55\xcb\xa9\xaf\x8b\x86\xa5\xcf\x5c\xad\xcc\x9c\xe3\x28\xe5\x3e\xc5\xbb\x29\xe7\x0a\xae\xdf\x74\x6d\xd3\xec\x5a\xf5\x53\xed\x77\x03\xa5\x7b\xea\x1f\x54\xd8\x82\xfb\x74\x5f\x38\xc3\xc6\x9e\x1f\x2a\x31\x39\x7a\x7a\x89\xb7\xa6\xb2\xfb\x06\xf7\x7d\xd5\xd0\xda\xe2\xd9\x7d\x54\x19\x31\x66\xdd\x3b\x36\x37\x2d\x3b\x8f\xa5\x77\xbe\xb3\x66\xc3\xef\x53\xf3\x74\x8d\x69\xa6\xa1\xac\x49\x54\x41\xe5\x44\xe1\x18\xd3\xfa\x9d\xef\xb0\x19\x86\xe5\x05\x61\x7b\x83\x7a\x7a\x60\x0b\xa8\xf1\x47\xef\x0b\xea\x3b\x70\x86\x41\x9f\x04\xfd\x01\xf2\xf5\x3d\xc5\x61\xd4\xef\x0d\xce\xeb\x8b\x71\x87\x30\x85\xe6\x82\x3f\x10\x16\x83\xfb\x1e\x3a\x9d\x6e\x49\xce\xf5\x16\x65\x95\x39\xe5\xec\xf5\x65\x81\x69\x56\x8a\xd0\xfa\x40\x09\x96\x6e\x1d\xcb\x15\x2f\xb3\x14\x31\xf2\xa8\x8f\x04\x49\x52\x0a\x74\x84\x56\x04\x17\xad\xae\x50\xb7\xa7\x66\xcd\x7c\xed\xb5\x7c\xdf\x43\x67\xe4\x0d\x7b\xc4\x19\x4d\x11\x65\x29\x79\xee\xf1\x92\x6e\x27\xd9\x7c\xfd\xb3\x1b\x65\x9a\xfe\x82\xa8\x36\x42\x18\xce\xb2\x35\x5a\x0a\xcc\xdc\x91\x86\x5a\xb0\xe0\xa6\x61\xdb\xa3\x8c\x2f\x69\xf2\xfa\xd2\x26\xa4\xc5\xd5\xae\x53\xde\x48\xf1\xcd\xeb\x0b\x23\x4f\xaf\x2f\xf5\x51\x31\x82\xc1\x6f\xf6\xde\x43\x71\xb4\xa4\x8f\xa4\x39\x75\x1e\xa2\x94\xc8\x42\x4f\xf1\x96\x55\x65\x5c\x49\x95\xa1\x96\xe3\xe7\x78\x7e\xa1\xab\x4d\xef\xdf\x12\xdb\x73\xb0\x23\xa2\x63\xeb\xc2\x59\x3d\xf7\x0e\xf2\x55\xcf\x5b\xcd\x68\x18\x57\x27\xfe\x0e\x8f\x05\xa3\xad\x33\x47\xcd\xd7\x5d\xa9\xb4\xb9\xf0\x07\x97\x48\x60\x16\x32\x2a\xcf\xd1\x23\xce\x4a\x82\x32\x22\xcd\x49\x9a\x6d\x5a\x57\x85\x39\x07\xe8\xa1\xd3\x7d\xd8\xa6\x4f\x58\x56\x46\x36\xc8\x44\x6b\xbe\x6c\x9f\xd4\x2b\x33\x7d\xe3\x04\xfa\xc6\x63\xf6\x14\xce\xac\x3d\x06\x9b\x39\xd3\x67\x18\x34\xae\x86\x8e\xa7\x81\x0b\x94\x71\x9c\x92\xd4\x8c\x1a\x2f\x55\x75\x59\xdf\x6f\x1c\xd4\xca\xc3\xed\xb0\xd6\xce\xb0\x9f\xf9\x6c\xf8\x36\x4e\x1f\x1b\xce\x49\x7b\x8d\xcb\x2c\x70\x80\xae\x38\xe0\x79\xae\x25\xd7\x70\x52\x10\xb1\xe0\x22\xd7\x8a\xc2\x8e\xe1\xec\xeb\xdd\x17\xeb\x69\x06\x68\xea\x53\xdf\x28\xee\xa3\xef\x92\x33\x37\xb7\x7b\xb4\xdd\x8c\xeb\xe5\xa3\xff\x26\x51\x8e\xd7\x6e\x61\xa4\xa5\xa8\x8f\x1d\x82\x27\x44\x4a\xfd\x23\x5f\xb4\x5d\x5d\x87\x76\x36\xe8\x25\x53\xce\xa5\xfe\x1d\x53\x7a\xaf\x17\x56\x91\xe7\x6e\x78\x9f\xb8\x78\x40\x4f\x24\xcb\xde\x84\x8e\x6f\x1a\x18\x2d\xb8\xb0\x24\xa0\x15\x66\x69\xa6\xa1\x70\xa6\x07\x76\xb9\x42\x54\x55\x62\xb3\x94\x19\x5e\x4a\x49\x04\xb2\x90\x89\x67\x17\x1d\xfb\x3e\xe9\x5e\xf1\x68\xca\xed\x08\x6a\x84\x3e\xef\xb4\x6f\x56\xb7\x95\x06\xe3\xad\x6e\xba\x5d\x34\xac\xce\xaa\x36\x12\xe5\xa5\xec\x78\xce\x16\x5c\x38\x5b\x53\x33\xdf\xe3\xb2\x0a\xdc\x90\xf4\x32\xd6\x31\xbf\x57\x58\x22\x49\x54\x4d\x28\x43\xec\x0f\xf6\x1f\x22\x12\x3e\x39\xeb\x8d\x4b\x29\x9c\xac\xcc\xf9\x5c\x16\x38\x31\x3b\x4f\x2d\xd2\x5e\xad\x45\x17\xe6\xca\xb3\xfe\x4a\x1a\x5d\x27\x0b\x92\xd0\x05\x25\x69\x35\x87\x3b\x63\xa3\xa7\xe6\xcf\xe4\xf9\x0d\x3a\xca\xd1\x68\xfa\xee\x97\xf6\x45\xc9\xed\x87\xfb\xe0\xdd\x96\x17\xe1\x34\x1a\x7b\xd7\x24\xe3\xa1\x6b\x92\x31\xf4\x9a\xe4\xdc\x1a\x42\xda\xa8\x33\x36\x96\xb4\x81\x5a\x11\xe6\xdc\x18\x7a\x5f\x92\xcf\x05\x4a\xb1\xc2\xd5\xe9\x40\x2f\x6a\x63\xc7\x46\x81\x02\xaf\x4e\x6a\x50\x9c\xa6\x7b\x22\x02\xaf\x50\x0a\x2c\xa8\x5a\x5b\x8f\xca\x5e\x62\x05\x5e\xa3\xb8\x39\x57\x96\x34\xdd\x1f\x14\x7a\x9b\x91\x92\x47\x9a\x58\xef\xc7\x82\x97\x2c\xe6\x9e\x68\x0c\xbd\x2a\xd8\x10\xa8\xb6\xda\xa3\xc0\x80\xee\x1d\x5f\x9a\xd1\x88\x40\x77\x66\x3d\x43\xf7\x13\xa6\x37\x5b\xfa\xd1\x1e\xc8\x7a\x3f\x30\xdf\x89\xd9\x13\x6f\xd0\xd5\x32\xb1\xb2\xf4\x9d\x8d\x7d\xa1\x1b\x95\x30\xed\x4d\x48\x5c\x74\xca\x68\xec\x3b\x03\x07\xf0\xb4\x7d\x85\xa3\x43\x61\x46\xe3\x31\x34\x7a\x63\xa9\xf7\xd5\xd6\xfc\x84\x4b\xb3\xde\x78\x2e\xcd\xba\x85\x6d\x3d\x13\x6f\xeb\x99\x0c\x6d\x3d\x13\xe8\xd6\x73\x63\x23\x4e\x91\x4c\x70\x8c\xfb\x60\x02\xdd\x6f\xee\x49\x34\x02\x70\x73\xf9\x44\xa5\x72\xaa\x30\x0a\x06\xb8\xa3\x18\x18\xac\x8f\x31\xaf\x2f\x7b\xa0\x05\xf7\x93\xff\x0a\xab\x89\x6b\xca\x52\xcd\xd9\xcf\x25\x4d\x7f\x89\x42\x03\x87\x04\xea\x83\x5f\xec\xda\x99\xf8\x61\xb2\x3d\x30\x6e\xc3\x32\x1e\xc2\x78\x34\x68\x10\xa0\x43\x4b\x32\x2e\xe3\xb5\xd0\xe4\xad\x17\x01\x38\xb8\x19\xbf\xbe\xf4\x2b\x75\x64\x17\x7e\xeb\x36\xc2\xde\x1e\x54\x71\x10\xda\xa0\xf5\xbf\x1e\x20\xcd\x8b\x15\x0c\x93\xf6\x99\x33\x03\x61\x7c\x05\x97\xe1\xb9\x8b\xee\x58\xb6\x46\x29\xd1\xa6\x39\x49\x6b\xd9\x59\x7f\x80\x23\x10\x40\x12\x34\x82\x39\xc5\x24\xe7\xcc\x06\xf6\xc7\x8c\x0a\x34\x68\xb9\xc6\xe1\x45\x14\x0c\x50\x07\x39\x71\xd9\x93\x4d\xae\x4f\xc6\x29\x51\x91\x91\xee\xa3\x89\x6f\xb4\x0c\xa1\x6a\x75\xf4\x63\x60\x81\xb6\xae\x3b\xbf\xc5\x2b\x41\x3f\x52\xb9\x8f\xbf\x7d\x81\x3c\xc5\x14\x5e\x20\x1b\x1c\xc5\xab\x0a\x3f\x3c\xb9\x4f\x55\xfc\x18\x38\xa0\x66\x2a\x6b\xff\xae\x59\x6e\x28\xe7\x8c\x2a\x2e\x68\x28\x72\x02\x67\x59\xeb\xef\x6e\xf9\x48\x84\x45\xed\x5c\x78\x7d\x11\x25\xd3\x07\xfc\x43\xc4\x05\x62\xbc\x6a\x2e\xfb\xe2\x9c\x7c\xba\x81\x6a\xab\x4d\x37\x2f\xb6\x91\xfd\xfa\xe2\xd3\xfd\xfa\xd2\x22\xdc\xf4\x52\x90\x34\x9a\x6e\x70\xb4\x74\x4b\xe9\x14\x24\x45\x25\x23\xcf\x85\x59\x95\xd9\xda\x27\x1d\xd2\x78\x80\x26\x68\x14\xad\xcd\x13\x21\x7b\x2c\x26\x3f\xb4\xa5\x07\x4a\x90\x8c\x60\xb9\x1b\x54\x6d\x1b\x1b\x47\xf1\x96\xa8\xd5\xa9\x67\x13\xbf\x1f\xb2\x89\xa7\xbe\x4d\xdc\x77\x52\x32\x5e\x6a\x77\x01\x9d\xe2\xfa\xda\xfc\x4d\x84\xbc\xa6\xbe\x85\x0c\x41\x35\xd7\xde\x7b\xc1\x42\x4f\x4d\x16\xb6\xe7\xde\x63\x17\x40\xe8\xc9\xc9\x02\x6e\x86\x45\xc6\x21\x42\x93\x6b\x2e\xf7\x77\x90\x4c\xc1\x36\xb4\x5a\x91\x56\x00\xa4\x94\x69\xe3\x54\xad\xe2\xe1\xc6\xc6\xf9\xd5\x3a\xd7\xee\x42\x07\xd4\xc8\x66\x8f\x79\xb5\xb5\xec\x75\xfe\x9e\x82\x0d\xed\x99\x74\x71\x38\x15\xbf\x4b\xe3\x5b\x14\xf6\xce\xea\x5f\xff\x46\xf3\xb5\x0a\x04\x97\x42\x48\x80\xa6\xdf\xcc\xda\xe2\xae\xef\xf5\x34\x51\x71\xb8\xd0\x2c\x9c\xdb\x0f\xf7\xaf\x2f\x26\xf8\x20\x5a\xcc\xbe\xc9\xdc\x8f\xe5\x42\x0e\xe2\xb1\xa0\x3a\xc9\xdc\x70\x1d\x29\x9e\x11\x81\x59\x42\x8c\x62\x45\x7b\xf2\x09\x55\x4c\xbf\x0a\xce\x96\x1e\x05\x39\x51\x2b\x9e\x22\xb5\x2e\x62\x76\xaf\xa9\x6f\x53\xf7\xa0\x1f\xfc\xeb\xdf\xe8\x0b\x16\x8a\x9a\xbb\x83\xfa\x12\xc1\xb0\xed\x67\x50\x43\x90\xa1\xda\xaa\x41\xe6\xcc\xdc\x78\xee\x03\x0a\xd5\x5a\x46\xd8\xaf\x2f\x56\x37\x93\x47\x93\xab\x1b\xa5\x27\xc1\xc9\x80\x2e\x50\xc2\xde\xa3\xe3\x0c\xe1\x34\x15\x44\xca\x3d\x26\x16\x54\x49\xd8\x89\xa5\x67\x90\xb9\x83\xc4\x76\x3f\xea\x39\x0f\x7f\x5d\x11\xdb\xf4\x67\xdd\x76\x5e\x2e\x16\x5a\xb1\xdb\xdb\xcb\x14\x2b\x7c\x24\x15\x17\x78\x49\x7e\x69\xdd\x41\xcd\xd7\x46\xf7\xb4\x3b\xae\x6f\x4d\x71\xa2\x4a\x9c\x55\xbf\x35\x3d\x1b\x7b\xac\x8a\xca\x0d\xdd\x97\x36\xf7\xf5\xb6\x7d\x5f\xdc\xe4\xd4\x37\x4d\xc3\x27\x2a\x73\xe1\x5a\x59\x66\xa8\x39\x37\xc6\x88\x1d\x1a\xb7\xcd\xb4\xe2\xc8\x31\x65\xe6\x7a\x6e\x3f\x23\x74\x0a\x8d\xaf\x66\xdc\x8d\x82\xb1\x32\x8c\x71\x99\xa2\x03\x85\xc5\x92\x28\xdc\x98\x98\xbb\xa2\x03\x6f\x13\x18\x21\xa9\xbb\x3a\x5e\x50\x21\x4d\xd6\x9a\x9e\x3b\xd6\x8f\xb2\x07\xfb\xd0\xcc\x3f\x6f\x16\xae\xb0\x44\x73\x6d\x82\x44\xe6\xcf\x8d\xa6\xbe\xf7\x7f\x57\xe8\xae\x17\x69\x17\x70\xe0\xb0\xdb\x14\x51\x73\xc9\x9c\xa6\x1e\x29\xfe\x22\xf7\x88\x8d\x8b\xb1\xf7\xe9\x05\x4e\x94\x86\x5e\x37\x35\xb6\x93\xac\xf5\x52\x9b\xd8\xcd\x6c\xa4\x2a\xa1\xc0\x04\x35\x87\x57\xdb\x00\xd5\x50\x07\x50\x8b\x40\xeb\x1b\x48\x32\x82\x45\xe4\xd0\x42\x77\xaa\xcd\xb1\xdd\xba\x9a\x1c\x95\xa1\x40\x8f\x01\x62\xa0\x9b\xd8\xae\xc4\x94\xec\x81\xf1\x27\xb6\x31\x72\x4f\x66\x3b\xda\x48\x9a\x18\xa0\x0c\x6a\x8c\x6f\x52\x36\x34\xf1\x9b\x59\xe3\xa8\x6a\x1d\x41\x5d\x9c\x92\x91\x21\x80\xba\x40\xb8\xf5\xc0\xcc\x91\x0a\xab\x52\xa2\xb2\x48\x23\xb3\x43\xa7\x53\xa0\x32\xb2\xee\x04\x97\xc7\x7c\x86\xae\xcc\x66\xf0\x99\x8b\x1c\x67\x07\x7e\xa7\x40\x67\x6f\xb0\xd3\x4b\xb2\x14\x38\x25\x69\xa0\x5b\xa0\x37\x37\xd8\xed\x2d\x35\xb1\x58\x81\x5e\x81\x6b\x35\xd8\xeb\x07\x13\x35\x1c\xe8\x14\xe8\x98\xed\x74\xda\x2b\x50\x60\xf9\x94\x4e\x77\x03\xa2\xf4\x16\x01\xa8\xc3\x7b\x32\x2f\x69\x96\x86\xe5\xe8\x19\x8f\xa0\x2e\x67\x8a\x17\x1b\x65\x92\x1c\xc6\x96\xdb\xdc\x13\xcf\x73\x75\x32\xe4\xb9\x3a\x81\xde\xe6\x36\x1c\x5e\x92\x79\xb9\xcc\xf8\xd2\xef\x0a\x38\xbb\x9b\xae\x2a\x9f\xa4\xdf\x15\x70\x46\xb7\xba\xda\xa8\xdd\xd0\xea\x69\x30\x57\x1e\xf9\x5d\x0a\xdb\x25\xaa\x72\xe3\x91\x20\x45\x93\x97\xd8\xea\x18\x5c\x2c\x68\xb3\xbf\x08\x65\x74\x02\xad\x12\xe4\xce\x09\x89\x49\x79\x2b\x8c\xba\xcd\x78\xf2\x10\xa5\x01\x4f\x7c\xe7\x0c\x08\xf4\xf5\x85\x4a\x54\xb2\x78\x5c\xe8\x05\x5f\x25\xd8\x84\xe7\x85\xc9\x78\x71\xf1\x50\x8b\x32\x0b\x78\xc1\x21\xc0\x50\xd7\xc5\x41\x05\x2d\x88\x2c\x33\xd5\x90\xe2\x72\x0d\x62\x2c\xff\x13\x68\xe4\x52\x2f\xb8\x5b\x4b\x71\xe0\xd0\x63\x07\xaf\xe1\xec\x49\x27\xbc\xef\x6f\x9c\xfe\x6c\xf2\x82\x8d\x53\xce\xb8\xcd\xa7\xc2\x6c\x8d\x8a\xca\x05\x02\xa1\x0f\xb8\x17\x31\x6e\x4f\x46\xaa\xa6\xd3\x27\xd0\x96\xdf\x70\x24\xa5\x6e\x27\x80\xd0\x00\x35\x23\xaf\x6b\xfb\xc8\xde\x86\x39\x4a\x0e\xab\xa0\x1e\x53\x6e\x8e\x17\x9d\xe8\x6c\xd8\xe5\xc5\x05\x67\x0b\xba\x84\x05\xf6\x9c\x7a\x5b\xc1\x50\xb1\xbf\xba\xf9\x56\x21\x27\x86\x06\xb4\xa0\xa6\x94\x15\x4e\x51\xca\x59\xcc\x59\xf3\x14\x1a\xe1\xb3\x24\xca\xa5\x10\x44\x23\x01\xb7\x26\xeb\x37\xb2\x58\xf1\xe1\x95\xa7\xd0\xed\xab\x0d\x17\x19\x33\x77\x1a\x08\xe5\xec\x0b\x8f\xae\xea\xd6\xb9\x01\x34\x1e\x9f\x58\x9f\xd8\x69\x20\xac\x73\x1b\xee\x03\x59\xc7\xe3\x05\xee\x46\xc3\xd9\x3c\xed\xe9\x69\xe2\x80\x62\x25\x1b\x28\x24\xb5\x15\xd1\xa5\xfd\xc6\x73\x09\xcd\xd1\xda\xc0\xdc\xcb\x6d\x7e\x1a\x28\x2b\xb5\x5d\xb2\x7b\x46\x43\x9e\x8e\xa0\x49\x89\xbf\x4b\xce\x50\xca\x93\x4a\x63\xf3\xf9\xef\x24\x09\xec\x3b\xd5\x24\x6b\x94\xc5\xdf\xec\xfa\x32\x19\xaa\xe6\x17\x36\x53\x67\x73\x62\x1c\xb6\xee\x1c\xd0\xdf\xc2\x6c\x05\x35\xf1\x2d\x2e\x8a\x6d\x25\xab\xc6\x5e\xf5\xa3\xf1\x71\x47\x05\x37\xd7\xd2\x36\x86\x96\x6f\xeb\x72\xe4\x75\xd9\xcd\x14\xa8\xbb\xfc\x78\x31\xbc\x43\x8c\xbd\xd0\xcf\xf1\x50\xe8\xe7\x18\x1c\xfa\xf9\xf1\x02\x29\x41\x97\x4b\x12\xe5\xb1\x19\x83\x43\x3f\x3f\x5e\xc4\x17\x6e\x1d\x83\xc3\x3f\x3f\x5e\xbc\xbe\xc4\x6d\x3c\x63\x70\xec\xe7\xc7\x8b\xba\x02\x4e\x64\xc4\xda\x18\x1c\x19\xf7\x9d\x26\x8a\xe6\xb5\xb9\x9e\x70\x26\x95\x28\x13\x15\xb3\x90\xc7\xe0\x38\xb9\x4f\x1c\x6b\x9b\xf5\x91\x08\x49\x50\x8e\x23\x82\xe5\xc6\xe0\x60\x39\x83\x65\x37\x57\x53\x2b\x61\xd7\x02\x67\xb7\x44\xe1\xeb\xd9\xc6\xe2\x99\x76\x17\xcf\xc4\x5b\xdb\x93\xa1\xc5\x33\x09\x54\x36\xeb\x89\x6e\xb9\xb5\xff\x68\xae\x91\x52\x7d\xe2\x5e\x86\xc2\x92\xb6\xca\x6c\x12\x48\x5e\x0b\xcb\xcc\x43\xd5\x5c\x68\x65\xd8\xb6\xd0\xe1\xb8\xa3\x40\xd2\x5c\xf0\x48\x71\x53\x17\x26\x46\xb7\xd7\xb3\x8d\x84\xec\xdd\xf0\x80\x97\x45\xb7\xe6\xd2\x64\x3f\x28\xe0\x51\xed\x9b\xab\x3e\xb1\x1f\x18\xf0\x68\xd6\x9c\x3a\x08\xfb\xa3\x24\x25\x31\xb5\x98\x73\x19\x88\x65\xeb\x6d\xaa\x7f\x75\x7b\x3d\x3b\xb4\xb9\x96\xee\x20\x45\x9e\x0b\xcc\x52\xb4\x10\xc4\x56\x1f\xfe\x27\x84\x68\x60\x9a\xca\x05\xcf\x0b\x9c\xa8\x81\x0a\xaa\xed\x26\x09\x2f\xb3\x94\xfd\xdd\x44\x60\x68\x85\x8c\xd2\xd2\x04\xef\x99\x58\x1f\x66\xb2\x3b\x0d\x95\x26\x7b\xef\x3f\x43\x65\x0c\x09\x35\x60\x20\x6b\xaf\x2f\x10\xe3\x7a\xa6\xed\x9d\xe8\xea\x7c\x93\x40\xb6\xde\x30\xd4\x1e\xa5\x00\x27\x81\x2c\xbd\x3e\xb0\xae\x92\x89\xb4\x1f\x27\x81\x34\xbd\xbe\x5a\xa8\x2e\x18\xa0\xc0\x02\xe7\xa4\x55\x52\x6b\x37\x38\x68\x88\x85\x9e\x08\xe6\x82\x23\x0a\x05\xea\x62\xa8\x2e\x51\xe2\x91\x76\xbf\x86\x6a\x2a\xc8\x47\x01\xee\x7e\xbb\x34\x37\x81\xb6\x65\x84\xa5\x30\x19\x83\xc3\xba\x18\x47\x39\x49\x29\x36\xb3\xf1\xf6\x7a\x16\x05\x06\x8d\xe5\x72\x45\xa5\x36\xca\x76\xcd\x69\xcc\xb6\x1e\xc8\xe0\xdc\x2a\x4f\x7d\xe6\x89\x82\x82\xaa\x12\x5a\x57\x43\xd9\xc7\x72\x08\xe4\x70\x6e\x65\xcd\xa5\x1a\xeb\xa3\x5d\x14\x22\x54\x99\x34\x88\x0b\xea\xa2\x6e\xa2\x31\xa1\x1a\xa5\x02\xa9\x0f\xc0\x05\x89\xaa\x89\x3f\x19\x83\x3d\x98\x3f\x92\x4d\xa8\xa2\x51\x82\x36\xd3\xb4\x61\x3a\xd6\x6e\x1a\x83\x2b\xb9\xb7\x66\x91\xc9\xe3\xda\x8b\x5b\xa8\xda\x69\x49\xd8\x7a\x1c\xf6\x42\x85\xea\x9f\x06\xd5\x15\x09\xdb\x07\x15\x9c\xac\xd0\x9e\xc0\xd5\xc6\x15\x7b\x3d\x34\x19\x83\xcb\xbb\xb7\x23\x5f\xf0\x7e\x93\x38\xe0\x95\xdb\x86\x69\xc3\x68\x5d\xb9\xe3\x68\x5c\xa8\x5e\x7a\x12\x5c\x35\x45\x14\x6d\xe1\x24\x65\xae\x5e\x4c\x1c\x88\x68\xed\x3b\xbe\x85\xad\x77\x3e\x81\x73\x94\x52\xf9\x00\x21\x0a\xaa\xb8\x1a\x61\x48\x42\x1e\xfa\x05\xb1\x33\x01\x50\x35\x66\xc6\xfd\x2f\x12\x0a\x54\xcd\xfd\xa5\x23\x05\x55\x81\x2e\x90\xb0\xe7\x04\xb6\x2b\x2a\x54\x07\x56\xc5\x19\x7f\x14\x2e\x54\x0b\x56\x26\xdf\xde\x88\xe0\xfa\x08\x9e\xa1\x92\x13\x85\x91\x0b\xef\x8d\x51\x0c\xe0\x42\x09\xdd\xcd\x6d\x6f\x60\x70\xc5\x84\x16\x50\x35\xd3\xa3\x4f\x2b\xe0\xd2\x09\x61\x41\xf7\xd4\xda\x84\x00\xef\xae\xeb\x8c\xf9\xb2\x2f\xec\xee\x86\x5a\x6b\x78\xf7\xc0\x05\x67\x00\x35\xee\x8e\xba\xc0\x92\x95\x78\x14\x2c\x54\x4f\xe1\x2c\xe7\x52\xa1\x45\x19\xa8\x2b\x0c\xc1\x81\x6a\xa6\xca\x4a\x31\x42\x8d\xb2\x52\xc6\x50\x65\x64\x8a\xf3\x63\x85\x33\xbe\x0c\xd5\x78\xdc\x01\xd2\x2f\xfd\xdb\xf7\x74\x91\x73\xfa\xc5\xbe\x76\x33\x19\xfb\x35\x7b\x87\xe7\xa7\xd9\xda\x1e\xb1\xa0\xbc\x94\x5a\x0b\xc8\xb8\x83\xe1\x64\x77\x33\x4c\xe2\x47\x82\x28\xe3\x69\xd4\xc4\x9c\xec\xae\x77\x0a\x5e\xd8\xd7\x26\x30\x32\x52\x8e\x82\xdd\x5d\xeb\x14\xa5\x5c\x99\xc0\x9a\x68\xd4\xc0\xe5\xde\xc5\xfd\xcd\xd7\x9b\x8b\xf3\x4f\x3e\xf0\x45\x33\x63\x77\xa9\x0f\x5b\x5f\xb0\xcc\xd6\xf2\xc8\xa4\x02\x6d\xdc\xb1\x9c\x74\xaf\x58\xbc\x34\xdc\xc9\xf4\xfd\xc0\x15\x4b\x20\x0d\xb7\xe7\x8a\xc5\xaa\xc8\xdf\x0c\x09\x11\xb2\x0a\x64\xd4\xf5\x00\xd9\x48\xa8\xdf\x70\x1a\x8f\xe5\x9f\xb3\x86\x9e\x31\xfa\xcd\xbd\x2b\xb4\x0b\x5e\x3d\x2c\x9f\xf8\xd2\x7f\x86\xa9\x3b\x26\x27\xde\x98\x0c\x5d\x7b\x4d\x03\xf5\x4a\xc3\x89\x0a\x34\x53\x44\x20\xb9\x66\x0a\xf7\x95\xd2\x05\x48\xeb\x04\x9e\x33\xe4\x10\xed\xd1\xd4\xba\x1c\x0a\xac\x56\x51\xa8\xc0\x1b\xa1\x26\x20\x3e\xe3\xcb\x23\xd3\xd4\x9e\x8f\xd4\xce\xca\xbf\x1e\xb4\x7f\xf2\x59\xf8\xfd\xac\xce\xc0\xbd\xf3\x16\xd3\xbb\xa1\x81\x7b\x07\xbd\x38\xd4\xc4\xdb\xb4\xbe\x82\x67\x34\x59\x6f\xbc\xf2\x75\xf7\x05\xab\xd5\x11\x7b\xcc\x17\xc3\xa1\x08\x53\xef\x36\x75\xda\x7d\xe2\xb5\xdb\xe9\x42\x98\x0a\xd5\x5b\x02\xa2\xa7\xde\xab\x61\xd3\xde\x10\x0c\xd7\xf1\x1c\x27\x0f\xdb\xfb\xf5\xe2\x30\xa6\xdd\x87\xad\xea\x7e\x6d\xdc\x5f\xb7\xbf\x4e\x87\xa7\x9e\x04\x4e\x87\xc2\xf5\x4e\xc1\x2f\x65\xcd\xbe\x9c\x5f\x5c\x6d\xfe\x0d\x3c\xb3\x4f\x03\xaf\x64\x6d\xdd\x8b\xaa\x73\xdd\x46\x31\xdc\xd6\x29\xcb\xd5\xd6\x0d\x95\xea\x1d\xa0\x63\xf7\xbd\x7f\x1b\x1d\x2b\xba\x5c\x99\xf7\x4d\x29\x37\x95\x07\x7f\xe7\x73\x24\xcb\x64\x85\xb0\xac\x62\x36\x29\x33\x7b\x5c\x15\xbb\x4d\xd9\x32\x94\x44\xa9\x3a\xb5\x80\x31\x7a\x5a\x35\x8e\x8e\x16\x0f\x50\x73\xa2\xf5\xa0\x44\xab\xc6\x2f\x79\x26\x49\x69\x5f\xa4\x30\x09\xd3\xf5\x7b\xc2\xbd\x05\x8f\x5d\x93\x76\xd1\x56\x93\x27\xd7\x5f\xd6\xb8\xcb\x4a\x50\x86\x2d\x8e\x80\x71\xb9\x26\xd7\xd8\x85\xbf\x3e\xda\x87\xac\xcf\x7d\x72\x07\x5a\xb5\x20\x81\x71\xf7\x55\x05\xbd\xfa\xe5\x03\x1f\xaf\xe0\x32\xd8\xac\x29\xa3\x77\x74\xbe\xec\x1a\x29\x7a\xe5\x06\xde\xa7\xee\x3e\x4d\x38\xb8\x78\xf5\x07\x50\x1b\xef\x1b\x4b\xc9\x82\x32\x92\x6e\x04\xc8\xb5\x7b\x82\x1e\x16\xff\x21\x39\x43\x5f\xd7\x05\xe9\x84\xda\x75\xb3\x81\xab\x5b\xaa\x0f\x3c\x5d\x9b\xf6\x3e\xa6\x9f\x8d\xd5\x7b\xe3\x55\xbd\xad\x69\x23\xec\x8c\x69\xfc\xe5\x6e\xe6\x75\x79\x0c\xbe\xd0\xfa\xc6\x70\xa9\x56\x5c\xd0\x3f\x49\x8a\xbe\x49\xd2\xcf\xc8\xb9\x6b\x67\xf3\xeb\xff\x9b\xe0\x94\xf8\xf2\x3b\x06\xfb\xb1\x8d\x3c\x8c\x10\x87\xe5\x67\xda\x7d\xc1\x6b\x33\x9d\xad\x7b\xdf\x47\x05\xfb\x8e\xfe\xe7\xc8\x69\x82\xa3\x9b\xd4\xf1\xb0\x05\x7f\xe3\x8b\x9f\xcb\x92\xa6\xbf\xa0\xef\x38\x2b\xfd\x71\x0c\x44\xb2\xf4\xe5\xd9\xd8\x29\x71\x6e\x5e\x64\x72\x75\xdb\x0b\x2c\xa5\xd5\xb4\x81\x01\x0d\x3d\xc2\xde\x57\x59\xc3\x3d\xf9\xda\x53\x50\xdc\x3d\xe9\xad\x87\xd0\xfd\xae\xae\x76\xf2\xe5\x6e\x66\x28\xb1\xf6\x8f\x29\x4f\x36\x53\x38\x09\xf8\xee\x90\x20\x36\x18\x3f\x48\x2a\x78\xee\x05\x46\xbf\xa6\xf3\x16\x67\xae\x04\xfb\x3f\x64\x28\xa9\x41\xcf\x55\x17\x94\x66\xba\x09\xcf\x8c\x1d\x9e\xcc\xa5\x1f\xf8\x02\x5d\x34\xd2\xe9\xa1\xeb\x2f\x90\x1f\xbc\x50\x2c\xe3\x6a\x65\x9e\xdd\xb3\xf3\xa9\x95\x0c\xda\x9b\x1f\x85\x70\xe7\x23\x63\x29\xd7\xed\xdf\xa0\x5f\x31\x55\x36\x62\x4d\xfd\x5d\x56\xe9\x40\xcd\xdd\x73\x8b\x4e\xb0\x8b\xe4\x5c\x4a\x9e\x50\x53\xa6\x5c\xcb\x28\xc1\x59\xd6\xeb\xa6\xae\x1a\x68\xa3\x41\x95\x42\x6b\x6a\xab\xeb\x4c\x51\x07\x59\x25\xa3\x06\x64\xda\x7d\x43\x40\xf2\x9c\x20\x45\xbd\x97\x79\x8e\x8f\xe1\x0a\xf7\x7b\xfb\x11\x8d\xea\x89\xc0\x8c\xe6\x81\x97\xdb\xed\xf0\x67\x19\x7f\x92\x88\xb3\x6c\x8d\x46\xd3\x77\xcd\xb3\x8a\x95\x43\xef\x0d\x2a\x6c\xb1\x2d\x5b\x26\x39\xf4\x64\x60\x97\x27\x97\xe2\x6e\x9e\xf1\x22\xa2\xca\x6f\xe0\xc2\x52\xd5\x65\x2e\xf4\xb6\x66\x0f\x73\x9c\x99\xf7\xaf\xe8\x07\x7e\x6d\x1e\x8a\x77\x29\xc5\x29\xc1\xdd\xc4\xbf\xe3\xd0\xf3\x85\x3d\xbd\x26\xf6\x19\x0f\xae\xec\x6d\xb7\x4d\x6f\xa2\xcd\x63\x4e\xad\x3e\xfd\xf7\x28\x7a\xfa\x34\xa2\xad\xaa\x6e\x01\x1e\xa5\x90\xdd\x27\x2c\xcd\x80\xd8\x69\xa1\xfb\x0a\x58\x79\x8d\xcc\xad\x89\x16\x58\x99\x63\xb8\x12\xbe\x27\xb2\xe0\x4c\xda\xf9\xc7\x4b\xb5\x11\x0e\x0b\x33\x83\x46\xbe\x19\x34\x1a\x34\x83\x46\x6f\xc1\xd5\xb3\x2f\x87\x2b\x67\xb7\x33\xcd\x75\x4b\x27\xf4\x42\x10\xd9\xf6\x82\x34\x22\x4b\xaa\x8a\xff\x26\x5d\xcd\x7d\xd2\x74\x62\x8a\xb7\xb0\xe6\xed\x4f\xc2\x54\xf3\xb6\x66\x4d\xfe\x0e\x7b\xdc\x3f\x4b\x22\xd6\x71\x9b\xdc\x0d\x5b\x64\xe5\xf3\xe5\x07\xb3\xf4\xcc\x40\x0c\xa8\xe7\xaa\xb1\x4f\x2a\x78\x91\x7d\xa5\x39\x69\xa2\xe5\xfa\x08\xae\xc2\xea\x94\x69\x4d\x04\xe5\x69\x53\xa6\x26\x40\x60\xe9\xa6\x96\xb5\x00\x8f\xf3\xc3\x69\x7e\x78\xac\xff\x5f\x1d\xbe\x5b\x1d\x1e\x8f\x56\x87\xa3\xc9\xea\xf0\x7d\x7a\x38\x7e\x9b\xb6\xe7\xde\xe7\xef\xb7\x57\x08\xa7\x39\x65\x5b\x72\x19\xfc\xc9\x37\x7e\xdb\x75\x21\xb4\x45\xd2\x7c\xb0\x5d\x24\x86\x88\x44\x65\xbd\xd2\xb8\xc0\xec\xef\x36\xf5\x84\x3d\x6a\xed\xab\x44\xd6\xb5\x34\x35\x1e\xf4\x2c\xfd\x89\x2f\x51\x81\x97\x7d\x79\x44\x2d\xb8\xcc\xb5\x0c\x80\x41\x0f\xcc\x17\xf5\x86\x39\x0c\xe7\xde\x6d\x69\x6d\xb0\xd2\x1b\xa8\xf3\xff\xe7\x81\xfa\xc1\xe3\xf0\x63\xe5\xdc\x2b\xc5\xff\x0b\x00\x00\xff\xff\x49\x87\xca\xe8\xf8\x89\x00\x00")

func ResourcesEventsYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesEventsYaml,
		"../resources/events.yaml",
	)
}

func ResourcesEventsYaml() (*asset, error) {
	bytes, err := ResourcesEventsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../resources/events.yaml", size: 35320, mode: os.FileMode(420), modTime: time.Unix(1605084227, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../resources/events.yaml": ResourcesEventsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"events.yaml": &bintree{ResourcesEventsYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
