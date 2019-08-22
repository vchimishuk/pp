package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/vchimishuk/opt"
	"github.com/vchimishuk/pp/category"
	"github.com/vchimishuk/pp/search"
)

type Response struct {
	Categories []*search.ClientCategory `json:"categories"`
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	descs := []*opt.Desc{
		{"h", "host", opt.ArgString, "addr", "net address to bind to"},
		{"p", "port", opt.ArgInt, "path", "port number to listen on"},
		{"s", "source", opt.ArgString, "url", "source server URL"},
	}
	opts, _, err := opt.Parse(os.Args[1:], descs)
	if err != nil {
		die("%s", err)
	}

	src := opts.StringOr("source", "http://localhost:8888/categories")
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("category_id")
		if idStr == "" {
			idStr = "0"
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := http.Get(src)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		scResp, err := category.ParseResponse(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		roots := search.Tree(scResp.Categories, id)
		b, err := json.Marshal(&Response{roots})
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		_, err = w.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"error: filed to write response: %s", err)
		}
	})
	addr := fmt.Sprintf("%s:%d", opts.StringOr("host", "localhost"),
		opts.IntOr("port", 8889))
	http.ListenAndServe(addr, nil)
}
