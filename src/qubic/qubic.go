package qubic

import (
	"fmt"
	"reflect"
	"strconv"
)

type Query struct {
	Sel      []string
	Fr       []string
	Wh       []string
	Jo       []string
	Gr       []string
	Orb_asc  []string
	Orb_desc []string
}

func NewQuery() *Query {
	return &Query{
		Sel:      []string{},
		Fr:       []string{},
		Wh:       []string{},
		Jo:       []string{},
		Gr:       []string{},
		Orb_asc:  []string{},
		Orb_desc: []string{},
	}
}

func (q *Query) Select(s interface{}) *Query {
	t := reflect.TypeOf(s)
	switch t.Kind() {
	case reflect.Slice:
		q.Sel = append(q.Sel, s.([]string)...)
	case reflect.String:
		q.Sel = append(q.Sel, s.(string))
	}
	return q
}

func (q *Query) From(f interface{}) *Query {
	t := reflect.TypeOf(f)
	switch t.Kind() {
	case reflect.Slice:
		q.Fr = append(q.Fr, f.([]string)...)
	case reflect.String:
		q.Fr = append(q.Fr, f.(string))
	}
	return q
}

func (q *Query) Where(w string, s interface{}) *Query {
	que := w
	t := reflect.TypeOf(s)
	switch t.Kind() {
	case reflect.String:
		que += " " + s.(string)
	case reflect.Uint:
		fmt.Println("unsigned")
		que += " " + strconv.FormatUint(uint64(s.(uint)), 10)
	case reflect.Uint8:
		que += " " + strconv.FormatUint(uint64(s.(uint8)), 10)
	case reflect.Uint16:
		que += " " + strconv.FormatUint(uint64(s.(uint16)), 10)
	case reflect.Uint32:
		que += " " + strconv.FormatUint(uint64(s.(uint32)), 10)
	case reflect.Uint64:
		que += " " + strconv.FormatUint(uint64(s.(uint64)), 10)
	case reflect.Int:
		que += " " + strconv.Itoa(s.(int))
	case reflect.Bool:
		if s == false {
			que += " 0"
		} else {
			que += " 1"
		}
	}
	fmt.Println(t.Kind())
	if que != w {
		q.Wh = append(q.Wh, que)
	}
	return q
}

func (q *Query) Where_in(w string, s interface{}) *Query {
	que := w
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Slice {
		que = que + " IN ("
		for k, v := range s.([]string) {
			if k < len(s.([]string))-1 {
				que = que + v + ", "
			} else {
				que = que + v + ")"
			}
		}
	}
	if que != w {
		q.Wh = append(q.Wh, que)
	}
	return q
}

func (q *Query) Join(j string, c string, t string) *Query {
	que := ""
	switch t {
	case "inner":
		que = "INNER JOIN " + j + " ON " + c
	case "right":
		que = "RIGHT JOIN " + j + " ON " + c
	case "left":
		que = "LEFT JOIN " + j + " ON " + c
	}
	if que != "" {
		q.Jo = append(q.Jo, que)
	}
	return q
}

func (q *Query) Groupby(s interface{}) *Query {
	t := reflect.TypeOf(s)
	switch t.Kind() {
	case reflect.Slice:
		q.Gr = append(q.Gr, s.([]string)...)
	case reflect.String:
		q.Gr = append(q.Gr, s.(string))
	}
	return q
}

func (q *Query) Orderby(s interface{}, o string) *Query {
	t := reflect.TypeOf(s)
	switch o {
	case "asc":
		switch t.Kind() {
		case reflect.Slice:
			q.Orb_asc = append(q.Orb_asc, s.([]string)...)
		case reflect.String:
			q.Orb_asc = append(q.Orb_asc, s.(string))
		}
		break
	case "desc":
		switch t.Kind() {
		case reflect.Slice:
			q.Orb_desc = append(q.Orb_desc, s.([]string)...)
		case reflect.String:
			q.Orb_desc = append(q.Orb_desc, s.(string))
		}
		break
	}
	return q
}

func (q *Query) Extract(s *string) {
	*s = `SELECT `
	for k, v := range q.Sel {
		if k < len(q.Sel)-1 {
			*s += v + ", "
		} else {
			*s += v
		}
	}
	*s += " FROM "
	for k, v := range q.Fr {
		if k < len(q.Fr)-1 {
			*s = *s + v + ", "
		} else {
			*s = *s + v
		}
	}
	if len(q.Jo) > 0 {
		*s += " "
		for _, v := range q.Jo {
			*s = *s + v
		}
	}
	if len(q.Wh) > 0 {
		*s += " WHERE "
		for k, v := range q.Wh {
			if k < len(q.Wh)-1 {
				*s = *s + v + " AND "
			} else {
				*s = *s + v
			}
		}
	}
	if len(q.Gr) > 0 {
		*s += ` GROUP BY `
		for k, v := range q.Gr {
			if k < len(q.Gr)-1 {
				*s += v + ", "
			} else {
				*s += v
			}
		}
	}
	if len(q.Orb_asc) > 0 || len(q.Orb_desc) > 0 {
		*s += ` ORDER BY `
		for k, v := range q.Orb_asc {
			if k < len(q.Orb_asc)-1 {
				*s += v + ", "
			} else {
				*s += v + " ASC"
			}
		}
		if len(q.Orb_asc) > 0 && len(q.Orb_desc) > 0 {
			*s += ", "
		}
		for k, v := range q.Orb_desc {
			if k < len(q.Orb_desc)-1 {
				*s += v + ", "
			} else {
				*s += v + " DESC"
			}
		}
	}
}
