package xconvey

import "github.com/smartystreets/goconvey/convey"

func NoError(c convey.C, err error) {
	c.So(err, convey.ShouldBeNil)
}

func NotNil(c convey.C, v interface{}) {
	c.So(v, convey.ShouldNotBeNil)
}

func Equals(c convey.C, actual interface{}, expected interface{}) {
	c.So(actual, convey.ShouldEqual, expected)
}

func NotEquals(c convey.C, actual interface{}, expected interface{}) {
	c.So(actual, convey.ShouldNotEqual, expected)
}
