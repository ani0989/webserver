package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://127.0.0.1:5000/companies"

// Client structure returned
type Client struct {
	Username string
	Password string
}

// NewBasicAuthClient ...Returns the client to authenticate
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// Todo structure returned
type Todo struct {
	ID int `json:"id"`
}

// PostReq ... Makes the post req
func (s *Client) PostReq(todo *Todo) error {
	fmt.Println(baseURL)
	fmt.Println("Post Request")
	j, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err

}

// GetReq ... Makes the get req
func (s *Client) GetReq(todo *Todo) error {
	fmt.Println(baseURL)
	fmt.Println("Get Request")
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
