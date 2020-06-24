package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nbytes    int64
	ll 		  *list.List
	cache 	  map[string]list.Element
	OnEvicted func(key string,value Value)
}

type entry struct {
	key	string
	vlue Value
}

type Value interface {
	len()
}

