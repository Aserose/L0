package streamer

import (
	"L0/pkg/customLogger"
	"github.com/nats-io/stan.go"
	"github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func TestNats(t *testing.T) {
	log := customLogger.NewLogger()
	st := newNatStream(log)
	subjectName := "test_test"

	convey.Convey("setup", t, func() {
		var msgFromChan []byte

		st.addChannel(subjectName, func(msg *stan.Msg) {
			msgFromChan = msg.Data
		})

		for _, t := range newTestData(newStrSlice(6)...) {
			convey.So(st.sendMsg(subjectName, t), convey.ShouldBeNil)
			time.Sleep(260 * time.Millisecond)
			convey.So(msgFromChan, convey.ShouldResemble, t)
		}
	})
}

func newTestData(msg ...string) [][]byte {
	testMsg := make([][]byte, len(msg))

	for i, m := range msg {
		testMsg[i] = []byte(m)
	}

	return testMsg
}

func newStrSlice(size int) []string {
	msgs := make([]string, size)

	for i := 0; i < size; i++ {
		msgs[i] = "testMsgs" + strconv.Itoa(i)
	}

	return msgs
}
