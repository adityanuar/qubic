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
	q.Like("person.name", "Ad", "after").Limit("100", 15).Orderby(i, "asc").Orderby(i, "desc").Groupby("user.id").Groupby(i).Join("user", "user.id = training.id", "left").Where("user.role =", "project manager", false).Where_in("user.id", w).From("training").From(i).Select(j).Select("competency").Where_raw("(hr_contract.end IS NULL OR hr_contract.end >= 54)").Extract(&s)
	fmt.Println(s)
}
