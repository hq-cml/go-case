package set

import (
    "testing"
    "runtime/debug"
)

//必须大写Test开头
func TestHashSetCreation(t *testing.T) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: %s\n", err)
        }
    }()
    t.Log("Starting TestHashSetCreation...")
    hs := NewHashSet()
    t.Logf("Create a HashSet value: %v\n", hs)
    if hs == nil {
        t.Errorf("The result of func NewHashSet is nil!\n")
    }
    isSet := IsSet(hs)
    if !isSet {
        t.Errorf("The value of HashSet is not Set!\n")
    } else {
        t.Logf("The HashSet value is a Set.\n")
    }
}
