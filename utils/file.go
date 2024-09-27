package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(pwd + path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := ""
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		content += (sc.Text() + "\n")
	}

	return strings.TrimRight(content, "\n")
}
