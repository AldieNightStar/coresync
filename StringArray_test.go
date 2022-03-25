package coresync

import "testing"

func TestStringArray(t *testing.T) {
	s := `Hello!|Hi!|This\|and\|That|0\\23`
	arr := StringToArray(s)
	if len(arr) != 4 {
		t.Fatal("Array len is not 4: ", len(arr))
	}
	if arr[0] != "Hello!" || arr[1] != "Hi!" || arr[2] != "This|and|That" || arr[3] != `0\23` {
		t.Fatal("Array data, is wrong!: ", arr)
	}
	ss := ArrayToString(arr)
	if ss != s {
		t.Fatal("Difference detected (Orig, New): ", s, ss)
	}
}
