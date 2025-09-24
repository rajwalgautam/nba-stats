package sportsblaze

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	defaultBasePath = "https://api.sportsblaze.com/nba/v1"
	dateLayout      = "2006-01-02"
)

type Options struct {
	ApiKey   string
	BasePath string
}

type Client struct {
	apiKey   string
	basePath string
	httpc    *http.Client
}

func New(opts Options) *Client {
	hc := &http.Client{Timeout: 20 * time.Second}
	c := &Client{
		apiKey: opts.ApiKey,
		httpc:  hc,
	}
	if opts.BasePath != "" {
		c.basePath = opts.BasePath
	} else {
		c.basePath = defaultBasePath
	}
	return c
}

func (c *Client) DailyBoxScores(date string) (*DailyBoxScores, error) {
	if err := validateDate(date); err != nil {
		return &DailyBoxScores{}, err
	}

	fullPath := fmt.Sprintf("%s/boxscores/daily/%s.json", c.basePath, date)

	b, err := c.get(fullPath)
	if err != nil {
		return &DailyBoxScores{}, err
	}

	dbs := new(DailyBoxScores)
	err = json.Unmarshal(b, dbs)
	if err != nil {
		return &DailyBoxScores{}, err
	}

	return dbs, nil
}

func (c *Client) get(path string) ([]byte, error) {
	withKey, err := c.addApiKey(path)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpc.Get(withKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s - got %d, expected 200", path, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s - error reading response: %+v", path, err)
	}
	return body, nil
}

func (c *Client) addApiKey(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("key", c.apiKey)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func validateDate(t string) error {
	_, err := time.Parse(dateLayout, t)
	return err
}
