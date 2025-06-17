package config

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	bucketsHeader = []string{"Name", "CreationTime", "LastModifiedTime", "Status"}
	objectsHeader = []string{"ObjectKey", "Size", "ContentType", "LastModified"}
)

func InitDir() error {
	err := os.MkdirAll(Dir, 0755)
	if err != nil {
		return err
	}

	return WriteFileWithHeader(path.Join(Dir, "buckets.csv"), bucketsHeader)
}

func WriteFileWithHeader(path string, header []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	if fileStat.Size() == 0 {
		_, err := file.WriteString(strings.Join(header, ",") + "\n")
		fmt.Println("Header for 'buckets.csv' written successfully!")
		if err != nil {
			return err
		}
	} else {
		var n int

		for i := range bucketsHeader {
			n += len(bucketsHeader[i])
		}

		n += len(bucketsHeader) - 1
		buf := make([]byte, n)
		file.Read(buf)
		if string(buf) != strings.Join(header, ",") {
			return fmt.Errorf("incorrect header in %s", path)
		}
	}

	return nil
}
