// generated by "charlatan -output=./consulresolver_charlatan.go consulResolver".  DO NOT EDIT.

package sd

import consul "github.com/hashicorp/consul/api"
import "reflect"

// consulResolverServiceInvocation represents a single call of FakeconsulResolver.Service
type consulResolverServiceInvocation struct {
	Parameters struct {
		Service     string
		Tag         string
		PassingOnly bool
		Q           *consul.QueryOptions
	}
	Results struct {
		Ident1 []*consul.ServiceEntry
		Ident2 *consul.QueryMeta
		Ident3 error
	}
}

// NewconsulResolverServiceInvocation creates a new instance of consulResolverServiceInvocation
func NewconsulResolverServiceInvocation(service string, tag string, passingOnly bool, q *consul.QueryOptions, ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) *consulResolverServiceInvocation {
	invocation := new(consulResolverServiceInvocation)

	invocation.Parameters.Service = service
	invocation.Parameters.Tag = tag
	invocation.Parameters.PassingOnly = passingOnly
	invocation.Parameters.Q = q

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2
	invocation.Results.Ident3 = ident3

	return invocation
}

// consulResolverTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type consulResolverTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeconsulResolver is a mock implementation of consulResolver for testing.
Use it in your tests as in this example:

	package example

	func TestWithconsulResolver(t *testing.T) {
		f := &sd.FakeconsulResolver{
			ServiceHook: func(service string, tag string, passingOnly bool, q *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeService ...
		f.AssertServiceCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeService.
*/
type FakeconsulResolver struct {
	ServiceHook func(string, string, bool, *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error)

	ServiceCalls []*consulResolverServiceInvocation
}

// NewFakeconsulResolverDefaultPanic returns an instance of FakeconsulResolver with all hooks configured to panic
func NewFakeconsulResolverDefaultPanic() *FakeconsulResolver {
	return &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
			panic("Unexpected call to consulResolver.Service")
		},
	}
}

// NewFakeconsulResolverDefaultFatal returns an instance of FakeconsulResolver with all hooks configured to call t.Fatal
func NewFakeconsulResolverDefaultFatal(t consulResolverTestingT) *FakeconsulResolver {
	return &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
			t.Fatal("Unexpected call to consulResolver.Service")
			return
		},
	}
}

// NewFakeconsulResolverDefaultError returns an instance of FakeconsulResolver with all hooks configured to call t.Error
func NewFakeconsulResolverDefaultError(t consulResolverTestingT) *FakeconsulResolver {
	return &FakeconsulResolver{
		ServiceHook: func(string, string, bool, *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
			t.Error("Unexpected call to consulResolver.Service")
			return
		},
	}
}

func (f *FakeconsulResolver) Reset() {
	f.ServiceCalls = []*consulResolverServiceInvocation{}
}

func (_f1 *FakeconsulResolver) Service(service string, tag string, passingOnly bool, q *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
	if _f1.ServiceHook == nil {
		panic("consulResolver.Service() called but FakeconsulResolver.ServiceHook is nil")
	}

	invocation := new(consulResolverServiceInvocation)
	_f1.ServiceCalls = append(_f1.ServiceCalls, invocation)

	invocation.Parameters.Service = service
	invocation.Parameters.Tag = tag
	invocation.Parameters.PassingOnly = passingOnly
	invocation.Parameters.Q = q

	ident1, ident2, ident3 = _f1.ServiceHook(service, tag, passingOnly, q)

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2
	invocation.Results.Ident3 = ident3

	return
}

// SetServiceStub configures consulResolver.Service to always return the given values
func (_f2 *FakeconsulResolver) SetServiceStub(ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
	_f2.ServiceHook = func(string, string, bool, *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
		return ident1, ident2, ident3
	}
}

// SetServiceInvocation configures consulResolver.Service to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (_f3 *FakeconsulResolver) SetServiceInvocation(calls_f4 []*consulResolverServiceInvocation, fallback_f5 func() ([]*consul.ServiceEntry, *consul.QueryMeta, error)) {
	_f3.ServiceHook = func(service string, tag string, passingOnly bool, q *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error) {
		for _, call := range calls_f4 {
			if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
				ident1 = call.Results.Ident1
				ident2 = call.Results.Ident2
				ident3 = call.Results.Ident3

				return
			}
		}

		return fallback_f5()
	}
}

// ServiceCalled returns true if FakeconsulResolver.Service was called
func (f *FakeconsulResolver) ServiceCalled() bool {
	return len(f.ServiceCalls) != 0
}

