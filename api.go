package todoist

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

const BaseUrl = "https://api.todoist.com/rest/v1/"

type Todoist struct {
	opts *Opts
}

type Opts struct {
	Token   string
	Client  *http.Client
	Timeout time.Duration
}

//goland:noinspection GoUnusedExportedFunction
func New(opts *Opts) *Todoist {
	if opts.Timeout == 0 {
		opts.Timeout = 15 * time.Second
	}

	if opts.Client == nil {
		opts.Client = &http.Client{
			Timeout: opts.Timeout,
		}
	}

	return &Todoist{
		opts: opts,
	}
}

func (t *Todoist) request(ctx context.Context, method string, endpoint string, params map[string]string, payload io.Reader, data interface{}) (err error) {
	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, method, BaseUrl+endpoint, payload); err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+t.opts.Token)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if params != nil && len(params) != 0 {
		query := req.URL.Query()
		for key, value := range params {
			query.Set(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	var res *http.Response
	if res, err = t.opts.Client.Do(req); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNoContent:
		return
	case http.StatusOK:
		if res.Header.Get("Content-Type") != "application/json" {
			return errors.New("invalid response content type")
		}

		if err = json.NewDecoder(res.Body).Decode(data); err != nil {
			return
		}

		return
	default:
		return errors.New(res.Status)
	}
}
