package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var mx = make(map[string]int)

var ch = make(chan map[string]int, 1)
var wg = &sync.WaitGroup{}

func main() {
	go func() { ch <- mx }()
	wg.Add(1)
	go dir("C:\\Users\\Rupak_Veerla\\Downloads", ch, wg)
	wg.Wait()
	close(ch)
	fileCount(ch)
}

func dir(path string, ch chan map[string]int, wg *sync.WaitGroup) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go dir(path+"\\"+file.Name(), ch, wg)
			continue
		}
		m := <-ch
		s := strings.Split(file.Name(), ".")
		m[s[len(s)-1]]++
		ch <- m
	}
	wg.Done()
}

func fileCount(ch chan map[string]int) {
	m := <-ch
	for k, v := range m {
		fmt.Printf("There are %d %q files\n", v, k)
	}
}
