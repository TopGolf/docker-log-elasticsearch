package v2

import (
	"errors"
	"net/http"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v3"
)

// MyRetrier ...
type MyRetrier struct {
	backoff elastic.Backoff
}

// NewMyRetrier ...
func NewMyRetrier(timeout int) *MyRetrier {
	return &MyRetrier{
		backoff: elastic.NewExponentialBackoff(100*time.Millisecond, time.Duration(timeout)*time.Second),
	}
}

// Retry ...
func (r *MyRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	// Fail hard on a specific error
	if err == syscall.ECONNREFUSED {
		return 0, false, errors.New("Elasticsearch or network down")
	}

	// Let the backoff strategy decide how long to wait and whether to stop
	wait, stop := r.backoff.Next(retry)

	return wait, stop, nil
}
