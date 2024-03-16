package cache

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) SetCache(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) GetCache(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("empty key")
	}
	return c.data[key], nil
}

func (c *Cache) DeleteCache(key string) {
	delete(c.data, key)
}
