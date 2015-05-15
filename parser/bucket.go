package parser

import (
	"encoding/xml"
)

type ListBucketResult struct {
	XMLName     xml.Name   `xml:"ListBucketResult"`
	Contents    []Contents `xml:"Contents"`
	IsTruncated string     `xml:"IsTruncated"`
	MaxKeys     string     `xml:"MaxKeys"`
	Name        string     `xml:"Name"`
}

type Contents struct {
	ETag         string `xml:"ETag"`
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	Owner        struct {
		DisplayName string `xml:"DisplayName"`
		ID          string `xml:"ID"`
	} `xml:"Owner"`
	Size         string `xml:"Size"`
	StorageClass string `xml:"StorageClass"`
	Type         string `xml:"Type"`
}

// <ListBucketResult>
//   <Name>rijin-develop-components-biz</Name>
//   <Prefix></Prefix>
//   <Marker></Marker>
//   <MaxKeys>100</MaxKeys>
//   <Delimiter></Delimiter>
//   <IsTruncated>false</IsTruncated>
//   <Contents>
//     <Key>asset/277cbccb-9107-40f8-6c0c-695dacc935a5</Key>
//     <LastModified>2015-05-15T05:43:20.000Z</LastModified>
//     <ETag>"A0326A955CE30068C4CCC7C2FC5BBA52"</ETag>
//     <Type>Normal</Type>
//     <Size>69796</Size>
//     <StorageClass>Standard</StorageClass>
//     <Owner>
//       <ID>1772225470167122</ID>
//       <DisplayName>1772225470167122</DisplayName>
//     </Owner>
//   </Contents>
//   <Contents>
//     <Key>asset/277cbccb-9107-40f8-6c0c-695dacc935a5</Key>
//     <LastModified>2015-05-15T05:43:20.000Z</LastModified>
//     <ETag>"A0326A955CE30068C4CCC7C2FC5BBA52"</ETag>
//     <Type>Normal</Type>
//     <Size>69796</Size>
//     <StorageClass>Standard</StorageClass>
//     <Owner>
//       <ID>1772225470167122</ID>
//       <DisplayName>1772225470167122</DisplayName>
//     </Owner>
//   </Contents>
// </ListBucketResult>
