package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
	"triple-s/config"
)

var (
	BucketsHeader = []string{"Name", "CreationTime", "LastModifiedTime", "Status"}
	ObjectsHeader = []string{"ObjectKey", "Size", "ContentType", "LastModified"}
)

func InitDir() error {
	err := os.MkdirAll(config.Dir, 0755)
	if err != nil {
		return err
	}

	return WriteFileWithHeader(path.Join(config.Dir, "buckets.csv"), BucketsHeader)
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
		if err != nil {
			return err
		}
	} else {
		var n int

		for i := range header {
			n += len(header[i])
		}

		n += len(header) - 1
		buf := make([]byte, n)
		file.Read(buf)
		if string(buf) != strings.Join(header, ",") {
			return fmt.Errorf("incorrect header in %s", path)
		}
	}

	return nil
}
