package sportsblaze

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		opts     Options
		testFunc func(t *testing.T, c *Client)
	}{
		{
			name: "happy path with default base path",
			opts: Options{},
			testFunc: func(t *testing.T, c *Client) {
				assert.Equal(t, defaultBasePath, c.basePath)
				assert.NotNil(t, c.httpc)
			},
		},
		{
			name: "happy path with custom base path",
			opts: Options{BasePath: "https://custom.api.com"},
			testFunc: func(t *testing.T, c *Client) {
				assert.Equal(t, "https://custom.api.com", c.basePath)
				assert.NotNil(t, c.httpc)
			},
		},
		{
			name: "happy path with api key",
			opts: Options{ApiKey: "some_key"},
			testFunc: func(t *testing.T, c *Client) {
				assert.Equal(t, "some_key", c.apiKey)
				assert.NotNil(t, c.httpc)
			},
		},
	}
	for _, tc := range tests {
		c := New(tc.opts)
		tc.testFunc(t, c)
	}
}

func TestDailyBoxScores(t *testing.T) {
	server := mockServer()
	defer server.Close()
	tests := []struct {
		name     string
		input    string
		client   *Client
		testFunc func(t *testing.T, dbs *DailyBoxScores, err error)
	}{
		{
			name:   "happy path",
			input:  "2024-10-24",
			client: New(Options{BasePath: server.URL}),
			testFunc: func(t *testing.T, dbs *DailyBoxScores, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, dbs)
			},
		},
		{
			name:   "no date provided",
			input:  "",
			client: New(Options{BasePath: server.URL}),
			testFunc: func(t *testing.T, dbs *DailyBoxScores, err error) {
				assert.Error(t, err)
				assert.Empty(t, dbs)
			},
		},
		{
			name:   "invalid date format, not YYYY-MM-DD",
			input:  "10-24-2024",
			client: New(Options{}),
			testFunc: func(t *testing.T, dbs *DailyBoxScores, err error) {
				assert.Error(t, err)
				assert.Empty(t, dbs)
			},
		},
		{
			name:   "http error",
			input:  "2024-10-24",
			client: New(Options{BasePath: server.URL + "/negative/http"}),
			testFunc: func(t *testing.T, dbs *DailyBoxScores, err error) {
				assert.Error(t, err)
				assert.Empty(t, dbs)
			},
		},
		{
			name:   "response unmarshal error",
			input:  "2024-10-24",
			client: New(Options{BasePath: server.URL + "/negative/json"}),
			testFunc: func(t *testing.T, dbs *DailyBoxScores, err error) {
				assert.Error(t, err)
				assert.Empty(t, dbs)
			},
		},
	}
	for _, tc := range tests {
		dbs, err := tc.client.DailyBoxScores(tc.input)
		tc.testFunc(t, dbs, err)
	}
}

func TestAddApiKey(t *testing.T) {
	c := New(Options{ApiKey: "some_key"})
	tests := []struct {
		name     string
		input    string
		testFunc func(t *testing.T, got string, err error)
	}{
		{
			name:  "happy path",
			input: "https://api.sportsblaze.com/nba/v1/daily/boxscores",
			testFunc: func(t *testing.T, got string, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "https://api.sportsblaze.com/nba/v1/daily/boxscores?key=some_key", got)
			},
		},
		{
			name:  "bad url",
			input: "https://this is a bad url",
			testFunc: func(t *testing.T, got string, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range tests {
		got, err := c.addApiKey(tc.input)
		tc.testFunc(t, got, err)
	}
}

func mockServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/boxscores/daily/2024-10-24.json" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{
				"date": "2024-10-24",
				"games": [
					{
						"gameId": "001",
						"homeTeam": "Lakers",
						"awayTeam": "Warriors",
						"homeScore": 102,
						"awayScore": 99
					}
				]
			}`)
			return
		}
		if r.URL.Path == "/negative/boxscores/daily/2024-10-24.json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if r.URL.Path == "/negative/json/boxscores/daily/2024-10-24.json" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `this really is not json`)
			return
		}
	})
	return httptest.NewServer(handler)
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		testFunc func(t *testing.T, err error)
	}{
		{
			name:  "happy path",
			input: "2024-10-24",
			testFunc: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:  "empty date",
			input: "",
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "invalid format",
			input: "10-24-2024",
			testFunc: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range tests {
		err := validateDate(tc.input)
		tc.testFunc(t, err)
	}
}
