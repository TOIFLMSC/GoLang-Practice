package gocopy

import (
	"flag" // importing flag package to work with flags in std
	"fmt"
	"io"        // importing io package to work with files
	"io/ioutil" // importing ioutil  to work with files
	"os"        // importing flag package to work with files, permissions and etc
	"time"

	"github.com/mitchellh/ioprogress" // importing ioprogress package to show visual animation of working with files
)

var from string
var to string
var offset int
var limit int

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to copy to")
	flag.IntVar(&offset, "offset", 0, "starting position of copying")
	flag.IntVar(&limit, "limit", 0, "value of copying")
}

func Copy(from string, to string, offset int64, limit int) error { // Main function for copying bytes from one file to another file

	buf := make([]byte, limit) // buffer for reading from first file, has limit argument to control amount of bytes which it can fit

	ffile, err := os.OpenFile(from, os.O_RDWR, 0644) // Opening first file for read and write with 0644 permisions

	if err != nil {
		return err
	}

	fdfile, err := os.OpenFile(from, os.O_RDWR, 0644) // Opening first file for read and write with 0644 permisions
	defer fdfile.Close()                              // defer that deletes descriptor for first file
	if err != nil {
		return err
	}

	bufdefault, err := ioutil.ReadAll(fdfile) // "reading" full file to the buffer, that will save all it's data
	if err != nil {
		return err
	}

	ifile, err := os.OpenFile(to, os.O_RDWR, 0644) // Opening second file for read and write with 0644 permisions
	defer ifile.Close()                            // defer that deletes descriptor for second file

	if limit == 0 && offset == 0 { // depending on the input data, we read from the file into the buffer for future writing
		buf, err = ioutil.ReadAll(ffile)
		if err != nil {
			return err
		}
	} else if limit == 0 && offset != 0 {
		ffile.Seek(offset, 0) // here we are shifting the descriptor by the required number of bits
		buf, err = ioutil.ReadAll(ffile)
		if err != nil {
			return err
		}
	} else {
		_, err = ffile.ReadAt(buf, offset)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(from, buf, 0644) // writing (copying) an information from the buffer into our second file
	if err != nil {
		return err
	}

	ffile.Close() // manual deliting of a descriptor for the first file to nullify seeker and rewrite file

	f2file, err := os.OpenFile(from, os.O_RDWR, 0644) // Opening first file for read and write with 0644 permisions
	defer f2file.Close()                              // defer that deletes descriptor for first file
	if err != nil {
		return err
	}

	ffilestats, err := f2file.Stat() // creating Stat file that has main information about first file
	if err != nil {
		return err
	}

	progressR := &ioprogress.Reader{ // creating an exemplar of special Reader that will show us reading progress and initialising it
		Reader:       f2file,            // file that implements interface os.Reader
		Size:         ffilestats.Size(), // size of the first file in bytes
		DrawInterval: time.Nanosecond,   // interval between operations in drawing
	}

	bufdel := make([]byte, 1) // creating a buffer with unit length for writing empty file

	err = ioutil.WriteFile(to, bufdel, 0644) // writing into second file from "empty" buffer, after this operation we will get empty file,
	if err != nil {                          // this operation clears file
		return err
	}

	buffer := make([]byte, 1) // creating a buffer with unit length for copying using buffer

	_, err = io.CopyBuffer(ifile, progressR, buffer) // copying an information from our special progress reader(first file) to the secong file
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(from, bufdefault, 0644) // writing back information in first file, which was wiped two actions back
	if err != nil {
		return err
	}

	return nil

}

func main() {
	flag.Parse()                                      // here we parse flags that we have typed in command line call
	fmt.Println(Copy(from, to, int64(offset), limit)) // executing copy function, that will copy information given the input from the command line call
}
