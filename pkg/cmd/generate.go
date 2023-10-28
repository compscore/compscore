package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for current checks",
	Long:  "Generate code for current checks",
	Run:   generateRun,
}

func generateRun(cmd *cobra.Command, args []string) {
	config.Init()

	var releases = make(map[structs.Release_s]bool)

	for _, check := range config.Checks {
		if check.Release.Tag == "" {
			check.Release.Tag = "latest"
		}
		releases[check.Release] = true
	}

	err := deleteImports("pkg/checks/imports")
	if err != nil {
		logrus.Fatalf("Failed to delete imports: %s", err.Error())
	}

	err = writeMain()
	if err != nil {
		logrus.Fatalf("Failed to write main: %s", err.Error())
	}

	err = writeChecks(releases)
	if err != nil {
		logrus.Fatalf("Failed to write checks: %s", err.Error())
	}

	fmtCmd := exec.Command("go", "fmt", "./...")
	fmtCmd.Stdout = os.Stdout
	fmtCmd.Stderr = os.Stderr
	err = fmtCmd.Run()
	if err != nil {
		logrus.Fatalf("Failed to run go fmt: %s", err.Error())
	}

	for release := range releases {
		getCmd := exec.Command("go", "get", fmt.Sprintf("github.com/%s/%s@%s", release.Org, release.Repo, release.Tag))
		getCmd.Stdout = os.Stdout
		getCmd.Stderr = os.Stderr
		fmt.Println("go", "get", fmt.Sprintf("github.com/%s/%s@%s", release.Org, release.Repo, release.Tag))
		err = getCmd.Run()
		if err != nil {
			logrus.Fatalf("Failed to run go get: %s", err.Error())
		}
	}

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr
	err = tidyCmd.Run()
	if err != nil {
		logrus.Fatalf("Failed to run go mod tidy: %s", err.Error())
	}
}

func deleteImports(path string) error {
	path = strings.TrimSuffix(path, "/")

	dir, err := os.Open(path)
	if err != nil {
		return err
	}

	files, err := dir.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(path + "/" + file)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeMain() error {
	outputFile, err := os.Create("pkg/checks/imports/main.go")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	tmplFile, err := os.Open("pkg/checks/template/main.go.tmpl")
	if err != nil {
		return err
	}
	defer tmplFile.Close()

	tmplString, err := io.ReadAll(tmplFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New("pkg/checks/imports/main.go").Parse(string(tmplString))
	if err != nil {
		return err
	}

	return tmpl.Execute(outputFile, nil)
}

func writeChecks(releases map[structs.Release_s]bool) error {
	tmplFile, err := os.Open("pkg/checks/template/check.go.tmpl")
	if err != nil {
		return err
	}
	defer tmplFile.Close()

	tmplString, err := io.ReadAll(tmplFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New("pkg/checks/imports/main.go").Parse(string(tmplString))
	if err != nil {
		return err
	}

	for release := range releases {
		outputFile, err := os.Create(fmt.Sprintf("pkg/checks/imports/%s-%s.go", release.Org, release.Repo))
		if err != nil {
			return err
		}
		defer outputFile.Close()

		err = tmpl.Execute(outputFile, release)
		if err != nil {
			return err
		}
	}

	return nil
}
