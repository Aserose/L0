package purchaseInfo

import (
	"L0/pkg/customLogger"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNatCache(t *testing.T) {
	logs := customLogger.NewLogger()
	n := newTestCache(logs)
	testData := [][]byte{
		[]byte("testing"),
		[]byte("test"),
	}

	convey.Convey("init", t, func() {

		convey.Convey("i/o", func() { n.io(testData) })

	})

}

type testNatCache struct {
	n natCache
}

func newTestCache(log customLogger.Logger) testNatCache {
	return testNatCache{
		n: newNatCache(log, make(map[int][]byte)),
	}
}

func (t testNatCache) io(testData [][]byte) {
	checkOutput := func(i int) int {
		t.n.put(i, testData[i])
		convey.So(t.n.get(i), convey.ShouldResemble, testData[i])
		return i
	}

	checkDelete := func(i int) {
		t.n.delete(i)
		convey.So(t.n.get(i), convey.ShouldBeNil)
	}

	for i := 0; i < len(testData); i++ {
		checkDelete(checkOutput(i))
	}
}
