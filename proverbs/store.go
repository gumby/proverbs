package proverbs

import "errors"

type Store interface {
	GetAll() ([]Proverb, error)
	Get(int) (Proverb, error)
}

type Proverb string

type InMemStore struct {
	data map[int]Proverb
}

func NewInMemStore() InMemStore {
	data := make(map[int]Proverb)
	for i, p := range load() {
		data[i] = p
	}
	return InMemStore{
		data: data,
	}
}

func load() []Proverb {
	return []Proverb{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
}
func (s InMemStore) GetAll() ([]Proverb, error) {
	proverbs := make([]Proverb, len(s.data))
	for i, p := range s.data {
		proverbs[i] = p
	}
	return proverbs, nil
}

func (s InMemStore) Get(id int) (Proverb, error) {
	if id >= len(s.data) {
		return "", errors.New("invalid id")
	}
	return s.data[id], nil
}
