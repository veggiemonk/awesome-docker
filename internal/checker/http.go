package checker

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/veggiemonk/awesome-docker/internal/cache"
)

const (
	defaultTimeout     = 30 * time.Second
	defaultConcurrency = 10
	userAgent          = "awesome-docker-checker/1.0"
)

// LinkResult holds the result of checking a single URL.
type LinkResult struct {
	URL         string
	RedirectURL string
	Error       string
	StatusCode  int
	OK          bool
	Redirected  bool
}

func shouldFallbackToGET(statusCode int) bool {
	switch statusCode {
	case http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusNotImplemented:
		return true
	default:
		return false
	}
}

// CheckLink checks a single URL. Uses HEAD first, falls back to GET.
func CheckLink(url string, client *http.Client) LinkResult {
	result := LinkResult{URL: url}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Track redirects
	var finalURL string
	origCheckRedirect := client.CheckRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		finalURL = req.URL.String()
		if len(via) >= 10 {
			return http.ErrUseLastResponse
		}
		return nil
	}
	defer func() { client.CheckRedirect = origCheckRedirect }()

	doRequest := func(method string) (*http.Response, error) {
		req, err := http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", userAgent)
		return client.Do(req)
	}

	resp, err := doRequest(http.MethodHead)
	if err != nil {
		resp, err = doRequest(http.MethodGet)
		if err != nil {
			result.Error = err.Error()
			return result
		}
	} else if shouldFallbackToGET(resp.StatusCode) {
		resp.Body.Close()
		resp, err = doRequest(http.MethodGet)
		if err != nil {
			result.Error = err.Error()
			return result
		}
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.OK = resp.StatusCode >= 200 && resp.StatusCode < 400

	if finalURL != "" && finalURL != url {
		result.Redirected = true
		result.RedirectURL = finalURL
	}

	return result
}

// CheckLinks checks multiple URLs concurrently.
func CheckLinks(urls []string, concurrency int, exclude *cache.ExcludeList) []LinkResult {
	if concurrency <= 0 {
		concurrency = defaultConcurrency
	}

	results := make([]LinkResult, len(urls))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, url := range urls {
		if exclude != nil && exclude.IsExcluded(url) {
			results[i] = LinkResult{URL: url, OK: true}
			continue
		}

		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			client := &http.Client{Timeout: defaultTimeout}
			results[idx] = CheckLink(u, client)
		}(i, url)
	}

	wg.Wait()
	return results
}
