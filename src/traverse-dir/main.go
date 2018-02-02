package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 40)

func walkDir(dir string, fileInfoCh chan<- fileInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "walkDir: %v\n", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, fileInfoCh, wg)
		} else {
			fileInfoCh <- fileInfo{entry.Size(), entry.Name()}
		}
	}
}

type fileInfo struct {
	size int64
	name string
}

func setTopTen(list []fileInfo, newItem fileInfo) []fileInfo {
	for i, d := range list {
		if d.size < newItem.size {
			copy(list[i+1:], list[i:])
			list[i] = newItem
			return list
		}
	}
	return list
}

func main() {
	// Determine the initial directories to traverse
	flag.Parse()
	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	// traverse the files
	fileInfoCh := make(chan fileInfo)
	var wg sync.WaitGroup

	for _, dir := range dirs {
		wg.Add(1)
		go walkDir(dir, fileInfoCh, &wg)
	}

	go func() {
		wg.Wait()
		close(fileInfoCh)
	}()

	// set the tick for report
	tick := time.Tick(500 * time.Millisecond)

	// genearte the report
	var totalFiles, totalSize int64
	topTenList := make([]fileInfo, 10)
loop:
	for {
		select {
		case info, ok := <-fileInfoCh:
			if !ok {
				break loop
			}
			totalFiles++
			totalSize += info.size
			topTenList = setTopTen(topTenList, info)
		case <-tick:
			fmt.Printf("%d files %.1f MB scanned\n", totalFiles, float64(totalSize)/1e6)
		}
	}

	fmt.Printf("%d files %.1f MB scanned\n", totalFiles, float64(totalSize)/1e6)
	fmt.Printf("Largest %d files: \n", int(math.Min(float64(len(topTenList)), float64(totalSize))))
	for _, info := range topTenList {
		fmt.Printf("%s %.1f MB \n", info.name, float64(info.size)/1e6)
	}
}
