package purchaseInfo

import (
	"L0/internal/datastruct"
	"L0/internal/repository/db"
	"L0/pkg/customLogger"
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPsqlPurchase(t *testing.T) {
	logs := customLogger.NewLogger()
	p := newTestPsqlPurchase(logs)
	testData := []datastruct.Order{
		newTestValue("1"),
		newTestValue("2"),
		newTestValue("3"),
	}

	convey.Convey("init", t, func() {

		convey.Convey("i/o", func() { p.io(testData) })

	})

}

type testPsqlPurchase struct {
	p PsqlPurchase
}

func newTestPsqlPurchase(logs customLogger.Logger) testPsqlPurchase {
	return testPsqlPurchase{
		p: NewPsqlPurchase(db.NewPsql(logs), logs),
	}
}

func (t testPsqlPurchase) io(orders []datastruct.Order) {
	checkIO := func(sourceOrder datastruct.Order) int {
		id := t.p.Put(sourceOrder)
		decodedOrder := datastruct.Order{}

		for _, val := range [][]byte{t.p.Get(id), t.p.GetCached(id)} {
			json.Unmarshal(val, &decodedOrder)
			convey.So(decodedOrder, convey.ShouldResemble, sourceOrder)
		}

		return id
	}

	checkDelete := func(id int) {
		t.p.Delete(id)
		for _, val := range [][]byte{t.p.Get(id), t.p.GetCached(id)} {
			convey.So(val, convey.ShouldBeNil)
		}
	}

	for _, order := range orders {
		checkDelete(checkIO(order))
	}
}

func newTestValue(customerId string) datastruct.Order {
	return datastruct.Order{
		CustomerID: customerId,
	}
}
