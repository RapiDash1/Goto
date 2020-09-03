package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os/exec"
	"sync"
)

func gotoWebsite(url *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	exec.Command("rundll32", "url.dll,FileProtocolHandler", <-*url).Start()
}

func getCommandLineInput() string {
	webKey := flag.String("key", "youtube", "Key saved in goto.json")
	flag.Parse()
	return *webKey
}

func getUrl(url *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	plan, _ := ioutil.ReadFile("goto.json")
	var data map[string]string
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}
	*url <- data[getCommandLineInput()]
	close(*url)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	url := make(chan string, 1)
	go getUrl(&url, &wg)
	go gotoWebsite(&url, &wg)
	wg.Wait()
}
