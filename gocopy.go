package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(from string, to string, offset int64, limit int64) error {
	fromFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.Create(to)
	if err != nil {
		return err
	}

	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	written, err := io.CopyN(toFile, fromFile, limit)
	if err != nil && err != io.EOF {
		return err
	}
	fmt.Println(written, "bytes were written.")

	err = toFile.Close()
	if err != nil {
		return err
	}
	return nil
}
