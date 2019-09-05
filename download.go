package common

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(url string, name string) (n int64, err error) {
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}
