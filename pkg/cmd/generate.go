package cmd

import (
	"fmt"
	"io"
	"os"
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

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateRun(cmd *cobra.Command, args []string) {
	config.Init()

	var releases = make(map[structs.Release_s]bool)

	for _, team := range config.RunningConfig.Teams {
		for _, check := range team.Checks {
			releases[check.Release] = true
		}
	}

	for release := range releases {
		fmt.Printf("%s/%s@%s\n", release.Org, release.Repo, release.Tag)
	}

	err := deleteImports("pkg/checks/imports")
	if err != nil {
		logrus.Fatalf("Failed to delete imports: %s", err.Error())
	}

	err = writeMain()
	if err != nil {
		logrus.Fatalf("Failed to write main: %s", err.Error())
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

	tmplString, err := io.ReadAll(tmplFile)
	if err != nil {
		return err
	}

	fmt.Println(string(tmplString))

	tmpl, err := template.New("pkg/checks/imports/main.go").Parse(string(tmplString))
	if err != nil {
		return err
	}

	return tmpl.Execute(outputFile, nil)
}
