package checks

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func GetLatestRelease(organization string, repo string) (tag string, err error) {
	githubClient := github.NewClient(nil)

	release, response, err := githubClient.Repositories.GetLatestRelease(context.Background(), organization, repo)
	if response.StatusCode != 200 {
		return "", fmt.Errorf("release endpoint returned %d for: %s/%s", response.StatusCode, organization, repo)
	}

	if err != nil {
		return "", err
	}

	return release.GetTagName(), nil
}
