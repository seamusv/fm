package http

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/seamusv/fm-integration/encoding"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type (
	Client struct {
		client   http.Client
		url      string
		username string
		password string
		err      error
	}

	ClientBuilder func() *Client
)

func NewClient(url, username, password string) ClientBuilder {
	return func() *Client {
		jar, _ := cookiejar.New(nil)
		return &Client{
			client:   http.Client{Jar: jar},
			url:      url,
			username: username,
			password: password,
		}
	}
}

func (c *Client) Login(profile, organisation string, businessDate time.Time) {
	if c.err != nil {
		return
	}
	params := fmt.Sprintf("ffff%s", hexString("Logon_user="+c.username)) +
		fmt.Sprintf("ffff%s", hexString("Logon_password="+c.password)) +
		fmt.Sprintf("ffff%s", hexString("Logon_newpassword=")) +
		fmt.Sprintf("ffff%s", hexString("Logon_profile="+profile)) +
		fmt.Sprintf("ffff%s", hexString("Logon_org="+organisation)) +
		fmt.Sprintf("ffff%s", hexString("Logon_bdate="+businessDate.Format("02/01/2006")))
	request := transDocument{
		Connection: connectionDocument{
			Cmd:    "connect",
			Params: params,
		},
	}
	xmlBytes, _ := xml.Marshal(request)
	res, err := c.post(xmlBytes)
	if err != nil {
		c.err = err
	} else {
		if res.Body != nil {
			defer res.Body.Close()
		}
	}
}

func (c *Client) Logout() {
	tmpErr := c.err
	c.err = nil

	c.Execute("EXIT")

	request := transDocument{
		Connection: connectionDocument{
			Cmd: "disconnect",
		},
	}
	xmlBytes, _ := xml.Marshal(request)
	res, err := c.post(xmlBytes)
	if err != nil {
		c.err = err
	} else {
		if res.Body != nil {
			defer res.Body.Close()
		}
	}
	c.err = tmpErr
}

func (c *Client) Execute(command string, messageCodes ...string) *encoding.Response {
	if c.err != nil {
		return nil
	}

	return c.ExecuteFields(command, struct{}{}, messageCodes...)
}

func (c *Client) ExecuteFields(command string, v interface{}, messageCodes ...string) *encoding.Response {
	if c.err != nil {
		return nil
	}

	xmlBytes, err := encoding.Marshal(command, v)
	if err != nil {
		c.err = err
		return nil
	}
	res, err := c.post(xmlBytes)
	if err != nil {
		c.err = err
		return nil
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)
	response, err := encoding.Parse(body)
	if err != nil {
		c.err = err
		return nil
	}
	if err := response.MessageContainsOneOf(messageCodes...); err != nil {
		c.err = err
		return nil
	}

	return response
}

func (c *Client) Err() error {
	return c.err
}

func (c *Client) post(data []byte) (*http.Response, error) {
	return c.client.Post(c.url, "text/xml", bytes.NewReader(data))
}

func hexString(s string) string {
	runes := []rune(s)
	result := ""
	for _, r := range runes {
		result += fmt.Sprintf("%04x", int(r))
	}
	return result
}

type (
	transDocument struct {
		XMLName    xml.Name `xml:"trans"`
		Connection connectionDocument
	}

	connectionDocument struct {
		XMLName xml.Name `xml:"connection"`
		Cmd     string   `xml:"cmd,attr"`
		Params  string   `xml:"parms,attr"`
	}
)
