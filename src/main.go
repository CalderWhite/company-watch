package main

import (
	"gopkg.in/cheggaaa/pb.v1"
	"io/ioutil"
	"math"
	"nameLogic"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// set up classifier
	c, _ := nameLogic.NewClassifier()

	// read in test data
	dat, err := ioutil.ReadFile("test_data.txt")
	check(err)
	s := string(dat)
	// loop on the test data roughly the amount of times we have comments (12 Million)
	scount := int(12 * math.Pow(10, 6))
	pbar := pb.New(count)
	pbar.Start()
	// execute on the test data
	for i := 0; i < count; i++ {
		c.FindNamesInText(s)
		pbar.Increment()
	}
	pbar.FinishPrint("Done.")
}
