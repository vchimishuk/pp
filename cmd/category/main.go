package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/vchimishuk/opt"
	"github.com/vchimishuk/pp/category"
)

const defaultFile = "categories.json"

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func parseCategories(r io.Reader) ([]*category.SourceCategory, error) {
	var cats []*category.SourceCategory

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &cats)

	return cats, err
}

func main() {
	descs := []*opt.Desc{
		{"f", "file", opt.ArgString, "path", "categories file to serve"},
		{"h", "host", opt.ArgString, "addr", "net address to bind to"},
		{"p", "port", opt.ArgInt, "path", "port number to listen on"},
	}
	opts, _, err := opt.Parse(os.Args[1:], descs)
	if err != nil {
		die("%s", err)

	}

	fname := opts.StringOr("file", defaultFile)
	f, err := os.Open(fname)
	if err != nil {
		die("failed to open %s: %s", fname, err)
	}
	cats, err := parseCategories(f)
	if err != nil {
		die("failed to parse %s: %s", fname, err)
	}

	data, err := json.Marshal(category.SourceResponse{cats})
	f.Close()
	if err != nil {
		die("response serialization failed: %s", err)
	}
	dataLen := strconv.Itoa(len(data))

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", dataLen)
		_, err := w.Write(data)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"error: filed to write response: %s", err)
		}
	})
	addr := fmt.Sprintf("%s:%d", opts.StringOr("host", "localhost"),
		opts.IntOr("port", 8888))
	http.ListenAndServe(addr, nil)
}
