package main

type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublishableDate string `json:"publish_date"`
	ISBN            string `json:"isbn"`
}
