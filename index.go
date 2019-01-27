package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const XkcdUrl = "https://xkcd.com/"

type Comic struct {
	Month      string
	Number     int `json:"num"`
	Year       string
	Transcript string
	Img        string
	Title      string
	SafeTitle  string `json:"safe_title"`
}

func ComicCount() int {
	dir, _ := os.Getwd()
	path := dir + "/items/"

	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			count++
		}
	}

	return count
}

func GetComic(number int) (*Comic, error) {
	comicUrl := "https://xkcd.com/" + strconv.Itoa(number) + "/info.0.json"
	resp, err := http.Get(comicUrl)
	fmt.Printf("getting comic: %d\n", number)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("comic %d not found\n", number)
		}
		return nil, errors.New("not_found")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("error reading response body: %s", err)
	}
	resp.Body.Close()

	var result Comic
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("error unmarshalling response body: %s", err)
		return nil, err
	}
	return &result, nil

}

func UpdateIndex(c *Comic) error {
	dir, _ := os.Getwd()
	path := dir + "/items/"

	f, err := os.OpenFile(path+"index.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("\nNumber: " + strconv.Itoa(c.Number) + "\nImg url: " + c.Img + "\nTitle: " + c.Title + "\nTranscript:\n\n" + c.Transcript + "\n\n***")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (c *Comic) Save() error {
	fmt.Printf("saving comic %s\n", strconv.Itoa(c.Number))

	dir, _ := os.Getwd()
	path := dir + "/items/"

	filename := strconv.Itoa(c.Number) + ".csv"
	content := []byte(strings.Join([]string{strconv.Itoa(c.Number), c.Year,
		c.Title, c.SafeTitle, c.Year, c.Transcript, c.Img}, ","))
	return ioutil.WriteFile(path+filename, content, 0600)
}

func SearchIndex(term string) ([]string, error) {
	dir, _ := os.Getwd()
	path := dir + "/items/"

	b, err := ioutil.ReadFile(path + "index.txt")

	var result = make([]string, 0)

	if err != nil {
		fmt.Print(err)
	}

	content := string(b)

	arr := strings.Split(content, "***")

	for _, v := range arr {
		if strings.Contains(v, term) {
			result = append(result, v)
		}
	}

	return result, nil
}
