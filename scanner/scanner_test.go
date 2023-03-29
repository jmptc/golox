package scanner

import "testing"

func TestAtEndOfSource(t *testing.T) {
    scanner := NewScanner("Hello")

    // set current right before end
    scanner.current = len(scanner.source) - 1
    if scanner.AtEnd() {
        t.Errorf("Not expecting AtEnd to be true here. current: %d len: %d", scanner.current, len(scanner.source))
    }

    // set current after source
    scanner.current = len(scanner.source)
    if !scanner.AtEnd() {
        t.Errorf("Expected AtEnd to be true. current: %d len: %d", scanner.current, len(scanner.source))
    }
}
