package gocopy

// import ( DON'T WORK RIGHT NOW // DON'T WORK RIGHT NOW // DON'T WORK RIGHT NOW // DON'T WORK RIGHT NOW
// 	"fmt"
// 	"io/ioutil"
// 	"path/filepath"
// 	"testing"
// )

// func TestGoCopy(t *testing.T) {
// 	testfile1, err := ioutil.TempFile("", "test1.*.txt")
// 	if err != nil {
// 		t.Fatalf("Testing impossible, cannot create file!")
// 	}
// 	testfile2, err := ioutil.TempFile("", "test2.*.txt")
// 	if err != nil {
// 		t.Fatalf("Testing impossible, cannot create file!")
// 	}
// 	testfile3, err := ioutil.TempFile("", "test3.*.txt")
// 	if err != nil {
// 		t.Fatalf("Testing impossible, cannot create file!")
// 	}
// 	testfile4, err := ioutil.TempFile("", "test4.*.txt")
// 	if err != nil {
// 		t.Fatalf("Testing impossible, cannot create file!")
// 	}
// 	defer os.Remove(testfile1.Name())
// 	defer os.Remove(testfile2.Name())
// 	defer os.Remove(testfile3.Name())
// 	defer os.Remove(testfile4.Name())
// 	exPath1 := filepath.Join("/var/folders/2n/44zwvbbj127bzl5d572f250m0000gn/T/", testfile1.Name())
// 	exPath2 := filepath.Join("/var/folders/2n/44zwvbbj127bzl5d572f250m0000gn/T/", testfile2.Name())
// 	exPath3 := filepath.Join("/var/folders/2n/44zwvbbj127bzl5d572f250m0000gn/T/", testfile3.Name())
// 	exPath4 := filepath.Join("/var/folders/2n/44zwvbbj127bzl5d572f250m0000gn/T/", testfile4.Name())
// 	a1 := make([]byte, 1024)
// 	a2 := make([]byte, 65536)
// 	a3 := make([]byte, 0)
// 	a4 := make([]byte, 0)
// 	for i := 0; i < 1024; i++ {
// 		a1 = append(a1, 1)
// 	}
// 	for i := 0; i < 1024; i++ {
// 		a2 = append(a2, 2)
// 	}
// 	_, err = testfile1.Write(a1)
// 	if err != nil {
// 		t.Fatalf("Cannot write in file 1")
// 	}
// 	_, err = testfile2.Write(a2)
// 	if err != nil {
// 		t.Fatalf("Cannot write in file 2")
// 	}
// 	fmt.Printf("Complete")
// 	fmt.Println(Copy(exPath1, exPath3, 0, 0))
// 	fmt.Println(Copy(exPath2, exPath4, 200, 30000))
// 	fmt.Printf("Complete")
// 	a3, err = ioutil.ReadAll(testfile3)
// 	if err != nil {
// 		t.Fatalf("Cannot read from file 3")
// 	}
// 	a4, err = ioutil.ReadAll(testfile4)
// 	if err != nil {
// 		t.Fatalf("Cannot read from file 4")
// 	}
// 	for j := 0; j <= len(a1); j++ {
// 		if a1[j] != a3[j] {
// 			t.Fatalf("Test failed in first copying")
// 		}
// 	}
// 	for k := 0; k <= len(a1); k++ {
// 		if a2[k] != a4[k] {
// 			t.Fatalf("Test failed in second copying")
// 		}
// 	}
// }
