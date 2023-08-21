package main

import (
	"context"
	"learning/ent"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// go generate ./ent
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=entDemo password=Password123#@! sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// CreateUser(ctx, client)
	// QueryUserWithName(ctx, client, "Gruxeon")
	// QueryAllUser(ctx, client)
	// CreateCars(ctx, client)

	// u, _ := QueryUserWithName(ctx, client, "DucKM")
	// GetUserCars(ctx, u)

	// cars, _ := GetAllCars(ctx, client)
	// for _, c := range cars {
	// 	GetCarOwner(ctx, c)
	// }
	// CreateGraph(ctx, client)
	QueryGithub(ctx, client)
	QueryDukeKhieuCars(ctx, client)
	QueryGroupWithUsers(ctx, client)
}
