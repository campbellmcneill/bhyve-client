package bhyveconnect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.orbitbhyve.com/v1/session"

// Client with username and password
type Client struct {
	Username string
	Password string
}

// NewBasicAuthClient creates the type
func NewBasicAuthClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// HelloWorld test
func HelloWorld(){
	fmt.Printf("Hello World\n")
}

// SessionToken returned once authenticated
type SessionToken struct {
	user_id       int    `json:"user_id"`
	orbit_api_key string `json:"orbit_api_key"`
}

// Login to BHyve
func (s *SessionToken) Login(client *Client) error {
	url := fmt.Sprintf(baseURL)
	fmt.Println(url)
	j, err := json.Marshal(client)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.SetBasicAuth(client.Username, client.Password)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

// doRequest sends an HTTP request
func (s *SessionToken) doRequest(req *http.Request) ([]byte, error) {

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
