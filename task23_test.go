package main

import "testing"

func TestTask23(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	res := RemoveAt(s, 2)
	expect := []int{1, 2, 4, 5}
	if len(res) != 4 || res[2] != 4 || res[3] != 5 {
		t.Errorf("RemoveAt failed: got %v, want %v", res, expect)
	}

	s2 := []int{1, 2, 3}
	res2 := RemoveAt(s2, 0)
	expect2 := []int{2, 3}
	if len(res2) != 2 || res2[0] != 2 || res2[1] != 3 {
		t.Errorf("RemoveAt failed: got %v, want %v", res2, expect2)
	}

	s3 := []int{1, 2, 3}
	res3 := RemoveAt(s3, 2)
	expect3 := []int{1, 2}
	if len(res3) != 2 || res3[0] != 1 || res3[1] != 2 {
		t.Errorf("RemoveAt failed: got %v, want %v", res3, expect3)
	}

	s4 := []string{"a"}
	res4 := RemoveAt(s4, 0)
	if len(res4) != 0 {
		t.Errorf("RemoveAt failed: got %v, want []", res4)
	}

	type bigStruct struct{ data [1024]byte }
	arr := make([]*bigStruct, 2)
	arr[0] = &bigStruct{}
	arr[1] = &bigStruct{}
	arr = RemoveAt(arr, 0)
	if arr[0] == nil {
		t.Errorf("RemoveAt zeroed wrong element")
	}
}
