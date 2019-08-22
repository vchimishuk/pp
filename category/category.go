package category

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type SourceCategory struct {
	Id     int    `json:"id"`
	Parent int    `json:"parent_id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Ndocs  int    `json:"doc_count"`
}

type SourceResponse struct {
	Categories []*SourceCategory `json:"categories"`
}

func ParseResponse(r io.Reader) (*SourceResponse, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("IO error: %s", err)
	}
	var resp SourceResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("JSON parsing error: %s", err)
	}

	return &resp, nil
}
