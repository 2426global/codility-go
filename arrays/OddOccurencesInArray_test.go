package arrays

import (
"testing"
. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestOddOccurences(c *C) {
	var solutions = [...]int {5, 7, 3, 2, 1}

	c.Assert(OddOccurences([]int{1,1,5,2,2,3,3}), Equals, solutions[0])
	c.Assert(OddOccurences([]int{1,1,5,2,2,3,5,7,3}), Equals, solutions[1])
}

func (s *MySuite) BenchmarkOddOccurences(c *C) {
	for i := 0; i < c.N; i++ {
		OddOccurences([]int{1,1,5,2,2,3,5,7,3})
	}
}
