package render

import (
	"bytes"
	"fmt"
	"os"
)

func Write(filename string, buf *bytes.Buffer) error {
	if buf == nil {
		return fmt.Errorf("buffer cannot be nil")
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	_, err = file.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}
	return nil
}
