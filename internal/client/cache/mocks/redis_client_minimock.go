// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/biryanim/leadgenmarket_tz/internal/client/cache.RedisClient -o redis_client_minimock.go -n RedisClientMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RedisClientMock implements mm_cache.RedisClient
type RedisClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	funcCloseOrigin    string
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mRedisClientMockClose

	funcGet          func(ctx context.Context, key string) (p1 interface{}, err error)
	funcGetOrigin    string
	inspectFuncGet   func(ctx context.Context, key string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mRedisClientMockGet

	funcSet          func(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
	funcSetOrigin    string
	inspectFuncSet   func(ctx context.Context, key string, value interface{}, expiration time.Duration)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mRedisClientMockSet
}

// NewRedisClientMock returns a mock for mm_cache.RedisClient
func NewRedisClientMock(t minimock.Tester) *RedisClientMock {
	m := &RedisClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mRedisClientMockClose{mock: m}

	m.GetMock = mRedisClientMockGet{mock: m}
	m.GetMock.callArgs = []*RedisClientMockGetParams{}

	m.SetMock = mRedisClientMockSet{mock: m}
	m.SetMock.callArgs = []*RedisClientMockSetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRedisClientMockClose struct {
	optional           bool
	mock               *RedisClientMock
	defaultExpectation *RedisClientMockCloseExpectation
	expectations       []*RedisClientMockCloseExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// RedisClientMockCloseExpectation specifies expectation struct of the RedisClient.Close
type RedisClientMockCloseExpectation struct {
	mock *RedisClientMock

	results      *RedisClientMockCloseResults
	returnOrigin string
	Counter      uint64
}

// RedisClientMockCloseResults contains results of the RedisClient.Close
type RedisClientMockCloseResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmClose *mRedisClientMockClose) Optional() *mRedisClientMockClose {
	mmClose.optional = true
	return mmClose
}

// Expect sets up expected params for RedisClient.Close
func (mmClose *mRedisClientMockClose) Expect() *mRedisClientMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("RedisClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &RedisClientMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the RedisClient.Close
func (mmClose *mRedisClientMockClose) Inspect(f func()) *mRedisClientMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for RedisClientMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by RedisClient.Close
func (mmClose *mRedisClientMockClose) Return(err error) *RedisClientMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("RedisClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &RedisClientMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &RedisClientMockCloseResults{err}
	mmClose.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Set uses given function f to mock the RedisClient.Close method
func (mmClose *mRedisClientMockClose) Set(f func() (err error)) *RedisClientMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the RedisClient.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the RedisClient.Close method")
	}

	mmClose.mock.funcClose = f
	mmClose.mock.funcCloseOrigin = minimock.CallerInfo(1)
	return mmClose.mock
}

// Times sets number of times RedisClient.Close should be invoked
func (mmClose *mRedisClientMockClose) Times(n uint64) *mRedisClientMockClose {
	if n == 0 {
		mmClose.mock.t.Fatalf("Times of RedisClientMock.Close mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmClose.expectedInvocations, n)
	mmClose.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmClose
}

func (mmClose *mRedisClientMockClose) invocationsDone() bool {
	if len(mmClose.expectations) == 0 && mmClose.defaultExpectation == nil && mmClose.mock.funcClose == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmClose.mock.afterCloseCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmClose.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Close implements mm_cache.RedisClient
func (mmClose *RedisClientMock) Close() (err error) {
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
			mmClose.t.Fatal("No results are set for the RedisClientMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to RedisClientMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished RedisClientMock.Close invocations
func (mmClose *RedisClientMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of RedisClientMock.Close invocations
func (mmClose *RedisClientMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *RedisClientMock) MinimockCloseDone() bool {
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
func (m *RedisClientMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RedisClientMock.Close")
		}
	}

	afterCloseCounter := mm_atomic.LoadUint64(&m.afterCloseCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to RedisClientMock.Close at\n%s", m.CloseMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && afterCloseCounter < 1 {
		m.t.Errorf("Expected call to RedisClientMock.Close at\n%s", m.funcCloseOrigin)
	}

	if !m.CloseMock.invocationsDone() && afterCloseCounter > 0 {
		m.t.Errorf("Expected %d calls to RedisClientMock.Close at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CloseMock.expectedInvocations), m.CloseMock.expectedInvocationsOrigin, afterCloseCounter)
	}
}

type mRedisClientMockGet struct {
	optional           bool
	mock               *RedisClientMock
	defaultExpectation *RedisClientMockGetExpectation
	expectations       []*RedisClientMockGetExpectation

	callArgs []*RedisClientMockGetParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// RedisClientMockGetExpectation specifies expectation struct of the RedisClient.Get
type RedisClientMockGetExpectation struct {
	mock               *RedisClientMock
	params             *RedisClientMockGetParams
	paramPtrs          *RedisClientMockGetParamPtrs
	expectationOrigins RedisClientMockGetExpectationOrigins
	results            *RedisClientMockGetResults
	returnOrigin       string
	Counter            uint64
}

// RedisClientMockGetParams contains parameters of the RedisClient.Get
type RedisClientMockGetParams struct {
	ctx context.Context
	key string
}

// RedisClientMockGetParamPtrs contains pointers to parameters of the RedisClient.Get
type RedisClientMockGetParamPtrs struct {
	ctx *context.Context
	key *string
}

// RedisClientMockGetResults contains results of the RedisClient.Get
type RedisClientMockGetResults struct {
	p1  interface{}
	err error
}

// RedisClientMockGetOrigins contains origins of expectations of the RedisClient.Get
type RedisClientMockGetExpectationOrigins struct {
	origin    string
	originCtx string
	originKey string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mRedisClientMockGet) Optional() *mRedisClientMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for RedisClient.Get
func (mmGet *mRedisClientMockGet) Expect(ctx context.Context, key string) *mRedisClientMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisClientMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &RedisClientMockGetParams{ctx, key}
	mmGet.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for RedisClient.Get
func (mmGet *mRedisClientMockGet) ExpectCtxParam1(ctx context.Context) *mRedisClientMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisClientMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &RedisClientMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx
	mmGet.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGet
}

// ExpectKeyParam2 sets up expected param key for RedisClient.Get
func (mmGet *mRedisClientMockGet) ExpectKeyParam2(key string) *mRedisClientMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisClientMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &RedisClientMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.key = &key
	mmGet.defaultExpectation.expectationOrigins.originKey = minimock.CallerInfo(1)

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the RedisClient.Get
func (mmGet *mRedisClientMockGet) Inspect(f func(ctx context.Context, key string)) *mRedisClientMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for RedisClientMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by RedisClient.Get
func (mmGet *mRedisClientMockGet) Return(p1 interface{}, err error) *RedisClientMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisClientMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &RedisClientMockGetResults{p1, err}
	mmGet.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// Set uses given function f to mock the RedisClient.Get method
func (mmGet *mRedisClientMockGet) Set(f func(ctx context.Context, key string) (p1 interface{}, err error)) *RedisClientMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the RedisClient.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the RedisClient.Get method")
	}

	mmGet.mock.funcGet = f
	mmGet.mock.funcGetOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// When sets expectation for the RedisClient.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mRedisClientMockGet) When(ctx context.Context, key string) *RedisClientMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisClientMock.Get mock is already set by Set")
	}

	expectation := &RedisClientMockGetExpectation{
		mock:               mmGet.mock,
		params:             &RedisClientMockGetParams{ctx, key},
		expectationOrigins: RedisClientMockGetExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up RedisClient.Get return parameters for the expectation previously defined by the When method
func (e *RedisClientMockGetExpectation) Then(p1 interface{}, err error) *RedisClientMock {
	e.results = &RedisClientMockGetResults{p1, err}
	return e.mock
}

// Times sets number of times RedisClient.Get should be invoked
func (mmGet *mRedisClientMockGet) Times(n uint64) *mRedisClientMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of RedisClientMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	mmGet.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGet
}

func (mmGet *mRedisClientMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements mm_cache.RedisClient
func (mmGet *RedisClientMock) Get(ctx context.Context, key string) (p1 interface{}, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	mmGet.t.Helper()

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, key)
	}

	mm_params := RedisClientMockGetParams{ctx, key}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := RedisClientMockGetParams{ctx, key}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("RedisClientMock.Get got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.key != nil && !minimock.Equal(*mm_want_ptrs.key, mm_got.key) {
				mmGet.t.Errorf("RedisClientMock.Get got unexpected parameter key, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originKey, *mm_want_ptrs.key, mm_got.key, minimock.Diff(*mm_want_ptrs.key, mm_got.key))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("RedisClientMock.Get got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGet.GetMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the RedisClientMock.Get")
		}
		return (*mm_results).p1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, key)
	}
	mmGet.t.Fatalf("Unexpected call to RedisClientMock.Get. %v %v", ctx, key)
	return
}

// GetAfterCounter returns a count of finished RedisClientMock.Get invocations
func (mmGet *RedisClientMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of RedisClientMock.Get invocations
func (mmGet *RedisClientMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to RedisClientMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mRedisClientMockGet) Calls() []*RedisClientMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*RedisClientMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *RedisClientMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *RedisClientMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisClientMock.Get at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to RedisClientMock.Get at\n%s", m.GetMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to RedisClientMock.Get at\n%s with params: %#v", m.GetMock.defaultExpectation.expectationOrigins.origin, *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Errorf("Expected call to RedisClientMock.Get at\n%s", m.funcGetOrigin)
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to RedisClientMock.Get at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), m.GetMock.expectedInvocationsOrigin, afterGetCounter)
	}
}

type mRedisClientMockSet struct {
	optional           bool
	mock               *RedisClientMock
	defaultExpectation *RedisClientMockSetExpectation
	expectations       []*RedisClientMockSetExpectation

	callArgs []*RedisClientMockSetParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// RedisClientMockSetExpectation specifies expectation struct of the RedisClient.Set
type RedisClientMockSetExpectation struct {
	mock               *RedisClientMock
	params             *RedisClientMockSetParams
	paramPtrs          *RedisClientMockSetParamPtrs
	expectationOrigins RedisClientMockSetExpectationOrigins
	results            *RedisClientMockSetResults
	returnOrigin       string
	Counter            uint64
}

// RedisClientMockSetParams contains parameters of the RedisClient.Set
type RedisClientMockSetParams struct {
	ctx        context.Context
	key        string
	value      interface{}
	expiration time.Duration
}

// RedisClientMockSetParamPtrs contains pointers to parameters of the RedisClient.Set
type RedisClientMockSetParamPtrs struct {
	ctx        *context.Context
	key        *string
	value      *interface{}
	expiration *time.Duration
}

// RedisClientMockSetResults contains results of the RedisClient.Set
type RedisClientMockSetResults struct {
	err error
}

// RedisClientMockSetOrigins contains origins of expectations of the RedisClient.Set
type RedisClientMockSetExpectationOrigins struct {
	origin           string
	originCtx        string
	originKey        string
	originValue      string
	originExpiration string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSet *mRedisClientMockSet) Optional() *mRedisClientMockSet {
	mmSet.optional = true
	return mmSet
}

// Expect sets up expected params for RedisClient.Set
func (mmSet *mRedisClientMockSet) Expect(ctx context.Context, key string, value interface{}, expiration time.Duration) *mRedisClientMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{}
	}

	if mmSet.defaultExpectation.paramPtrs != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by ExpectParams functions")
	}

	mmSet.defaultExpectation.params = &RedisClientMockSetParams{ctx, key, value, expiration}
	mmSet.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// ExpectCtxParam1 sets up expected param ctx for RedisClient.Set
