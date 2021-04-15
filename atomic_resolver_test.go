package snowflake_test

import (
	"testing"

	"github.com/godruoyi/go-snowflake"
)

func TestAtomicResolver(t *testing.T) {
	id, _ := snowflake.AtomicResolver(1)

	if id != 0 {
		t.Error("Sequence should be equal 0")
	}
}

func BenchmarkCombinationParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			snowflake.AtomicResolver(1)
		}
	})
}

func BenchmarkAtomicResolver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		snowflake.AtomicResolver(1)
	}
}
