package main

import (
	"fmt"
)

const LESSON_NUMBER int = 10

type time struct {
	hours   int
	minutes int
	seconds int
}

func createTime(hours, minutes, seconds int) *time {
	newTime := time{hours, minutes, seconds}
	return &newTime
}

func (t time) String() string {
	return fmt.Sprintf("%d:%d:%d", t.hours, t.minutes, t.seconds)
}

type runner struct {
	time 
	fname string
	lname string
}

func createRunner(fname, lname string) *runner {
	newRunner := runner{fname: fname, lname: lname}
	return &newRunner
}

func (r *runner) finish(t time) {
	r.time = t
}

func (r runner) String() string {
	return fmt.Sprintf("%s, %s - %s", r.lname, r.fname, r.time)
}

func PrintWhatILearned() {
	runnerA := createRunner("Ben", "Bryant")
	runnerA.finish(*createTime(0, 15, 30))
	fmt.Println(runnerA)
	runnerB := runner{
		fname: "John",
		lname: "Doe",
		time: time{
			hours: 0,
			minutes: 16,
			seconds: 06,
		},
	}
	fmt.Println(runnerB)
}

func main() {
	fmt.Printf("\n%d. Embedded Structs.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
