package checks

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"plugin"
	"strings"

	"github.com/compscore/compscore/pkg/helpers"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/google/go-github/github"
)

type releaseAssetCacheStruct struct {
	Path string
	Tag  string
}

var (
	releaseAssetCache  map[string]releaseAssetCacheStruct                                                                                                                                          = make(map[string]releaseAssetCacheStruct)
	checkFunctionCache map[string]func(ctx context.Context, target string, command string, expectedOutput string, username string, password string, options map[string]interface{}) (bool, string) = make(map[string]func(ctx context.Context, target string, command string, expectedOutput string, username string, password string, options map[string]interface{}) (bool, string))
)

func GetCheckFunction(organization string, repo string, tag string) (func(ctx context.Context, target string, command string, expectedOutput string, username string, password string, options map[string]interface{}) (bool, string), error) {
	checkFunction, ok := checkFunctionCache[organization+"/"+repo+"/"+tag]
	if ok {
		return checkFunction, nil
	}

	file, err := GetReleaseAsset(organization, repo, tag)
	if err != nil {
		return nil, err
	}

	plugin, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}

	runSymbol, err := plugin.Lookup("Run")
	if err != nil {
		return nil, err
	}

	runFunc, ok := runSymbol.(func(ctx context.Context, target string, command string, expectedOutput string, username string, password string, options map[string]interface{}) (bool, string))
	if !ok {
		return nil, fmt.Errorf("failed to cast Run to func")
	}

	checkFunctionCache[organization+"/"+repo+"/"+tag] = runFunc

	return runFunc, nil
}

func GetReleaseAsset(organization string, repo string, tag string) (path string, err error) {
	path, _, err = GetReleaseAssetWithTag(organization, repo, tag)
	return path, err
}

func GetReleaseAssetWithTag(organization string, repo string, tag string) (path string, tagParsed string, err error) {
	releaseExists, ok := releaseAssetCache[organization+"/"+repo+"/"+tag]
	if ok {
		return releaseExists.Path, releaseExists.Tag, nil
	}

	if tag == "latest" || tag == "" {
		tag, err = GetLatestRelease(organization, repo)
		if err != nil {
			return "", "", err
		}

	}

	path = GeneratePath(organization, repo, tag)

	exists, err := helpers.FileExists(path)
	if err != nil {
		return "", tag, err
	}

	if exists {
		releaseAssetCache[organization+"/"+repo+"/"+tag] = releaseAssetCacheStruct{
			Path: path,
			Tag:  tag,
		}

		return path, tag, nil
	}

	path, err = DownloadReleaseAsset(organization, repo, tag)
	if err != nil {
		return "", tag, err
	}

	releaseAssetCache[organization+"/"+repo+"/"+tag] = releaseAssetCacheStruct{
		Path: path,
		Tag:  tag,
	}

	return path, tag, nil
}

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

func DownloadReleaseAsset(organization string, repo string, tag string) (filepath string, err error) {
	githubClient := github.NewClient(nil)

	release, response, err := githubClient.Repositories.GetReleaseByTag(context.Background(), organization, repo, tag)
	if response.StatusCode != 200 {
		return "", fmt.Errorf("release endpoint returned %d", response.StatusCode)
	}

	if err != nil {
		return "", err
	}

	var asset *github.ReleaseAsset

	for _, a := range release.Assets {
		if strings.HasSuffix(a.GetName(), ".so") {
			asset = &a
			break
		}
	}

	if asset == nil {
		return "", fmt.Errorf("failed to find asset")
	}

	resp, err := http.Get(asset.GetBrowserDownloadURL())

	filepath = GeneratePath(organization, repo, tag)

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

func GeneratePath(organization string, repo string, tag string) string {
	return "plugins/" + organization + "-" + repo + "-" + tag + ".so"
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
