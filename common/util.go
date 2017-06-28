package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Decode(r *http.Response, v interface{}) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("[ReadAll] ")
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		log.Fatal("[Unmarshal] ")
	}

	if valid, ok := v.(interface {
		OK() error
	}); ok {
		err = valid.OK()
		if err != nil {
			log.Fatal("[Validation] ")
		}
	}
}
