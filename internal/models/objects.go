package models

import "encoding/xml"

type Object struct {
	XMLName      xml.Name `xml:"object"`
	ObjectKey    string   `xml:"objectKey"`
	ContentType  string   `xml:"contentType"`
	Size         string   `xml:"size"`
	LastModified string   `xml:"lastModified"`
}

type Objects struct {
	XMLName xml.Name `xml:"objects"`
	Objects []Object `xml:"object"`
}