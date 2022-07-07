package main

import (
	"ebook-downloader/utils"
	"fmt"
	"os"
)

func main() {

	var (
		//oreillyExecutable = os.Getenv("OREILLY_DOWNLOADER_PATH")
		booksRootPath = os.Getenv("BOOKS_ROOT_PATH")
	)

	err := utils.FormatEbooks(booksRootPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//
	//executableArgs := []string{
	//	oreillyExecutable,
	//	"--kindle",
	//	"--root-path", booksRootPath,
	//}
	//
	//downloader := utils.SafariDownloader{
	//	ExecutablePath: os.Getenv("PYTHON_EXECUTABLE"),
	//	ExecutableArgs: executableArgs,
	//}
	//
	//books := utils.ReadBookList("bookList.txt")
	//
	//var wg sync.WaitGroup
	//ch := make(chan string, len(books))
	//
	//for i := 0; i < 20; i++ {
	//	go func(workerId int) {
	//		fmt.Println("Worker", workerId, "working...")
	//		for {
	//			isbn, open := <-ch
	//
	//			if open == false {
	//				break
	//			}
	//
	//			fmt.Printf("Downloading ebook %v...\n", isbn)
	//			downloader.DownloadEbook(isbn)
	//			fmt.Printf("Finished download of ebook %v.\n", isbn)
	//			wg.Done()
	//		}
	//	}(i)
	//}
	//
	//for _, book := range books {
	//	wg.Add(1)
	//	ch <- book
	//}
	//
	//wg.Wait()
	//close(ch)
	//
	//fmt.Println("Exit, but why?")

	utils.ExtractEbooks(booksRootPath, booksRootPath+"/Extracted")
}

//func downloadFromOreilly() {
//	var (
//		oreillyExecutable = os.Getenv("OREILLY_DOWNLOADER_PATH")
//		booksRootPath     = os.Getenv("BOOKS_ROOT_PATH")
//	)
//
//	executableArgs := []string{
//		oreillyExecutable,
//		"--kindle",
//		"--root-path", booksRootPath,
//	}
//
//	downloader := utils.SafariDownloader{
//		ExecutablePath: os.Getenv("PYTHON_EXECUTABLE"),
//		ExecutableArgs: executableArgs,
//	}
//
//	books := utils.ReadBookList("bookList.txt")
//
//	var wg sync.WaitGroup
//	ch := make(chan string, len(books))
//
//	for i := 0; i < 20; i++ {
//		worker(i, ch)
//	}
//
//	for _, book := range books {
//		wg.Add(1)
//		ch <- book
//	}
//
//	wg.Wait()
//	close(ch)
//
//	fmt.Println("Exit, but why?")
//}

//func worker(workerId int, ch chan string) {
//	go func() {
//		fmt.Println("Worker", workerId, "working...")
//		for {
//			isbn, open := <-ch
//
//			if open == false {
//				break
//			}
//
//			fmt.Printf("Downloading ebook %v...\n", isbn)
//			downloader.DownloadEbook(isbn)
//			fmt.Printf("Finished download of ebook %v.\n", isbn)
//			wg.Done()
//		}
//	}(i)
//}
