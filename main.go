package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type Book struct {
    title string `json:"title"`
    authors []string
}
type Books struct {
    books []Book
}

func getListOfBooks(library_id string) Books {
    resp,err := http.Get("http://localhost:8000/calibre/ajax/books/"+library_id)
    if err != nil {
        fmt.Println("error ocurred")
        log.Fatalln(err)
    }
    
    defer resp.Body.Close()

    body,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(body))
    fmt.Println("function which returns list of books")
    
    return Books{books:[]Book{ Book{title: "title", authors:[]string{"author"}} }}
}

func main() {
    fmt.Println(getListOfBooks("Calibre_Library"))
    fmt.Println("main func")
}
