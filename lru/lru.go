package lru

import "container/list"

type Cache struct {
	//最大存储空间
	maxBytes  int64
	//当前存储大小
	nbytes    int64
	//循环链表
	ll 		  *list.List
	//字典中key对应链表的节点
	cache 	  map[string]*list.Element
	OnEvicted func(key string,value Value)
}

type entry struct {
	key	string
	value Value
}

type Value interface {
	Len() int
}

func New(maxByetes int64,onEvicted func(string,Value))*Cache  {
	return &Cache{
		maxBytes:  maxByetes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache)RemoveOldest()  {
	ele := c.ll.Back()
	if ele!=nil{
		c.ll.Remove(ele)
		//将Value接口断言成entry数据类型获取存储的数据
		kv := ele.Value.(*entry)
		delete(c.cache,kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}