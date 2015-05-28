/*
 * http://golang.org/doc/code.html#Testing
 * http://golang.org/pkg/testing/
 */
package testing

import (
	"testing"
)

func TestFullname(t *testing.T) {
	expected := "Jose Moreno"	
	got := Fullname("Jose", "Moreno")
	if got != expected {
		t.Errorf("Fullname = %q, want %q", got, expected)
	}
}

func BenchmarkFullname(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		Fullname("name", "surname")
	}
}