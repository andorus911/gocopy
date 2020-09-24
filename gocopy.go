package main

import (
	"github.com/cheggaaa/pb/v3"
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

	if limit < 1 {
		stat, err := fromFile.Stat()
		if err != nil {
			return err
		}
		limit = stat.Size()
	}

	reader := io.LimitReader(fromFile, limit)
	tmpl := `{{ green "Copying:" }} {{ bar . "[" "=" ">" "-" | green}} {{speed . | green }} {{percent . | green}}`
	bar := pb.ProgressBarTemplate(tmpl).Start64(limit)
	bar.Set("my_green_string", "green").Set("my_blue_string", "blue")
	barReader := bar.NewProxyReader(reader)

	bar.Start()
	_, err = io.Copy(toFile, barReader)
	if err != nil && err != io.EOF {
		return err
	}
	bar.Finish()

	err = toFile.Close()
	if err != nil {
		return err
	}
	return nil
}
