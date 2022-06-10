package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ExtractEbooks(currentPath string, targetPath string) {
	booksDirs, err := os.ReadDir(currentPath)

	if err != nil {
		fmt.Println(err)
	}

	justNameRg, _ := regexp.Compile("\\([0-9]{3,}\\)")

	for _, bookDir := range booksDirs {
		isbn := string(justNameRg.Find([]byte(bookDir.Name())))
		isbnIdxRange := justNameRg.FindIndex([]byte(bookDir.Name()))
		if len(isbnIdxRange) <= 0 {
			continue
		}

		isbn = strings.ReplaceAll(strings.ReplaceAll(isbn, "(", ""), ")", "")

		bookName := bookDir.Name()[:isbnIdxRange[0]-1]
		bookName = strings.ReplaceAll(bookName, " ", "_")
		bookName = strings.ReplaceAll(bookName, ".", "")

		currentEbookPath := currentPath + "/" + bookDir.Name() + "/" + isbn + ".epub"
		newEbookPath := targetPath + "/" + bookName + ".epub"

		err := os.Rename(currentEbookPath, newEbookPath)

		if err != nil {
			fmt.Println("The path is invalid:", currentEbookPath)
			fmt.Println(err)
		}
	}
}
