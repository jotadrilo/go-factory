package generator

import (
	"bytes"
	"os"
	"path/filepath"
)

type DataWriter interface {
	Write(path string, data []byte) error
}

type FileWriter struct{}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

func (w *FileWriter) Write(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

type BufferWriter struct {
	b *bytes.Buffer
}

func NewBufferWriter(b *bytes.Buffer) *BufferWriter {
	return &BufferWriter{b: b}
}

func (w *BufferWriter) Write(_ string, data []byte) error {
	w.b.Write(data)
	return nil
}
