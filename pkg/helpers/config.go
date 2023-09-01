package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func FileExists(filepath string) (bool, error) {
	stat, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if stat.IsDir() {
		return false, fmt.Errorf("file is a directory: %s", filepath)
	}

	return true, nil
}

func GetEditor() string {
	editors := []string{"nano", "vim", "emacs", "vi"}
	for _, editor := range editors {
		path, err := exec.LookPath(editor)
		if err == nil && path != "" {
			return editor
		}
	}
	return ""
}

func EditFile(filepath string, editor string) error {
	editorCmd := exec.Command(editor, filepath)

	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	return editorCmd.Run()
}
