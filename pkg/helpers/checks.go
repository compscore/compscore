package helpers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/compscore/compscore/pkg/config"
	"github.com/google/go-github/github"
)

func GetReleaseAsset(organization string, repo string, tag string) (path string, err error) {
	if tag == "" {
		tag = "latest"
	}

	filename := "plugins/" + organization + "-" + repo + "-" + tag + ".so"

	exists, err := FileExists(filename)
	if err != nil {
		return "", err
	}

	if exists {
		return filename, nil
	}

	return DownloadReleaseAsset(organization, repo, tag)
}

func DownloadReleaseAsset(organization string, repo string, tag string) (filepath string, err error) {
	githubClient := github.NewClient(nil)

	var release *github.RepositoryRelease

	if tag == "latest" {
		release, _, err = githubClient.Repositories.GetLatestRelease(context.Background(), organization, repo)
	} else {
		release, _, err = githubClient.Repositories.GetReleaseByTag(context.Background(), organization, repo, tag)
	}

	if err != nil {
		return "", err
	}

	var asset *github.ReleaseAsset

	for _, a := range release.Assets {
		if a.GetName() == config.CheckFileName {
			asset = &a
			break
		}
	}

	if asset == nil {
		return "", fmt.Errorf("failed to find asset")
	}

	resp, err := http.Get(asset.GetBrowserDownloadURL())

	filepath = "plugins/" + organization + "-" + repo + "-" + tag + ".so"

	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}
