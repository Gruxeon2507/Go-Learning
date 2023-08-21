package main

import (
	"context"
	"fmt"
	"learning/ent"
	"learning/ent/user"
	"log"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(20).SetName("Gruxeon").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUserWithName(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Name(name)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querrying user: %w", err)
	}
	log.Println("user returned", u)
	return u, nil
}
func QueryAllUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querrying user: %w", err)
	}
	for i, v := range u {
		fmt.Println("User no", i, ": ", v)
	}
	return u, nil
}

func GetCarOwner(ctx context.Context, car *ent.Car) (*ent.User, error) {
	u, err := car.QueryOwner().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get car owner failed: %w", err)
	}
	fmt.Println("Owner of", car, "is", u)
	return u, nil
}
