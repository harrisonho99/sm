package fetch

import (
	"io"
	"net/http"
	"os"
)

func Fetch(url string) {
	resp, err := http.Get(url)
	checkError(err)
	_, err = io.Copy(os.Stdout, resp.Body)
	checkError(err)
	resp.Body.Close()
}

func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
