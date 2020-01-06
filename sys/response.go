package sys

import "encoding/json"

type Response struct {
	Info    bool
	Message string
}

func (res *Response) Parse() []byte {
	json, err := json.Marshal(res)
	if err != nil {
		return nil
	}

	return json
}
