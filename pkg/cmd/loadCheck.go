package cmd

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var testCmd = &cobra.Command{
	Use: "test",
	Run: testRun,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func testRun(cmd *cobra.Command, args []string) {
	config.Init()

	out, err := yaml.Marshal(config.RunningConfig)
	if err != nil {
		panic(err)
	}

	println(string(out))
}
