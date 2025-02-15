package main

import "testing"

func TestEvenOrOdd(t *testng.T) {
  result := EvenOrOdd(10)
  if result != "even" {
    t.Errorf("expected: even, actual: %s", result)
  }
}
