package utils

func CreateObjectsCSV(path string) error {
	return WriteFileWithHeader(path, objectsHeader)
}