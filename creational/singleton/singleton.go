package singleton

// Singleton is an interface that should be implemented by a counter type
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance should return the current counter object
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

// AddOne will increment the counter by one
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
