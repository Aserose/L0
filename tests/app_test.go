package tests

import (
	"L0/internal/app"
	"L0/internal/datastruct"
	"L0/pkg/customLogger"
	"bytes"
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

type testValue [2]datastruct.Order

func TestApp(t *testing.T) {
	go app.Run()

	logs := customLogger.NewLogger()
	d := data{}
	testData := testValue{
		d.NewTestData("3tr232trf32"),
		d.NewTestData("53gv32vvg2323"),
	}

	testClient := newTestClient(logs)

	convey.Convey("setup", t, func() {

		convey.Convey("purchase info i/o", func() { testClient.ioPurchase(testData) })
		convey.Convey("error info i/o", func() { testClient.ioErrorInfo() })

	})

}

type testClient struct {
	client *http.Client
	log    customLogger.Logger
}

func newTestClient(log customLogger.Logger) testClient {
	return testClient{
		client: &http.Client{},
		log:    log,
	}
}

func (t testClient) do(req *http.Request) []byte {
	resp, err := t.client.Do(req)
	if err != nil {
		t.log.Error(t.log.CallInfoStr(), err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.log.Error(t.log.CallInfoStr(), err.Error())
	}
	return b
}

func (t testClient) ioErrorInfo() {
	errorInfo := datastruct.ErrorInfo{}
	incorrectValues := []string{"incorrect value", "123", `"241"`, `"PUT"`}

	convey.So(string(t.outputError()), convey.ShouldEqual, `none`)
	for _, v := range incorrectValues {
		t.inputPurchase([]byte(v), true)
		time.Sleep(200 * time.Millisecond)

		if err := json.Unmarshal(t.outputError(), &errorInfo); err != nil {
			t.log.Error(t.log.CallInfoStr(), err.Error())
		}

		convey.So(errorInfo.Data, convey.ShouldEqual, v)
	}
}

func (t testClient) outputError() []byte {
	resp, _ := t.client.Get("http://localhost:8000/error")
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}

func (t testClient) ioPurchase(data testValue) {
	var (
		testId   string
		testBool = []bool{false, true}
	)

	input := func(isChannel bool, order datastruct.Order) (id string) {
		byteValue, _ := json.Marshal(order)
		id = t.inputPurchase(byteValue, isChannel)
		if isChannel {
			convey.So(id, convey.ShouldBeEmpty)
			time.Sleep(250 * time.Millisecond)
			return increment(testId)
		}
		convey.So(id, convey.ShouldNotBeEmpty)
		return
	}

	output := func(id string, order datastruct.Order) string {
		for _, isCache := range testBool {
			convey.So(t.outputPurchase(id, isCache).CustomerID, convey.ShouldResemble, order.CustomerID)
		}
		return id
	}

	for i, isChannel := range testBool {
		testId = output(input(isChannel, data[i]), data[i])
	}
}

func (t testClient) outputPurchase(id string, isCache bool) datastruct.Order {
	order := datastruct.Order{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/purchase", nil)
	req.URL.RawQuery = url.Values{
		`id`:    {id},
		`cache`: {t.interpret(isCache)},
	}.Encode()
	if err != nil {
		t.log.Error(t.log.CallInfoStr(), err.Error())
	}

	if err := json.Unmarshal(t.do(req), &order); err != nil {
		t.log.Error(t.log.CallInfoStr(), err.Error())
	}

	return order
}

func (t testClient) inputPurchase(testData []byte, isChannel bool) string {
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/purchase", bytes.NewBuffer(testData))
	req.URL.RawQuery = url.Values{
		"ch": {t.interpret(isChannel)},
	}.Encode()
	if err != nil {
		t.log.Error(t.log.CallInfoStr(), err.Error())
	}

	return string(t.do(req))
}

func (t testClient) interpret(val bool) string {
	if val {
		return `1`
	}
	return `0`
}

func increment(num string) string {
	numInt, _ := strconv.Atoi(num)
	numInt += 1
	return strconv.Itoa(numInt)
}

type data struct{}

func (d data) NewTestData(testCustomerID string) datastruct.Order {
	order := datastruct.Order{}
	json.Unmarshal(d.getTestJson(), &order)

	order.CustomerID = testCustomerID

	return order
}

func (d data) getTestJson() []byte {
	f, _ := os.Open(filepath.Join(d.getDirectory(d.getFilename(), 1), "model.json"))
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	return byteValue
}

func (d data) getDirectory(filename string, depth int) string {
	if depth != 0 {
		return d.getDirectory(filepath.Dir(filename), depth-1)
	}
	return filepath.Dir(filename)
}

func (d data) getFilename() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}