// AssertServiceCalled calls t.Error if FakeconsulResolver.Service was not called
func (f *FakeconsulResolver) AssertServiceCalled(t consulResolverTestingT) {
	t.Helper()
	if len(f.ServiceCalls) == 0 {
		t.Error("FakeconsulResolver.Service not called, expected at least one")
	}
}

// ServiceNotCalled returns true if FakeconsulResolver.Service was not called
func (f *FakeconsulResolver) ServiceNotCalled() bool {
	return len(f.ServiceCalls) == 0
}

// AssertServiceNotCalled calls t.Error if FakeconsulResolver.Service was called
func (f *FakeconsulResolver) AssertServiceNotCalled(t consulResolverTestingT) {
	t.Helper()
	if len(f.ServiceCalls) != 0 {
		t.Error("FakeconsulResolver.Service called, expected none")
	}
}

// ServiceCalledOnce returns true if FakeconsulResolver.Service was called exactly once
func (f *FakeconsulResolver) ServiceCalledOnce() bool {
	return len(f.ServiceCalls) == 1
}

// AssertServiceCalledOnce calls t.Error if FakeconsulResolver.Service was not called exactly once
func (f *FakeconsulResolver) AssertServiceCalledOnce(t consulResolverTestingT) {
	t.Helper()
	if len(f.ServiceCalls) != 1 {
		t.Errorf("FakeconsulResolver.Service called %d times, expected 1", len(f.ServiceCalls))
	}
}

// ServiceCalledN returns true if FakeconsulResolver.Service was called at least n times
func (f *FakeconsulResolver) ServiceCalledN(n int) bool {
	return len(f.ServiceCalls) >= n
}

// AssertServiceCalledN calls t.Error if FakeconsulResolver.Service was called less than n times
func (f *FakeconsulResolver) AssertServiceCalledN(t consulResolverTestingT, n int) {
	t.Helper()
	if len(f.ServiceCalls) < n {
		t.Errorf("FakeconsulResolver.Service called %d times, expected >= %d", len(f.ServiceCalls), n)
	}
}

// ServiceCalledWith returns true if FakeconsulResolver.Service was called with the given values
func (_f6 *FakeconsulResolver) ServiceCalledWith(service string, tag string, passingOnly bool, q *consul.QueryOptions) (found bool) {
	for _, call := range _f6.ServiceCalls {
		if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
			found = true
			break
		}
	}

	return
}

// AssertServiceCalledWith calls t.Error if FakeconsulResolver.Service was not called with the given values
func (_f7 *FakeconsulResolver) AssertServiceCalledWith(t consulResolverTestingT, service string, tag string, passingOnly bool, q *consul.QueryOptions) {
	t.Helper()
	var found bool
	for _, call := range _f7.ServiceCalls {
		if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeconsulResolver.Service not called with expected parameters")
	}
}

// ServiceCalledOnceWith returns true if FakeconsulResolver.Service was called exactly once with the given values
func (_f8 *FakeconsulResolver) ServiceCalledOnceWith(service string, tag string, passingOnly bool, q *consul.QueryOptions) bool {
	var count int
	for _, call := range _f8.ServiceCalls {
		if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
			count++
		}
	}

	return count == 1
}

// AssertServiceCalledOnceWith calls t.Error if FakeconsulResolver.Service was not called exactly once with the given values
func (_f9 *FakeconsulResolver) AssertServiceCalledOnceWith(t consulResolverTestingT, service string, tag string, passingOnly bool, q *consul.QueryOptions) {
	t.Helper()
	var count int
	for _, call := range _f9.ServiceCalls {
		if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeconsulResolver.Service called %d times with expected parameters, expected one", count)
	}
}

// ServiceResultsForCall returns the result values for the first call to FakeconsulResolver.Service with the given values
func (_f10 *FakeconsulResolver) ServiceResultsForCall(service string, tag string, passingOnly bool, q *consul.QueryOptions) (ident1 []*consul.ServiceEntry, ident2 *consul.QueryMeta, ident3 error, found bool) {
	for _, call := range _f10.ServiceCalls {
		if reflect.DeepEqual(call.Parameters.Service, service) && reflect.DeepEqual(call.Parameters.Tag, tag) && reflect.DeepEqual(call.Parameters.PassingOnly, passingOnly) && reflect.DeepEqual(call.Parameters.Q, q) {
			ident1 = call.Results.Ident1
			ident2 = call.Results.Ident2
			ident3 = call.Results.Ident3
			found = true
			break
		}
	}

	return
}