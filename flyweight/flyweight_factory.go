package flyweight

import "sync"

var lock sync.Mutex
var obj Flyweight

// GetFlyweight get flyweight by key
func GetFlyweight(key string) Flyweight {
	if key == "1" {
		lock.Lock()
		defer lock.Unlock()

		if obj == nil {
			obj = &ConcretFlyweight{}
		}
		return obj
	}

	panic("key invalid")
}
