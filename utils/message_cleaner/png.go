package message_cleaner

import (
	tExt "github.com/chrisbward/go-png-chunks"
)

func CleanPngTEXt(body []byte) ([]byte, error) {
	bodyBuffer, err := tExt.WritetEXtToPngBytes(body, tExt.TEXtChunk{})
	if err != nil {
		return nil, err
	}

	return bodyBuffer.Bytes(), nil
}
