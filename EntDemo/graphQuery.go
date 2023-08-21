package main

import (
	"context"
	"fmt"
	"learning/ent"
	"learning/ent/car"
	"learning/ent/group"
	"learning/ent/user"
	"time"
)

func CreateGraph(ctx context.Context, client *ent.Client) error {
	DukeKhieu, err := client.User.Create().SetName("DukeKhieu").SetAge(20).Save(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new user: %w", err)
	}
	KhieuDuc, err := client.User.Create().SetName("KhieuDuc").SetAge(21).Save(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new user: %w", err)
	}

	err = client.Car.Create().SetModel("BWM").SetRegisteredAt(time.Now()).SetOwner(DukeKhieu).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new car: %w", err)
	}

	err = client.Car.Create().SetModel("Audi").SetRegisteredAt(time.Now()).SetOwner(DukeKhieu).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new car: %w", err)
	}

	err = client.Car.Create().SetModel("MayBach").SetRegisteredAt(time.Now()).SetOwner(KhieuDuc).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new car: %w", err)
	}

	err = client.Group.Create().SetName("GitLab").AddUsers(KhieuDuc, DukeKhieu).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new group: %w", err)
	}

	err = client.Group.Create().SetName("GitHub").AddUsers(KhieuDuc).Exec(ctx)
	if err != nil {
		return fmt.Errorf("Error creating new group: %w", err)
	}
	fmt.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.Query().Where(group.Name("GitHub")).QueryUsers().QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("Error Querrying GitHub Cars: %w", err)
	}
	fmt.Println("GitHub Group cars returned:", cars)
	fmt.Println()
	return nil
}

func QueryDukeKhieuCars(ctx context.Context, client *ent.Client) error {
	duke := client.User.Query().Where(
		user.HasCars(),
		user.Name("DukeKhieu"),
	).OnlyX(ctx)
	//like only but panic when a error happened
	cars := duke.QueryGroups().QueryUsers().QueryCars().Where(
		car.Not(
			car.Model("BWM"),
		),
	).AllX(ctx)
	fmt.Println("cars returned: ", cars)
	fmt.Println()
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups := client.Group.Query().Where(group.HasUsers()).AllX(ctx)
	fmt.Println("groups returned:", groups)
	fmt.Println()
	return nil
}
