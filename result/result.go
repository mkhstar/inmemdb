package result

import (
	"encoding/json"
)

const (
	//errorStatus for Info Status
	errorStatus string = "error"
	//successStatus  for Info Status
	successStatus string = "success"
)

// Info struct that holds the result after execution
type Info struct {
	Error  error
	Result interface{}
}

type resultJSONMap map[string]interface{}

func (r resultJSONMap) JSON() string {
	result, _ := json.Marshal(r)
	return string(result)
}

// Resulter can send response in available formats
type Resulter interface {
	Echo() string
	JSON() string
}
