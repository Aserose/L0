package errorInfo

import (
	"L0/pkg/customLogger"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNativeCache(t *testing.T) {
	log := customLogger.NewLogger()

	ec := NewErrCache(log, make([]string, 1))
	testValue := []string{"1", "2", "3"}

	convey.Convey("setup", t, func() {

		convey.So(ec.GetLast(), convey.ShouldEqual, msgNoError)
		for _, v := range testValue {
			ec.Put(v)
			convey.So(ec.GetLast(), convey.ShouldEqual, v)
		}
	})
}
