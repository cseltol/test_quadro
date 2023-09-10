// package main implementsa client for BookShelf service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	pb "github.com/cseltol/test_quadro/testp"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

// printBook get the book for the given book query
func printBook(client pb.BookShelfClient, author *pb.Author) string {
	log.Printf("Getting book for Author %s\n", author.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	book, err := client.GetBook(ctx, author)
	if err != nil {
		log.Fatalf("client.GetBook failed: %v\n", err)
	}
	log.Println(book.GetName())
	return book.GetName()
}

func printAuthor(client pb.BookShelfClient, book *pb.Book) string {
	log.Printf("Getting author for book %s\n", book.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	author, err := client.GetAuthor(ctx, book)
	if err != nil {
		log.Fatalf("client.GetAuthor failed: %v\n", err)
	}
	log.Println(author.GetName())
	return author.GetName()
}

func main() {
	flag.Parse()

	// Set up connection to the server.
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewBookShelfClient(conn)

	authorName := "Булгаков М.А."
	printBook(client, &pb.Author{Name: &authorName})

	bookName := "Мастер и Маргарита"
	printAuthor(client, &pb.Book{Name: &bookName, Author: &pb.Author{}})
}
