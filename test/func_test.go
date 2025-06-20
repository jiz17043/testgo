package test

import (
    "testing"

    "github.com/jiz17043/testgo/src"
)

func TestMightBeUnstable(t *testing.T) {
    if !src.MightBeUnstable() {
        t.Fatal("test failed (flaky)")
    }
}
