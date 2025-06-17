package models

import "encoding/xml"

type Bucket struct {
	XMLName      xml.Name `xml:"bucket"`
	Name         string   `xml:"name"`
	CreationTime string   `xml:"creation_time"`
	LastModified string   `xml:"last_modified"`
	Status       string   `xml:"status"`
}

type Buckets struct {
	XMLName xml.Name `xml:"buckets"`
	Buckets []Bucket `xml:"bucket"`
}
