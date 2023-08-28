package helpers

import (
	"os"
	"os/exec"
)

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
