package hash

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// auto increment は id からサービスの規模や、user の id が判明してしまったり良くない。
// ulid は uuid のデメリットである、処理速度を解消するためのもの
func GetUlid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id.String()
}
