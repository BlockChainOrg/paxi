package wankeeper

import "github.com/ailidani/paxi"

// records of who owns the tokens
type tokens struct {
	id     paxi.ID
	tokens map[paxi.Key]paxi.ID
	// master token manager can create new tokens
	master bool
}

func newTokens(id paxi.ID) *tokens {
	t := &tokens{
		id:     id,
		tokens: make(map[paxi.Key]paxi.ID),
	}
	for i := 0; i <= paxi.GetConfig().Benchmark.K; i++ {
		k := paxi.Key(i + paxi.GetConfig().Benchmark.Min)
		if k == 0 {
			continue
		}
		t.tokens[k] = id
	}
	return t
}

func (t *tokens) contains(key paxi.Key) bool {
	id, exist := t.tokens[key]
	if t.master && !exist {
		t.tokens[key] = t.id
		return true
	}
	return id == t.id
}

func (t *tokens) add(key paxi.Key) {
	t.tokens[key] = t.id
}

func (t *tokens) get(key paxi.Key) paxi.ID {
	return t.tokens[key]
}

func (t *tokens) set(key paxi.Key, id paxi.ID) {
	t.tokens[key] = id
}
