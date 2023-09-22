package checks

import (
	"context"
	"fmt"

	"github.com/compscore/compscore/pkg/structs"
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

func GetAllGitRemotes(runningConfig structs.RunningConfig_s) []structs.Release_s {
	remoteMap := make(map[structs.Release_s]bool)

	for _, team := range runningConfig.Teams {
		for _, check := range team.Checks {
			remoteMap[check.Release] = true
		}
	}

	remoteSlice := make([]structs.Release_s, len(remoteMap))
	i := 0
	for remote := range remoteMap {
		remoteSlice[i] = remote
		i++
	}

	return remoteSlice
}
