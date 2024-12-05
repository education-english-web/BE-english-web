package imageutil

import (
	"bytes"
	"encoding/base64"
	"image"
	// Register image decoder
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

func ExtractFromImage(base64Image string) (image.Config, string, error) {
	var (
		conf image.Config
		ext  string
	)

	ss := strings.Split(base64Image, ",")
	b64 := ss[0]

	if len(ss) > 1 {
		b64 = ss[1]
	}

	unbased, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return conf, ext, err
	}

	r := bytes.NewReader(unbased)

	conf, ext, err = image.DecodeConfig(r)
	if err != nil {
		return conf, ext, err
	}

	return conf, ext, nil
}
