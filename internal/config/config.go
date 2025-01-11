package config

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed assets/config.ini
var configInit []byte

func Write(file string) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	configStr := string(configInit)

	lines := strings.Split(configStr, "\n")
	for _, line := range lines {
		_, _ = fmt.Fprintln(w, line)
	}
	return w.Flush()
}
