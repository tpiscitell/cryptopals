package utils

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func ReadBase64File(fname string) []byte {
	f, _ := os.Open(fname)
	defer f.Close()

	reader := base64.NewDecoder(base64.StdEncoding, f)
	data, _ := ioutil.ReadAll(reader)

	return data
}
