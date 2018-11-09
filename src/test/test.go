package main

import (
	"fmt"
	"qubic"
)

func main() {
	q := qubic.NewQuery()
	i := []string{"user", "project"}
	j := []string{"id", "ttl"}
	s := ""
	w := []string{"4", "3", "11"}
	q.Limit("100", 15).Orderby(i, "asc").Orderby(i, "desc").Groupby("user.id").Groupby(i).Join("user", "user.id = training.id", "left").Where("user.id =", 9223372036854775807).Where_in("user.id", w).From("training").From(i).Select(j).Select("competency").Extract(&s)
	fmt.Println(s)
}
