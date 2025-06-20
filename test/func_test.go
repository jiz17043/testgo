package test

import (
    "testing"

    "testgo/src"
)

func TestMightBeUnstable(t *testing.T) {
    if !src.MightBeUnstable() {
        t.Fatal("test failed (flaky)")
    }
}
