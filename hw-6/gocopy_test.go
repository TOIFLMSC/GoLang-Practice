package gocopy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGoCopy(t *testing.T) {
	f1test, err := os.Create("filetest1.txt")
	if err != nil {
		t.Fatalf("Testing impossible, cannot create file!")
	}
	f2test, err := os.Create("filetest2.txt")
	if err != nil {
		t.Fatalf("Testing impossible, cannot create file!")
	}
	f3test, err := os.Create("filetest3.txt")
	if err != nil {
		t.Fatalf("Testing impossible, cannot create file!")
	}
	f4test, err := os.Create("filetest4.txt")
	if err != nil {
		t.Fatalf("Testing impossible, cannot create file!")
	}
	defer os.Remove(f1test.Name())
	defer os.Remove(f2test.Name())
	defer os.Remove(f3test.Name())
	defer os.Remove(f4test.Name())
	f1testpath, err := filepath.Abs(filepath.Dir(f1test.Name()))
	if err != nil {
		t.Fatalf("Failed to test, cannot get test fail path")
	}
	f2testpath, err := filepath.Abs(filepath.Dir(f2test.Name()))
	if err != nil {
		t.Fatalf("Failed to test, cannot get test fail path")
	}
	f3testpath, err := filepath.Abs(filepath.Dir(f3test.Name()))
	if err != nil {
		t.Fatalf("Failed to test, cannot get test fail path")
	}
	f4testpath, err := filepath.Abs(filepath.Dir(f4test.Name()))
	if err != nil {
		t.Fatalf("Failed to test, cannot get test fail path")
	}
	exPath1 := filepath.Join(f1testpath, f1test.Name())
	exPath2 := filepath.Join(f2testpath, f2test.Name())
	exPath3 := filepath.Join(f3testpath, f3test.Name())
	exPath4 := filepath.Join(f4testpath, f4test.Name())
	a1 := make([]byte, 1024)
	a2 := make([]byte, 65536)
	a3 := make([]byte, 0)
	a4 := make([]byte, 0)
	for i := 0; i < 1024; i++ {
		a1 = append(a1, 1)
	}
	for i := 0; i < 1024; i++ {
		a2 = append(a2, 2)
	}
	_, err = f1test.Write(a1)
	if err != nil {
		t.Fatalf("Cannot write in file 1")
	}
	_, err = f2test.Write(a2)
	if err != nil {
		t.Fatalf("Cannot write in file 2")
	}
	fmt.Printf("Complete")
	fmt.Println(Copy(exPath1, exPath3, 0, 0))
	fmt.Println(Copy(exPath2, exPath4, 200, 30000))
	fmt.Printf("Complete")
	a3, err = ioutil.ReadAll(f3test)
	if err != nil {
		t.Fatalf("Cannot read from file 3")
	}
	a4, err = ioutil.ReadAll(f4test)
	if err != nil {
		t.Fatalf("Cannot read from file 4")
	}
	fmt.Println(len(a1))
	fmt.Println(len(a3))
	fmt.Printf("%T", a3[0])

	fmt.Printf("%T", a3[0])
	for j := 0; j < len(a1); j++ {
		if a1[j] != a3[j] {
			t.Fatalf("Test failed in first copying")
		}
	}
	fmt.Println(len(a2))
	fmt.Println(len(a4))
	for k := 0; k < len(a1); k++ {
		if a2[k] != a4[k] {
			t.Fatalf("Test failed in second copying")
		}
	}
}