func (mmSet *mRedisClientMockSet) ExpectCtxParam1(ctx context.Context) *mRedisClientMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &RedisClientMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.ctx = &ctx
	mmSet.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmSet
}

// ExpectKeyParam2 sets up expected param key for RedisClient.Set
func (mmSet *mRedisClientMockSet) ExpectKeyParam2(key string) *mRedisClientMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &RedisClientMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.key = &key
	mmSet.defaultExpectation.expectationOrigins.originKey = minimock.CallerInfo(1)

	return mmSet
}

// ExpectValueParam3 sets up expected param value for RedisClient.Set
func (mmSet *mRedisClientMockSet) ExpectValueParam3(value interface{}) *mRedisClientMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &RedisClientMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.value = &value
	mmSet.defaultExpectation.expectationOrigins.originValue = minimock.CallerInfo(1)

	return mmSet
}

// ExpectExpirationParam4 sets up expected param expiration for RedisClient.Set
func (mmSet *mRedisClientMockSet) ExpectExpirationParam4(expiration time.Duration) *mRedisClientMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{}
	}

	if mmSet.defaultExpectation.params != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Expect")
	}

	if mmSet.defaultExpectation.paramPtrs == nil {
		mmSet.defaultExpectation.paramPtrs = &RedisClientMockSetParamPtrs{}
	}
	mmSet.defaultExpectation.paramPtrs.expiration = &expiration
	mmSet.defaultExpectation.expectationOrigins.originExpiration = minimock.CallerInfo(1)

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the RedisClient.Set
func (mmSet *mRedisClientMockSet) Inspect(f func(ctx context.Context, key string, value interface{}, expiration time.Duration)) *mRedisClientMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for RedisClientMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by RedisClient.Set
func (mmSet *mRedisClientMockSet) Return(err error) *RedisClientMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisClientMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &RedisClientMockSetResults{err}
	mmSet.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSet.mock
}

