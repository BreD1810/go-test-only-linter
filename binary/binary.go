package binary

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const binaryName = "unoptimised-binary"

func buildUnoptimised(filePath string) {
	cmd := exec.Command("go", "build", "-o", binaryName, "-gcflags", "-N -l", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to build unoptimised binary: %s\n%s\n", err, output)
	}
}

func removeUnoptimised() {
	cmd := exec.Command("rm", binaryName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to delete generated binary: %s\n%s\n", err, output)
	}
}

func GetBinaryFunctions(filePath string) map[string]struct{} {
	buildUnoptimised(filePath)
	cmd := exec.Command("go", "tool", "nm", binaryName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to inspect binary: %s\n%s", err, output)
	}

	resString := string(output)
	lines := strings.Split(strings.TrimSpace(resString), "\n")

	funcNames := make(map[string]struct{})

	for n, line := range lines {
		splits := strings.Split(strings.TrimSpace(line), " ")
		if len(splits) < 3 {
			fmt.Fprintf(os.Stderr, "WRONG OUTPUT ON LINE %d: %s\n", n, splits)
			continue
		}

		funcName := strings.Split(strings.TrimSpace(line), " ")[2]
		if strings.HasPrefix(funcName, "fmt") ||
			strings.HasPrefix(funcName, "reflect") ||
			strings.HasPrefix(funcName, "go") ||
			strings.HasPrefix(funcName, "_cgo") ||
			strings.HasPrefix(funcName, "runtime") ||
			strings.HasPrefix(funcName, "internal") ||
			strings.HasPrefix(funcName, "unicode") ||
			strings.HasPrefix(funcName, "strconv") ||
			strings.HasPrefix(funcName, "sync") ||
			strings.HasPrefix(funcName, "sort") ||
			strings.HasPrefix(funcName, "time") ||
			strings.HasPrefix(funcName, "syscall") ||
			strings.HasPrefix(funcName, "math") ||
			strings.HasPrefix(funcName, "type") ||
			strings.HasPrefix(funcName, "io") ||
			strings.HasPrefix(funcName, "os") {
			continue
		}

		funcNames[funcName] = struct{}{}
	}

	removeUnoptimised()

	return funcNames
}
