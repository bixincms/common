package common

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Download(url string, name string) (n int64, err error) {
	resp, err := http.Get(url)
	content_type := resp.Header.Get("Content-Type")
	split := strings.Split(content_type, "/")

	if len(split) == 2 {
		name = name[:strings.LastIndex(name, ".")] + "." + split[1]
	}
	out, err := os.Create(name)
	defer out.Close()
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}
