package httpUtils

import (
//	"encoding/json"
	"net/http"
	"time"
	"log"
	"errors"
	"io"
)

// Fetch makes network calls using the method (POST/GET..), the URL // to hit, headers to add (if any), and the body of the request.
// Feel free to add more stuff to before/after making the actual n/w call!
func Fetch(method string, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	// Create client with required custom parameters.
	// Options: Disable keep-alives, 30sec n/w call timeout.
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Duration(10 * time.Second),
	}
	// Create request.
	req, _ := http.NewRequest(method, url, body)
	// Add any required headers.
	for key, value := range header {
		req.Header.Add(key, value)
	} 
	// Perform said network call.
	res, err := client.Do(req)
		if err != nil {
		log.Print(err.Error())   // use glog it's amazing
		return nil, err
	}
	// If response from network call is not 200, return error too.
	if res.StatusCode != http.StatusOK {
		return res, errors.New("Network call did not return SUCCESS!")
	}
	return res, nil
  }
