// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kidlj/demo/ent/o2m2types/ent"
	"github.com/kidlj/demo/ent/o2m2types/ent/pet"
	"github.com/kidlj/demo/ent/o2m2types/ent/predicate"
	"github.com/kidlj/demo/ent/o2m2types/ent/user"

	_ "github.com/lib/pq"
)

func Example_O2M2Types() {
	client, err := ent.Open("postgres", "postgres://demo:demo@localhost/ent?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// User created: User(id=1, age=30, name=a8m)
	// a8m
	// 2
}

func Do(ctx context.Context, client *ent.Client) error {
	// Create the 2 pets.
	pedro, err := client.Pet.
		Create().
		SetName("pedro").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %v", err)
	}
	lola, err := client.Pet.
		Create().
		SetName("lola").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %v", err)
	}
	// Create the user, and add its pets on the creation.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddPets(pedro, lola).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}
	fmt.Println("User created:", a8m)
	// Output: User(id=1, age=30, name=a8m)

	// Query the owner. Unlike `Only`, `OnlyX` panics if an error occurs.
	owner := pedro.QueryOwner().OnlyX(ctx)
	fmt.Println(owner.Name)
	// Output: a8m

	// Traverse the sub-graph. Unlike `Count`, `CountX` panics if an error occurs.
	count := pedro.
		QueryOwner(). // a8m
		QueryPets().  // pedro, lola
		CountX(ctx)   // count
	fmt.Println(count)
	// Output: 2

	pets, err := client.Pet.
		Query().
		Where(pet.HasOwner()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets:", pets)

	pets2, err := client.Pet.
		Query().
		Where(pet.HasOwnerWith(user.Name("a8m"))).
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets2:", pets2)

	pets3, err := client.Pet.
		Query().
		Where(pet.Not(pet.NameHasPrefix("Ari"))).
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets3:", pets3)

	pets4, err := client.Pet.
		Query().
		Where(
			pet.Or(
				pet.HasOwner(),
				pet.Not(pet.NameHasPrefix("lo")),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets4:", pets4)

	pets5, err := client.Pet.
		Query().
		Where(
			pet.And(
				pet.HasOwner(),
				pet.Not(pet.NameHasPrefix("lo")),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets5:", pets5)

	pets6 := client.Pet.
		Query().
		Where(predicate.Pet(func(s *sql.Selector) {
			s.Where(sql.InInts(pet.OwnerColumn, 1, 2, 3))
		})).
		AllX(ctx)
	if err != nil {
		return fmt.Errorf("querying pets: %v", err)
	}
	fmt.Println("pets6:", pets6)

	return nil
}
