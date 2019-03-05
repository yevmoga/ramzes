package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Category struct {
	Id         string     `json:"cid"`
	Name       string     `json:"title"`
	UrlName    string     `json:"url_name"`
	Image      string     `json:"image"`
	Categories Categories `json:"categories"`
}

//описанием, количеством, ценой, картинкой
//В идеале ещё
//Ценой привязанной к доллару
type Product struct {
	Id          string `json:"cid"`
	Name        string `json:"title"`
	Decsription string `json:"url_name"`
	Image       string `json:"image"`
	PriceUah    string `json:"image"`
	PriceUsd    string `json:"image"`
}

var BaseUrl = "https://www.rcscomponents.kiev.ua/"

func main() {

	url := BaseUrl + "content.php?lang=russian"

	response, err := http.Get(url)
	if err != nil {
		log.Printf("error getting url: %s", url)
		return
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("failed to fetch %s: expected 200, got %d", url, response.StatusCode)
		return
	}

	var categories Categories
	if err := json.NewDecoder(response.Body).Decode(&categories); err != nil {
		log.Printf("failed to parse. Error: %s", err)
		return
	}

	products := []Product{}
	task := make(chan string)

	go func() {
		for {
			categories.getUrl(task)
		}
	}()

	go func() {
		for {
			result, err := fetchProduct(task)
			if err != nil {
				log.Printf("failed to parse. Error: %s", err)
			}
			products = append(products, result...)
		}
	}()

	//if err := toExcel(results); err != nil {
	//	log.Printf("could not write to file: %v", err)
	//}

	fmt.Scanln()
	fmt.Println("done")
}

func fetchProduct(url chan string) ([]Product, error) {
	// todo-moga
	fmt.Println("read: " + <-url)
	return []Product{}, nil
}

func toExcel(products chan []Product) error {
	// todo-moga
	return nil
}

type Categories []Category

func (Categories) len() int {
	// todo-moga
	return 1
}

func (categories Categories) getUrl(task chan string) {
	for _, c := range categories {
		fmt.Println("category: " + c.Id)
		if c.Categories != nil {
			//c.Categories.getUrl(task)
		} else {
			task <- fmt.Sprintf("%scatalog/%s/%s/showall", BaseUrl, c.UrlName, c.Id)
		}
	}
}
