package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Bodytojson(w io.Reader) (o map[string]interface{}, err error) {
	body, err := ioutil.ReadAll(w)
	if err != nil {
		return o, err
	}

	json.Unmarshal(body, &o)
	return o, nil
}
