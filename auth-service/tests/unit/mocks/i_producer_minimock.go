// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka.IProducer -o i_producer_minimock.go -n IProducerMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// IProducerMock implements mm_kafka.IProducer
type IProducerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	funcCloseOrigin    string
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mIProducerMockClose

	funcProduce          func(message string, topic string, key string) (err error)
	funcProduceOrigin    string
	inspectFuncProduce   func(message string, topic string, key string)
	afterProduceCounter  uint64
	beforeProduceCounter uint64
	ProduceMock          mIProducerMockProduce
}

// NewIProducerMock returns a mock for mm_kafka.IProducer
func NewIProducerMock(t minimock.Tester) *IProducerMock {
	m := &IProducerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mIProducerMockClose{mock: m}

	m.ProduceMock = mIProducerMockProduce{mock: m}
	m.ProduceMock.callArgs = []*IProducerMockProduceParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mIProducerMockClose struct {
	optional           bool
	mock               *IProducerMock
	defaultExpectation *IProducerMockCloseExpectation
	expectations       []*IProducerMockCloseExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// IProducerMockCloseExpectation specifies expectation struct of the IProducer.Close
type IProducerMockCloseExpectation struct {
	mock *IProducerMock

	results      *IProducerMockCloseResults
	returnOrigin string
	Counter      uint64
}

// IProducerMockCloseResults contains results of the IProducer.Close
type IProducerMockCloseResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmClose *mIProducerMockClose) Optional() *mIProducerMockClose {
	mmClose.optional = true
	return mmClose
}

// Expect sets up expected params for IProducer.Close
func (mmClose *mIProducerMockClose) Expect() *mIProducerMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("IProducerMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &IProducerMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the IProducer.Close
func (mmClose *mIProducerMockClose) Inspect(f func()) *mIProducerMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for IProducerMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by IProducer.Close
func (mmClose *mIProducerMockClose) Return(err error) *IProducerMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("IProducerMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &IProducerMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &IProducerMockCloseResults{err}
	mmClose.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Set uses given function f to mock the IProducer.Close method
func (mmClose *mIProducerMockClose) Set(f func() (err error)) *IProducerMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the IProducer.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the IProducer.Close method")
	}

	mmClose.mock.funcClose = f
	mmClose.mock.funcCloseOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Times sets number of times IProducer.Close should be invoked
func (mmClose *mIProducerMockClose) Times(n uint64) *mIProducerMockClose {
	if n == 0 {
		mmClose.mock.t.Fatalf("Times of IProducerMock.Close mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmClose.expectedInvocations, n)
	mmClose.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmClose
}

func (mmClose *mIProducerMockClose) invocationsDone() bool {
	if len(mmClose.expectations) == 0 && mmClose.defaultExpectation == nil && mmClose.mock.funcClose == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmClose.mock.afterCloseCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmClose.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Close implements mm_kafka.IProducer
func (mmClose *IProducerMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	mmClose.t.Helper()

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the IProducerMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to IProducerMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished IProducerMock.Close invocations
func (mmClose *IProducerMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of IProducerMock.Close invocations
func (mmClose *IProducerMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *IProducerMock) MinimockCloseDone() bool {
	if m.CloseMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CloseMock.invocationsDone()
}

// MinimockCloseInspect logs each unmet expectation
func (m *IProducerMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to IProducerMock.Close")
		}
	}

	afterCloseCounter := mm_atomic.LoadUint64(&m.afterCloseCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to IProducerMock.Close at\n%s", m.CloseMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to IProducerMock.Close at\n%s", m.funcCloseOrigin)
	}

	if !m.CloseMock.invocationsDone() && afterCloseCounter > 0 {
		m.t.Errorf("Expected %d calls to IProducerMock.Close at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CloseMock.expectedInvocations), m.CloseMock.expectedInvocationsOrigin, afterCloseCounter)
	}
}

type mIProducerMockProduce struct {
	optional           bool
	mock               *IProducerMock
	defaultExpectation *IProducerMockProduceExpectation
	expectations       []*IProducerMockProduceExpectation

	callArgs []*IProducerMockProduceParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// IProducerMockProduceExpectation specifies expectation struct of the IProducer.Produce
type IProducerMockProduceExpectation struct {
	mock               *IProducerMock
	params             *IProducerMockProduceParams
	paramPtrs          *IProducerMockProduceParamPtrs
	expectationOrigins IProducerMockProduceExpectationOrigins
	results            *IProducerMockProduceResults
	returnOrigin       string
	Counter            uint64
}

// IProducerMockProduceParams contains parameters of the IProducer.Produce
type IProducerMockProduceParams struct {
	message string
	topic   string
	key     string
}

// IProducerMockProduceParamPtrs contains pointers to parameters of the IProducer.Produce
type IProducerMockProduceParamPtrs struct {
	message *string
	topic   *string
	key     *string
}

// IProducerMockProduceResults contains results of the IProducer.Produce
type IProducerMockProduceResults struct {
	err error
}

// IProducerMockProduceOrigins contains origins of expectations of the IProducer.Produce
type IProducerMockProduceExpectationOrigins struct {
	origin        string
	originMessage string
	originTopic   string
	originKey     string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmProduce *mIProducerMockProduce) Optional() *mIProducerMockProduce {
	mmProduce.optional = true
	return mmProduce
}

// Expect sets up expected params for IProducer.Produce
func (mmProduce *mIProducerMockProduce) Expect(message string, topic string, key string) *mIProducerMockProduce {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	if mmProduce.defaultExpectation == nil {
		mmProduce.defaultExpectation = &IProducerMockProduceExpectation{}
	}

	if mmProduce.defaultExpectation.paramPtrs != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by ExpectParams functions")
	}

	mmProduce.defaultExpectation.params = &IProducerMockProduceParams{message, topic, key}
	mmProduce.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmProduce.expectations {
		if minimock.Equal(e.params, mmProduce.defaultExpectation.params) {
			mmProduce.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmProduce.defaultExpectation.params)
		}
	}

	return mmProduce
}

// ExpectMessageParam1 sets up expected param message for IProducer.Produce
func (mmProduce *mIProducerMockProduce) ExpectMessageParam1(message string) *mIProducerMockProduce {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	if mmProduce.defaultExpectation == nil {
		mmProduce.defaultExpectation = &IProducerMockProduceExpectation{}
	}

	if mmProduce.defaultExpectation.params != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Expect")
	}

	if mmProduce.defaultExpectation.paramPtrs == nil {
		mmProduce.defaultExpectation.paramPtrs = &IProducerMockProduceParamPtrs{}
	}
	mmProduce.defaultExpectation.paramPtrs.message = &message
	mmProduce.defaultExpectation.expectationOrigins.originMessage = minimock.CallerInfo(1)

	return mmProduce
}

// ExpectTopicParam2 sets up expected param topic for IProducer.Produce
func (mmProduce *mIProducerMockProduce) ExpectTopicParam2(topic string) *mIProducerMockProduce {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	if mmProduce.defaultExpectation == nil {
		mmProduce.defaultExpectation = &IProducerMockProduceExpectation{}
	}

	if mmProduce.defaultExpectation.params != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Expect")
	}

	if mmProduce.defaultExpectation.paramPtrs == nil {
		mmProduce.defaultExpectation.paramPtrs = &IProducerMockProduceParamPtrs{}
	}
	mmProduce.defaultExpectation.paramPtrs.topic = &topic
	mmProduce.defaultExpectation.expectationOrigins.originTopic = minimock.CallerInfo(1)

	return mmProduce
}

// ExpectKeyParam3 sets up expected param key for IProducer.Produce
func (mmProduce *mIProducerMockProduce) ExpectKeyParam3(key string) *mIProducerMockProduce {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	if mmProduce.defaultExpectation == nil {
		mmProduce.defaultExpectation = &IProducerMockProduceExpectation{}
	}

	if mmProduce.defaultExpectation.params != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Expect")
	}

	if mmProduce.defaultExpectation.paramPtrs == nil {
		mmProduce.defaultExpectation.paramPtrs = &IProducerMockProduceParamPtrs{}
	}
	mmProduce.defaultExpectation.paramPtrs.key = &key
	mmProduce.defaultExpectation.expectationOrigins.originKey = minimock.CallerInfo(1)

	return mmProduce
}

// Inspect accepts an inspector function that has same arguments as the IProducer.Produce
func (mmProduce *mIProducerMockProduce) Inspect(f func(message string, topic string, key string)) *mIProducerMockProduce {
	if mmProduce.mock.inspectFuncProduce != nil {
		mmProduce.mock.t.Fatalf("Inspect function is already set for IProducerMock.Produce")
	}

	mmProduce.mock.inspectFuncProduce = f

	return mmProduce
}

// Return sets up results that will be returned by IProducer.Produce
func (mmProduce *mIProducerMockProduce) Return(err error) *IProducerMock {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	if mmProduce.defaultExpectation == nil {
		mmProduce.defaultExpectation = &IProducerMockProduceExpectation{mock: mmProduce.mock}
	}
	mmProduce.defaultExpectation.results = &IProducerMockProduceResults{err}
	mmProduce.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmProduce.mock
}

// Set uses given function f to mock the IProducer.Produce method
func (mmProduce *mIProducerMockProduce) Set(f func(message string, topic string, key string) (err error)) *IProducerMock {
	if mmProduce.defaultExpectation != nil {
		mmProduce.mock.t.Fatalf("Default expectation is already set for the IProducer.Produce method")
	}

	if len(mmProduce.expectations) > 0 {
		mmProduce.mock.t.Fatalf("Some expectations are already set for the IProducer.Produce method")
	}

	mmProduce.mock.funcProduce = f
	mmProduce.mock.funcProduceOrigin = minimock.CallerInfo(1)
	return mmProduce.mock
}

// When sets expectation for the IProducer.Produce which will trigger the result defined by the following
// Then helper
func (mmProduce *mIProducerMockProduce) When(message string, topic string, key string) *IProducerMockProduceExpectation {
	if mmProduce.mock.funcProduce != nil {
		mmProduce.mock.t.Fatalf("IProducerMock.Produce mock is already set by Set")
	}

	expectation := &IProducerMockProduceExpectation{
		mock:               mmProduce.mock,
		params:             &IProducerMockProduceParams{message, topic, key},
		expectationOrigins: IProducerMockProduceExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmProduce.expectations = append(mmProduce.expectations, expectation)
	return expectation
}

// Then sets up IProducer.Produce return parameters for the expectation previously defined by the When method
func (e *IProducerMockProduceExpectation) Then(err error) *IProducerMock {
	e.results = &IProducerMockProduceResults{err}
	return e.mock
}

// Times sets number of times IProducer.Produce should be invoked
func (mmProduce *mIProducerMockProduce) Times(n uint64) *mIProducerMockProduce {
	if n == 0 {
		mmProduce.mock.t.Fatalf("Times of IProducerMock.Produce mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmProduce.expectedInvocations, n)
	mmProduce.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmProduce
}

func (mmProduce *mIProducerMockProduce) invocationsDone() bool {
	if len(mmProduce.expectations) == 0 && mmProduce.defaultExpectation == nil && mmProduce.mock.funcProduce == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmProduce.mock.afterProduceCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmProduce.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Produce implements mm_kafka.IProducer
func (mmProduce *IProducerMock) Produce(message string, topic string, key string) (err error) {
	mm_atomic.AddUint64(&mmProduce.beforeProduceCounter, 1)
	defer mm_atomic.AddUint64(&mmProduce.afterProduceCounter, 1)

	mmProduce.t.Helper()

	if mmProduce.inspectFuncProduce != nil {
		mmProduce.inspectFuncProduce(message, topic, key)
	}

	mm_params := IProducerMockProduceParams{message, topic, key}

	// Record call args
	mmProduce.ProduceMock.mutex.Lock()
	mmProduce.ProduceMock.callArgs = append(mmProduce.ProduceMock.callArgs, &mm_params)
	mmProduce.ProduceMock.mutex.Unlock()

	for _, e := range mmProduce.ProduceMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmProduce.ProduceMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmProduce.ProduceMock.defaultExpectation.Counter, 1)
		mm_want := mmProduce.ProduceMock.defaultExpectation.params
		mm_want_ptrs := mmProduce.ProduceMock.defaultExpectation.paramPtrs

		mm_got := IProducerMockProduceParams{message, topic, key}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.message != nil && !minimock.Equal(*mm_want_ptrs.message, mm_got.message) {
				mmProduce.t.Errorf("IProducerMock.Produce got unexpected parameter message, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmProduce.ProduceMock.defaultExpectation.expectationOrigins.originMessage, *mm_want_ptrs.message, mm_got.message, minimock.Diff(*mm_want_ptrs.message, mm_got.message))
			}

			if mm_want_ptrs.topic != nil && !minimock.Equal(*mm_want_ptrs.topic, mm_got.topic) {
				mmProduce.t.Errorf("IProducerMock.Produce got unexpected parameter topic, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmProduce.ProduceMock.defaultExpectation.expectationOrigins.originTopic, *mm_want_ptrs.topic, mm_got.topic, minimock.Diff(*mm_want_ptrs.topic, mm_got.topic))
			}

			if mm_want_ptrs.key != nil && !minimock.Equal(*mm_want_ptrs.key, mm_got.key) {
				mmProduce.t.Errorf("IProducerMock.Produce got unexpected parameter key, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmProduce.ProduceMock.defaultExpectation.expectationOrigins.originKey, *mm_want_ptrs.key, mm_got.key, minimock.Diff(*mm_want_ptrs.key, mm_got.key))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmProduce.t.Errorf("IProducerMock.Produce got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmProduce.ProduceMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmProduce.ProduceMock.defaultExpectation.results
		if mm_results == nil {
			mmProduce.t.Fatal("No results are set for the IProducerMock.Produce")
		}
		return (*mm_results).err
	}
	if mmProduce.funcProduce != nil {
		return mmProduce.funcProduce(message, topic, key)
	}
	mmProduce.t.Fatalf("Unexpected call to IProducerMock.Produce. %v %v %v", message, topic, key)
	return
}

// ProduceAfterCounter returns a count of finished IProducerMock.Produce invocations
func (mmProduce *IProducerMock) ProduceAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProduce.afterProduceCounter)
}

// ProduceBeforeCounter returns a count of IProducerMock.Produce invocations
func (mmProduce *IProducerMock) ProduceBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProduce.beforeProduceCounter)
}

// Calls returns a list of arguments used in each call to IProducerMock.Produce.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmProduce *mIProducerMockProduce) Calls() []*IProducerMockProduceParams {
	mmProduce.mutex.RLock()

	argCopy := make([]*IProducerMockProduceParams, len(mmProduce.callArgs))
	copy(argCopy, mmProduce.callArgs)

	mmProduce.mutex.RUnlock()

	return argCopy
}

// MinimockProduceDone returns true if the count of the Produce invocations corresponds
// the number of defined expectations
func (m *IProducerMock) MinimockProduceDone() bool {
	if m.ProduceMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ProduceMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ProduceMock.invocationsDone()
}

// MinimockProduceInspect logs each unmet expectation
func (m *IProducerMock) MinimockProduceInspect() {
	for _, e := range m.ProduceMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IProducerMock.Produce at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterProduceCounter := mm_atomic.LoadUint64(&m.afterProduceCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ProduceMock.defaultExpectation != nil && afterProduceCounter < 1 {
		if m.ProduceMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to IProducerMock.Produce at\n%s", m.ProduceMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to IProducerMock.Produce at\n%s with params: %#v", m.ProduceMock.defaultExpectation.expectationOrigins.origin, *m.ProduceMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProduce != nil && afterProduceCounter < 1 {
		m.t.Errorf("Expected call to IProducerMock.Produce at\n%s", m.funcProduceOrigin)
	}

	if !m.ProduceMock.invocationsDone() && afterProduceCounter > 0 {
		m.t.Errorf("Expected %d calls to IProducerMock.Produce at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.ProduceMock.expectedInvocations), m.ProduceMock.expectedInvocationsOrigin, afterProduceCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IProducerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockProduceInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IProducerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *IProducerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockProduceDone()
}
