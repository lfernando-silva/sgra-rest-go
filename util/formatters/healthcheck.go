package formatters

import (
	"encoding/json"
	"io/ioutil"

	//"net/http"
	"net/http/httptest"

	"github.com/lfernando-silva/sgra-rest-go/util/types"
)

func jsonUnmarshal(responseData []byte, r *types.THealthCheck) {
	json.Unmarshal(responseData, &r)
}

// func parseResponse(response *http.Response, r *types.THealthCheck) error{
// 	responseData, err := ioutil.ReadAll(response.Body)
//     jsonUnmarshal(responseData, r)
// 	return err
// }

func ParseTestResponse(response *httptest.ResponseRecorder, r *types.THealthCheck) {
	responseData, _ := ioutil.ReadAll(response.Body)
    jsonUnmarshal(responseData, r)
}