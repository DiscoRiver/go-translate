package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

var wg sync.WaitGroup
var channel chan string

func main() {
	path := os.Args[1]
	lan := os.Args[2]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		go func() {
			trans(scanner.Text(), lan)
		}()
		time.Sleep(40 * time.Millisecond)
	}
	wg.Wait()
}

func trans(text string, tLanguage string) {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Sets the target language
	target, err := language.Parse(tLanguage)
	if err != nil {
		log.Fatalln(err)
	}

	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(translations[0].Text)
	wg.Done()
}
