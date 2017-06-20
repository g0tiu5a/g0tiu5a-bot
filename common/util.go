package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Decode(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	//decode
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	//check
	if valid, ok := v.(interface {
		OK() error
	}); ok {
		err = valid.OK()
		if err != nil {
			return err
		}
	}

	return nil
}
