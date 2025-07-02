package src

import (
    "math/rand"
    "time"
)

func MightBeUnstable() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64() > 0.3
}
