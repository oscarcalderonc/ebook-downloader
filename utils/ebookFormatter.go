package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func FormatEbooks(ebookDirectory string) error {

	books, err := ioutil.ReadDir(ebookDirectory)

	if err != nil {
		return err
	}

	//ebook-convert Books/ distributed-computing-with-go.azw3

	convertArguments := []string{
		"--keep-ligatures", "--embed-all-fonts", "--output-profile=kindle_pw", "--margin-bottom=0", "--margin-left=0", "--margin-right=0", "--margin-top=0", "--filter-css=color,margin-left,margin-right",
	}

	wg := sync.WaitGroup{}
	maxGoroutines := 20
	var mxCount sync.Mutex
	currGoroutines := 0

	var mxFailed sync.Mutex
	failedBooks := []string{}

	for _, path := range books {

		for currGoroutines >= maxGoroutines {
		}

		//fmt.Println("Increasing goroutines to", currGoroutines+1)
		wg.Add(1)
		currGoroutines++
		fmt.Println(path.Name())
		bookName := strings.Split(path.Name(), ".")[0]

		go func(bookName string) {
			defer func() {
				mxCount.Lock()
				defer mxCount.Unlock()
				//fmt.Println("Reducing goroutines to", currGoroutines-1)
				wg.Done()
				currGoroutines--
			}()
			fullBookPath := ebookDirectory + "/" + bookName
			args := []string{fullBookPath + ".epub", fullBookPath + ".azw3"}
			args = append(args, convertArguments...)

			cmd := exec.Command("ebook-convert", args...)
			err := cmd.Run()
			if err != nil {
				fmt.Println(fmt.Errorf("%v: %v", bookName, err))
				mxFailed.Lock()
				defer mxFailed.Unlock()
				failedBooks = append(failedBooks, bookName)
			}
		}(bookName)
	}

	wg.Wait()
	fmt.Println("Finish book conversion")
	fmt.Println("Failures:", failedBooks)

	books, err = ioutil.ReadDir(ebookDirectory)

	for _, book := range books {
		if strings.Contains(book.Name(), ".azw3") == true {
			err = os.Rename(ebookDirectory+"/"+book.Name(), ebookDirectory+"/kindle/"+book.Name())

			if err != nil {
				fmt.Println(fmt.Errorf("%v: %v", "Issue trying to move the book", book.Name()))
			}
		}
	}

	return nil
}
