package main

import (
	"context"
	"fmt"
	"learning/ent"
	"log"
	"time"
)

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created:", tesla)

	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created:", ford)

	u, err := client.User.Create().SetAge(30).SetName("DucKM").AddCars(ford, tesla).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func GetUserCars(ctx context.Context, user *ent.User) ([]*ent.Car, error) {
	cars, err := user.QueryCars().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error querying user's cars: %w", err)
	}
	fmt.Println(user.Name + "'s Cars List: ")
	for _, v := range cars {
		fmt.Println(v)
	}
	return cars, nil
}

func GetAllCars(ctx context.Context, client *ent.Client) ([]*ent.Car, error) {
	cars, err := client.Car.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error querying cars: %w", err)
	}
	fmt.Println("Car list: ")
	for _, c := range cars {
		fmt.Println(c)
	}
	return cars, nil
}
