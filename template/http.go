package template

import (
	"bytes"
	_ "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "os"
	_ "time"
)

type APIResponse struct {
	Headers http.Header
	Body    []byte
}

var baseUrl string = "https://api.vpsie.com/apps/v2"

func SendHttpRequest(path string, payload []byte) (*APIResponse, error) {
	url := baseUrl + path
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var headers http.Header

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	headers = resp.Header
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyAsString := string(body)
	fmt.Println("response Body:", bodyAsString)
	fmt.Println("response Status:", resp.Status)

	status := resp.StatusCode
	if !(status >= 200 && status <= 299) {
		return nil, errors.New("Status: " + resp.Status + " result: " + bodyAsString)
	}

	return &APIResponse{
		Headers: headers,
		Body:    body}, nil

}

func SendPutRequest(path string, payload []byte) (string, error) {
	url := baseUrl + path
	fmt.Println("URL:>", url)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyAsString := string(body)
	fmt.Println("response Body:", bodyAsString)
	fmt.Println("response Status:", resp.Status)
	status := resp.StatusCode
	if !(status >= 200 && status <= 299) {
		return "", errors.New("Status: " + resp.Status + " result: " + bodyAsString)
	}
	return bodyAsString, nil

}

func SendGetRequest(path string, vals map[string]string) (string, error) {
	fullUrl := baseUrl + path + "?"

	for k, v := range vals {
		fullUrl = fullUrl + (k + "=" + url.QueryEscape(v)) + "&"
	}
	fmt.Println("URL:>", fullUrl)

	req, err := http.NewRequest("GET", fullUrl, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyAsString := string(body)

	fmt.Println("response Body:", bodyAsString)
	fmt.Println("response Status:", resp.Status)
	status := resp.StatusCode
	if !(status >= 200 && status <= 299) {
		return "", errors.New("Status: " + resp.Status + " result: " + bodyAsString)
	}
	return bodyAsString, nil
}