// Set uses given function f to mock the RedisClient.Set method
func (mmSet *mRedisClientMockSet) Set(f func(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)) *RedisClientMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the RedisClient.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the RedisClient.Set method")
	}

	mmSet.mock.funcSet = f
	mmSet.mock.funcSetOrigin = minimock.CallerInfo(1)
	return mmSet.mock
}

// When sets expectation for the RedisClient.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mRedisClientMockSet) When(ctx context.Context, key string, value interface{}, expiration time.Duration) *RedisClientMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisClientMock.Set mock is already set by Set")
	}

	expectation := &RedisClientMockSetExpectation{
		mock:               mmSet.mock,
		params:             &RedisClientMockSetParams{ctx, key, value, expiration},
		expectationOrigins: RedisClientMockSetExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up RedisClient.Set return parameters for the expectation previously defined by the When method
func (e *RedisClientMockSetExpectation) Then(err error) *RedisClientMock {
	e.results = &RedisClientMockSetResults{err}
	return e.mock
}

// Times sets number of times RedisClient.Set should be invoked
func (mmSet *mRedisClientMockSet) Times(n uint64) *mRedisClientMockSet {
	if n == 0 {
		mmSet.mock.t.Fatalf("Times of RedisClientMock.Set mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSet.expectedInvocations, n)
	mmSet.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSet
}

func (mmSet *mRedisClientMockSet) invocationsDone() bool {
	if len(mmSet.expectations) == 0 && mmSet.defaultExpectation == nil && mmSet.mock.funcSet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSet.mock.afterSetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Set implements mm_cache.RedisClient
func (mmSet *RedisClientMock) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	mmSet.t.Helper()

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, key, value, expiration)
	}

	mm_params := RedisClientMockSetParams{ctx, key, value, expiration}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, &mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_want_ptrs := mmSet.SetMock.defaultExpectation.paramPtrs

		mm_got := RedisClientMockSetParams{ctx, key, value, expiration}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSet.t.Errorf("RedisClientMock.Set got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSet.SetMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.key != nil && !minimock.Equal(*mm_want_ptrs.key, mm_got.key) {
				mmSet.t.Errorf("RedisClientMock.Set got unexpected parameter key, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSet.SetMock.defaultExpectation.expectationOrigins.originKey, *mm_want_ptrs.key, mm_got.key, minimock.Diff(*mm_want_ptrs.key, mm_got.key))
			}

			if mm_want_ptrs.value != nil && !minimock.Equal(*mm_want_ptrs.value, mm_got.value) {
				mmSet.t.Errorf("RedisClientMock.Set got unexpected parameter value, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSet.SetMock.defaultExpectation.expectationOrigins.originValue, *mm_want_ptrs.value, mm_got.value, minimock.Diff(*mm_want_ptrs.value, mm_got.value))
			}

			if mm_want_ptrs.expiration != nil && !minimock.Equal(*mm_want_ptrs.expiration, mm_got.expiration) {
				mmSet.t.Errorf("RedisClientMock.Set got unexpected parameter expiration, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSet.SetMock.defaultExpectation.expectationOrigins.originExpiration, *mm_want_ptrs.expiration, mm_got.expiration, minimock.Diff(*mm_want_ptrs.expiration, mm_got.expiration))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("RedisClientMock.Set got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSet.SetMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the RedisClientMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(ctx, key, value, expiration)
	}
	mmSet.t.Fatalf("Unexpected call to RedisClientMock.Set. %v %v %v %v", ctx, key, value, expiration)
	return
}

// SetAfterCounter returns a count of finished RedisClientMock.Set invocations
func (mmSet *RedisClientMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of RedisClientMock.Set invocations
func (mmSet *RedisClientMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to RedisClientMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mRedisClientMockSet) Calls() []*RedisClientMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*RedisClientMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *RedisClientMock) MinimockSetDone() bool {
	if m.SetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SetMock.invocationsDone()
}

// MinimockSetInspect logs each unmet expectation
func (m *RedisClientMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisClientMock.Set at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSetCounter := mm_atomic.LoadUint64(&m.afterSetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && afterSetCounter < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to RedisClientMock.Set at\n%s", m.SetMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to RedisClientMock.Set at\n%s with params: %#v", m.SetMock.defaultExpectation.expectationOrigins.origin, *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && afterSetCounter < 1 {
		m.t.Errorf("Expected call to RedisClientMock.Set at\n%s", m.funcSetOrigin)
	}

	if !m.SetMock.invocationsDone() && afterSetCounter > 0 {
		m.t.Errorf("Expected %d calls to RedisClientMock.Set at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SetMock.expectedInvocations), m.SetMock.expectedInvocationsOrigin, afterSetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RedisClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockGetInspect()

			m.MinimockSetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RedisClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RedisClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockGetDone() &&
		m.MinimockSetDone()
}
