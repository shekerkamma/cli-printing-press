package generator

import (
	"fmt"
	"os"
	"os/exec"
)

// EnsureGoModuleReady hydrates go.sum so freshly generated CLIs build in a
// cold module cache without a manual `go mod tidy`.
func EnsureGoModuleReady(outputDir string) error {
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = outputDir
	tidyCmd.Stdout = os.Stderr
	tidyCmd.Stderr = os.Stderr
	if err := tidyCmd.Run(); err != nil {
		return fmt.Errorf("running go mod tidy: %w", err)
	}
	return nil
}
