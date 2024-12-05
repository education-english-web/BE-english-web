package fileutil

import (
	"path/filepath"
	"strings"
)

const (
	mMIMETypeUnknown = "Unknown"
)

var mMIMETypeMap = map[string]string{
	"pdf":  "application/pdf",
	"doc":  "application/msword",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"xls":  "application/vnd.ms-excel",
	"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"ppt":  "application/vnd.ms-powerpoint",
	"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"odt":  "application/vnd.oasis.opendocument.text",
	"ods":  "application/vnd.oasis.opendocument.spreadsheet",
	"odp":  "application/vnd.oasis.opendocument.presentation",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"png":  "image/png",
	"gif":  "image/gif",
	"txt":  "text/plain",
	"csv":  "text/csv",
	"zip":  "application/zip",
}

// ExtractFromFilename returns name and file extension(lower case) from filename
func ExtractFromFilename(filePath string) (string, string) {
	var name, ext string

	baseFile := filepath.Base(filePath)
	if baseFile == "." {
		return "", ""
	}

	var i int

	for i = len(baseFile) - 1; i > -1; i-- {
		if baseFile[i] == '.' {
			break
		}
	}

	switch {
	case i < 0:
		name, ext = baseFile, ""
	case i == 0:
		name, ext = "", baseFile[1:]
	case i > 0:
		name, ext = baseFile[:i], baseFile[i+1:]
	}

	return name, strings.ToLower(ext)
}

// GetMIMEType return MIME type from Ext
func GetMIMEType(ext string) string {
	if mimeType, ok := mMIMETypeMap[ext]; ok {
		return mimeType
	}

	return mMIMETypeUnknown
}
