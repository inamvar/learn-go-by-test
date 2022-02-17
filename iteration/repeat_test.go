package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){

	 assertCorrectOutput := func(t testing.TB, repeated, expected string){
		t.Helper()
		if repeated != expected{
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	 }

	 t.Run("Repeat 'a' for 5 times", func(t *testing.T){
		repeated := Repeat("a", 5);
		expected := "aaaaa"
		assertCorrectOutput(t, repeated, expected)
	 })

    
	
}

func BenchmarkRepeat( b *testing.B){
	for i:=0 ; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	repeated :=  Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}