package csv

import (
	"encoding/csv"
	"os"
)

type Writer struct {
	header    []string
	autoFlush bool
	file      *os.File
	writer    *csv.Writer
}

func NewWriter(file string, header []string) (*Writer, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(f)
	err = writer.Write(header)
	if err != nil {
		return nil, err
	}
	return &Writer{header: header, autoFlush: false, file: f, writer: writer}, nil
}

func (w *Writer) SetAutoFlush(autoFlush bool) {
	w.autoFlush = autoFlush
}

func (w *Writer) Write(record []string) error {
	err := w.writer.Write(record)
	if err != nil {
		return err
	}
	if w.autoFlush {
		w.writer.Flush()
	}
	return nil
}

func (w *Writer) Flush() {
	w.writer.Flush()
}

func (w *Writer) Close() error {
	w.writer.Flush()
	return w.file.Close()
}
