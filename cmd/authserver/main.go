package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	api "serviceauth/internal/api/user"
	repos "serviceauth/internal/repository/user"
	serv "serviceauth/internal/service/user"
	desc "serviceauth/pkg/auth_v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DBConfig struct {
	user     string
	password string
	port     string
	dbname   string
}

func (db DBConfig) String() string {
	return fmt.Sprintf("postgres://%v:%v@localhost:%v/%v", db.user, db.password, db.port, db.dbname)
}

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbconfig := DBConfig{
		user:     os.Getenv("PG_USER"),
		password: os.Getenv("PG_PASSWORD"),
		port:     os.Getenv("PG_PORT"),
		dbname:   os.Getenv("PG_DATABASE_NAME"),
	}
	lis, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.New(ctx, dbconfig.String())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repository := repos.NewRepository(pool)
	service := serv.NewService(repository)
	server := api.NewImplementation(service)
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthenticationServer(s, server)

	log.Printf("server listening %v", lis.Addr().String())

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
