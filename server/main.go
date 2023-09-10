// package main implementsa server for BookShelf service.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/cseltol/test_quadro/repo"
	pb "github.com/cseltol/test_quadro/testp"
)

var (
	port       = flag.Int("port", 50051, "The server port")
)

type bookShelfServer struct {
	pb.UnimplementedBookShelfServer
}

func (s *bookShelfServer) GetAuthor(ctx context.Context, book *pb.Book) (*pb.Author, error) {
	conn := repo.GetConnection()
	defer conn.Close()

	var author *pb.Author
	err := conn.QueryRow(`
	SELECT author FROM authors
	WHERE author.name = (
		SELECT book FROM books
		WHERE book.name = $1
	);`, 
	book.GetName(),
	).Scan(&author)

	if err != nil {
		return nil, errors.New("could not find the author of specified book")
	}
	return author, nil
}

func (s *bookShelfServer) GetBook(ctx context.Context, author *pb.Author) (*pb.Book, error) {
	conn := repo.GetConnection()
	defer conn.Close()

	var book *pb.Book
	err := conn.QueryRow(`
	SELECT book FROM books
	WHERE book.author = (
		SELECT author FROM authors
		WHERE author.name = $1
	);`,
	author.GetName(),
	).Scan(&book)

	if err != nil {
		return nil, errors.New("could not find the book by specified author")
	}

	return book, nil
}

func main() {
	repo.InitDB()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterBookShelfServer(grpcServer, &bookShelfServer{})
	grpcServer.Serve(lis)
}
