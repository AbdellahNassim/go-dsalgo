package main

import (
	"fmt"
	"go-dsalgo/pkg/set"
)

func RunHashSetImplementation() {

	hashSet := set.NewHashSet()

	hashSet.Add(1, 2, 3, 4, 5)

	hashSet2 := set.NewHashSet([]interface{}{6,7,8,9,10}...)

	fmt.Print(hashSet2.String())
	fmt.Print(hashSet.Values())

	comb := hashSet.Union(hashSet2)
	fmt.Print(comb.String())
}