package utils

import (
	"fmt"
	"os/exec"
)

type EbookDownloader interface {
	DownloadEbook(string) string
}

type SafariDownloader struct {
	ExecutablePath string
	ExecutableArgs []string
}

func (s *SafariDownloader) DownloadEbook(isbn string) string {

	completeArguments := append(s.ExecutableArgs, isbn) //"--book-path", isbn,
	cmd := exec.Command(s.ExecutablePath, completeArguments...)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
	return ""
}
