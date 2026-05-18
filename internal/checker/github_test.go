package checker

import (
	"errors"
	"testing"
)

func TestExtractGitHubRepo(t *testing.T) {
	tests := []struct {
		url   string
		owner string
		name  string
		ok    bool
	}{
		{"https://github.com/docker/compose", "docker", "compose", true},
		{"https://github.com/moby/moby", "moby", "moby", true},
		{"https://github.com/user/repo/", "user", "repo", true},
		{"https://github.com/user/repo?tab=readme-ov-file", "user", "repo", true},
		{"https://github.com/user/repo#readme", "user", "repo", true},
		{"https://github.com/user/repo.git", "user", "repo", true},
		{"https://www.github.com/user/repo", "user", "repo", true},
		{"https://github.com/user/repo/issues", "", "", false},
		{"https://github.com/user/repo/wiki", "", "", false},
		{"https://github.com/apps/dependabot", "", "", false},
		{"https://example.com/not-github", "", "", false},
		{"https://github.com/user", "", "", false},
	}

	for _, tt := range tests {
		owner, name, ok := ExtractGitHubRepo(tt.url)
		if ok != tt.ok {
			t.Errorf("ExtractGitHubRepo(%q): ok = %v, want %v", tt.url, ok, tt.ok)
			continue
		}
		if ok {
			if owner != tt.owner || name != tt.name {
				t.Errorf(
					"ExtractGitHubRepo(%q) = (%q, %q), want (%q, %q)",
					tt.url,
					owner,
					name,
					tt.owner,
					tt.name,
				)
			}
		}
	}
}

func TestPartitionLinks(t *testing.T) {
	urls := []string{
		"https://github.com/docker/compose",
		"https://example.com/tool",
		"https://github.com/moby/moby",
		"https://github.com/user/repo/issues",
		"dozzle",
		"#projects",
	}
	gh, ext := PartitionLinks(urls)
	if len(gh) != 2 {
		t.Errorf("github links = %d, want 2", len(gh))
	}
	if len(ext) != 2 {
		t.Errorf("external links = %d, want 2", len(ext))
	}
}

func TestIsGitHubAuthError(t *testing.T) {
	tests := []struct {
		err  error
		want bool
	}{
		{errors.New("non-200 OK status code: 401 Unauthorized body: \"Bad credentials\""), true},
		{errors.New("Resource not accessible by integration"), true},
		{errors.New("dial tcp: lookup api.github.com: no such host"), false},
		{errors.New("context deadline exceeded"), false},
	}

	for _, tt := range tests {
		got := isGitHubAuthError(tt.err)
		if got != tt.want {
			t.Errorf("isGitHubAuthError(%q) = %v, want %v", tt.err, got, tt.want)
		}
	}
}
