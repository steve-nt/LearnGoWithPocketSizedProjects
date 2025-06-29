package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello world!!!"

	got := greet()

	if got != want {
		//mark the test as failed
		t.Errorf("expected: %q, got: %q", want, got)

	}

}

/*
func ExampleMain() {
	main() // Calls the main function

	// Output:
	// Hello world
}
*/
/*With the
// Output:
// Hello world
we specify the result of the main function call.
This is useful for testing the output of the main function.
want:
Hello world
FAIL
exit status 1
FAIL    learngo-pockets/hello   0.022s
*/
