package main

import (
	"ebook-downloader/utils"
	"os"
)

func main() {

	var (
		oreillyExecutable = os.Getenv("OREILLY_DOWNLOADER_PATH")
		booksRootPath    = os.Getenv("BOOKS_ROOT_PATH")
	)

	executableArgs := []string{
		oreillyExecutable,
		"--kindle",
		"--root-path", booksRootPath,
	}

	downloader := utils.SafariDownloader{
		ExecutablePath: os.Getenv("PYTHON_EXECUTABLE"),
		ExecutableArgs: executableArgs,
	}

	downloader.DownloadEbook("9781484226919")
}
