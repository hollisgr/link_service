package link_service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (s *linkService) getFile(link string) (*bytes.Buffer, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("file download error: %v", err)
	}
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("file read err: %v", err)
	}
	return buffer, nil
}

func (s *linkService) createArchive(buffers []*bytes.Buffer, filenames []string) ([]byte, error) {
	archBuffer := new(bytes.Buffer)
	archiver := zip.NewWriter(archBuffer)
	for i, buffer := range buffers {
		writer, err := archiver.Create(filenames[i])
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(writer, buffer)
		if err != nil {
			return nil, err
		}
	}

	err := archiver.Close()
	if err != nil {
		return nil, fmt.Errorf("archiver close error: %v", err)

	}
	return archBuffer.Bytes(), nil
}
