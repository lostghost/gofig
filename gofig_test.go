package gofig

import (
	"os"
	"testing"

	. "launchpad.net/gocheck"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) SetUpSuite(c *C) {
	os.Clearenv()
	os.Setenv("FOO", "bar")
}

func (s *BaseSuite) TearDownSuite(c *C) {
	os.Clearenv()
}

func (s *BaseSuite) TestString(c *C) {
	value, err := String("FOO")
	c.Assert(err, IsNil)
	c.Check(value, Equals, "bar")
}

func (s *BaseSuite) TestMissingSting(c *C) {
	_, err := String("NOPE")
	c.Check(err, NotNil)
}

func (s *BaseSuite) TestDefaultString(c *C) {
	value := StringDefault("FOO", "banana")
	c.Check(value, Equals, "bar")
}

func (s *BaseSuite) TestMissingDefaultString(c *C) {
	value := StringDefault("NOPE", "banana")
	c.Check(value, Equals, "banana")
}

func (s *BaseSuite) TestStringVar(c *C) {
	var value string
	err := StringVar(&value, "FOO")
	c.Assert(err, IsNil)
	c.Check(value, Equals, "bar")
}

func (s *BaseSuite) TestMissingStringVar(c *C) {
	var value string
	err := StringVar(&value, "NOPE")
	c.Check(err, NotNil)
}

func (s *BaseSuite) TestDefaultStringVar(c *C) {
	var value string
	StringVarDefault(&value, "FOO", "banana")
	c.Check(value, Equals, "bar")
}

func (s *BaseSuite) TestMissingDefaultStringVar(c *C) {
	var value string
	StringVarDefault(&value, "NOPE", "banana")
}

func (s *BaseSuite) TestInt(c *C) {

}
