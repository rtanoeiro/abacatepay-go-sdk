package fetch

import "fmt"

type HTTPError struct {
	StatusCode int
	Body       []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("fetch: http %d: %s", e.StatusCode, e.Body)
}
