package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ReadBookList(path string) []string {
	booksFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer booksFile.Close()

	bookFolders, err := os.ReadDir(os.Getenv("BOOKS_ROOT_PATH"))

	if err != nil {
		fmt.Println(err)
	}

	bookIdRegexp, _ := regexp.Compile("[0-9]{7,}")
	booksUnique := make(map[string]int)

	for _, bookFolder := range bookFolders {
		//fmt.Println(bookFolder.Name())
		match := string(bookIdRegexp.Find([]byte(bookFolder.Name())))
		if _, value := booksUnique[match]; !value {
			booksUnique[match] = 1
		} else {
			fmt.Println("This is repeated ", match)
		}
	}

	reader := bufio.NewReader(booksFile)

	bookIds := []string{}

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		match := string(bookIdRegexp.Find(line))
		if _, value := booksUnique[match]; !value {
			booksUnique[match] = 1
			bookIds = append(bookIds, match)
		} else {
			fmt.Println("This is repeated ", match)
		}
	}

	fmt.Println(bookIds)
	return bookIds
}
