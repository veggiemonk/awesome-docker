package checker

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// RepoInfo holds metadata about a GitHub repository.
type RepoInfo struct {
	Owner      string
	Name       string
	URL        string
	IsArchived bool
	IsDisabled bool
	IsPrivate  bool
	PushedAt   time.Time
	Stars      int
	Forks      int
	HasLicense bool
}

// ExtractGitHubRepo extracts owner/name from a GitHub URL.
// Returns false for non-repo URLs (issues, wiki, apps, etc.).
func ExtractGitHubRepo(rawURL string) (owner, name string, ok bool) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", false
	}

	host := strings.ToLower(u.Hostname())
	if host != "github.com" && host != "www.github.com" {
		return "", "", false
	}

	path := strings.Trim(u.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}

	// Skip known non-repository top-level routes.
	switch parts[0] {
	case "apps", "features", "topics":
		return "", "", false
	}

	name = strings.TrimSuffix(parts[1], ".git")
	if name == "" {
		return "", "", false
	}

	return parts[0], name, true
}

func isHTTPURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

func isGitHubAuthError(err error) bool {
	if err == nil {
		return false
	}
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "401 unauthorized") ||
		strings.Contains(s, "bad credentials") ||
		strings.Contains(s, "resource not accessible by integration")
}

// PartitionLinks separates URLs into GitHub repos and external HTTP(S) links.
func PartitionLinks(urls []string) (github, external []string) {
	for _, url := range urls {
		if _, _, ok := ExtractGitHubRepo(url); ok {
			github = append(github, url)
		} else if isHTTPURL(url) {
			external = append(external, url)
		}
	}
	return
}

// GitHubChecker uses the GitHub GraphQL API.
type GitHubChecker struct {
	client *githubv4.Client
}

// NewGitHubChecker creates a checker with the given OAuth token.
func NewGitHubChecker(token string) *GitHubChecker {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	return &GitHubChecker{client: githubv4.NewClient(httpClient)}
}

// CheckRepo queries a single GitHub repository.
func (gc *GitHubChecker) CheckRepo(ctx context.Context, owner, name string) (RepoInfo, error) {
	var query struct {
		Repository struct {
			IsArchived     bool
			IsDisabled     bool
			IsPrivate      bool
			PushedAt       time.Time
			StargazerCount int
			ForkCount      int
			LicenseInfo    *struct {
				Name string
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	vars := map[string]any{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	if err := gc.client.Query(ctx, &query, vars); err != nil {
		return RepoInfo{}, fmt.Errorf("github query %s/%s: %w", owner, name, err)
	}

	r := query.Repository
	return RepoInfo{
		Owner:      owner,
		Name:       name,
		URL:        fmt.Sprintf("https://github.com/%s/%s", owner, name),
		IsArchived: r.IsArchived,
		IsDisabled: r.IsDisabled,
		IsPrivate:  r.IsPrivate,
		PushedAt:   r.PushedAt,
		Stars:      r.StargazerCount,
		Forks:      r.ForkCount,
		HasLicense: r.LicenseInfo != nil,
	}, nil
}

// CheckRepos queries multiple repos in sequence with rate limiting.
func (gc *GitHubChecker) CheckRepos(ctx context.Context, urls []string, batchSize int) ([]RepoInfo, []error) {
	if batchSize <= 0 {
		batchSize = 50
	}

	var results []RepoInfo
	var errs []error

	for i, url := range urls {
		owner, name, ok := ExtractGitHubRepo(url)
		if !ok {
			continue
		}

		info, err := gc.CheckRepo(ctx, owner, name)
		if err != nil {
			errs = append(errs, err)
			if isGitHubAuthError(err) {
				break
			}
			continue
		}
		results = append(results, info)

		if (i+1)%batchSize == 0 {
			time.Sleep(1 * time.Second)
		}
	}

	return results, errs
}
