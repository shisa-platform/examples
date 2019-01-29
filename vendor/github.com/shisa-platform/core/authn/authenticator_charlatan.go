// generated by "charlatan -output=./authenticator_charlatan.go Authenticator".  DO NOT EDIT.

package authn

import "reflect"
import "github.com/ansel1/merry"
import "github.com/shisa-platform/core/context"
import "github.com/shisa-platform/core/httpx"
import "github.com/shisa-platform/core/models"

// AuthenticatorAuthenticateInvocation represents a single call of FakeAuthenticator.Authenticate
type AuthenticatorAuthenticateInvocation struct {
	Parameters struct {
		Ident1 context.Context
		Ident2 *httpx.Request
	}
	Results struct {
		Ident3 models.User
		Ident4 merry.Error
	}
}

// NewAuthenticatorAuthenticateInvocation creates a new instance of AuthenticatorAuthenticateInvocation
func NewAuthenticatorAuthenticateInvocation(ident1 context.Context, ident2 *httpx.Request, ident3 models.User, ident4 merry.Error) *AuthenticatorAuthenticateInvocation {
	invocation := new(AuthenticatorAuthenticateInvocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	return invocation
}

// AuthenticatorChallengeInvocation represents a single call of FakeAuthenticator.Challenge
type AuthenticatorChallengeInvocation struct {
	Results struct {
		Ident1 string
	}
}

// AuthenticatorTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type AuthenticatorTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeAuthenticator is a mock implementation of Authenticator for testing.
Use it in your tests as in this example:

	package example

	func TestWithAuthenticator(t *testing.T) {
		f := &authn.FakeAuthenticator{
			AuthenticateHook: func(ident1 context.Context, ident2 *httpx.Request) (ident3 models.User, ident4 merry.Error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeAuthenticate ...
		f.AssertAuthenticateCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeAuthenticate.
*/
type FakeAuthenticator struct {
	AuthenticateHook func(context.Context, *httpx.Request) (models.User, merry.Error)
	ChallengeHook    func() string

	AuthenticateCalls []*AuthenticatorAuthenticateInvocation
	ChallengeCalls    []*AuthenticatorChallengeInvocation
}

// NewFakeAuthenticatorDefaultPanic returns an instance of FakeAuthenticator with all hooks configured to panic
func NewFakeAuthenticatorDefaultPanic() *FakeAuthenticator {
	return &FakeAuthenticator{
		AuthenticateHook: func(context.Context, *httpx.Request) (ident3 models.User, ident4 merry.Error) {
			panic("Unexpected call to Authenticator.Authenticate")
		},
		ChallengeHook: func() (ident1 string) {
			panic("Unexpected call to Authenticator.Challenge")
		},
	}
}

// NewFakeAuthenticatorDefaultFatal returns an instance of FakeAuthenticator with all hooks configured to call t.Fatal
func NewFakeAuthenticatorDefaultFatal(t AuthenticatorTestingT) *FakeAuthenticator {
	return &FakeAuthenticator{
		AuthenticateHook: func(context.Context, *httpx.Request) (ident3 models.User, ident4 merry.Error) {
			t.Fatal("Unexpected call to Authenticator.Authenticate")
			return
		},
		ChallengeHook: func() (ident1 string) {
			t.Fatal("Unexpected call to Authenticator.Challenge")
			return
		},
	}
}

// NewFakeAuthenticatorDefaultError returns an instance of FakeAuthenticator with all hooks configured to call t.Error
func NewFakeAuthenticatorDefaultError(t AuthenticatorTestingT) *FakeAuthenticator {
	return &FakeAuthenticator{
		AuthenticateHook: func(context.Context, *httpx.Request) (ident3 models.User, ident4 merry.Error) {
			t.Error("Unexpected call to Authenticator.Authenticate")
			return
		},
		ChallengeHook: func() (ident1 string) {
			t.Error("Unexpected call to Authenticator.Challenge")
			return
		},
	}
}

func (f *FakeAuthenticator) Reset() {
	f.AuthenticateCalls = []*AuthenticatorAuthenticateInvocation{}
	f.ChallengeCalls = []*AuthenticatorChallengeInvocation{}
}

func (_f1 *FakeAuthenticator) Authenticate(ident1 context.Context, ident2 *httpx.Request) (ident3 models.User, ident4 merry.Error) {
	if _f1.AuthenticateHook == nil {
		panic("Authenticator.Authenticate() called but FakeAuthenticator.AuthenticateHook is nil")
	}

	invocation := new(AuthenticatorAuthenticateInvocation)
	_f1.AuthenticateCalls = append(_f1.AuthenticateCalls, invocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	ident3, ident4 = _f1.AuthenticateHook(ident1, ident2)

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	return
}

// SetAuthenticateStub configures Authenticator.Authenticate to always return the given values
func (_f2 *FakeAuthenticator) SetAuthenticateStub(ident3 models.User, ident4 merry.Error) {
	_f2.AuthenticateHook = func(context.Context, *httpx.Request) (models.User, merry.Error) {
		return ident3, ident4
	}
}

// SetAuthenticateInvocation configures Authenticator.Authenticate to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (_f3 *FakeAuthenticator) SetAuthenticateInvocation(calls_f4 []*AuthenticatorAuthenticateInvocation, fallback_f5 func() (models.User, merry.Error)) {
	_f3.AuthenticateHook = func(ident1 context.Context, ident2 *httpx.Request) (ident3 models.User, ident4 merry.Error) {
		for _, call := range calls_f4 {
			if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
				ident3 = call.Results.Ident3
				ident4 = call.Results.Ident4

				return
			}
		}

		return fallback_f5()
	}
}

// AuthenticateCalled returns true if FakeAuthenticator.Authenticate was called
func (f *FakeAuthenticator) AuthenticateCalled() bool {
	return len(f.AuthenticateCalls) != 0
}

// AssertAuthenticateCalled calls t.Error if FakeAuthenticator.Authenticate was not called
func (f *FakeAuthenticator) AssertAuthenticateCalled(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) == 0 {
		t.Error("FakeAuthenticator.Authenticate not called, expected at least one")
	}
}

// AuthenticateNotCalled returns true if FakeAuthenticator.Authenticate was not called
func (f *FakeAuthenticator) AuthenticateNotCalled() bool {
	return len(f.AuthenticateCalls) == 0
}

// AssertAuthenticateNotCalled calls t.Error if FakeAuthenticator.Authenticate was called
func (f *FakeAuthenticator) AssertAuthenticateNotCalled(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) != 0 {
		t.Error("FakeAuthenticator.Authenticate called, expected none")
	}
}

// AuthenticateCalledOnce returns true if FakeAuthenticator.Authenticate was called exactly once
func (f *FakeAuthenticator) AuthenticateCalledOnce() bool {
	return len(f.AuthenticateCalls) == 1
}

// AssertAuthenticateCalledOnce calls t.Error if FakeAuthenticator.Authenticate was not called exactly once
func (f *FakeAuthenticator) AssertAuthenticateCalledOnce(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) != 1 {
		t.Errorf("FakeAuthenticator.Authenticate called %d times, expected 1", len(f.AuthenticateCalls))
	}
}

// AuthenticateCalledN returns true if FakeAuthenticator.Authenticate was called at least n times
func (f *FakeAuthenticator) AuthenticateCalledN(n int) bool {
	return len(f.AuthenticateCalls) >= n
}

// AssertAuthenticateCalledN calls t.Error if FakeAuthenticator.Authenticate was called less than n times
func (f *FakeAuthenticator) AssertAuthenticateCalledN(t AuthenticatorTestingT, n int) {
	t.Helper()
	if len(f.AuthenticateCalls) < n {
		t.Errorf("FakeAuthenticator.Authenticate called %d times, expected >= %d", len(f.AuthenticateCalls), n)
	}
}

// AuthenticateCalledWith returns true if FakeAuthenticator.Authenticate was called with the given values
func (_f6 *FakeAuthenticator) AuthenticateCalledWith(ident1 context.Context, ident2 *httpx.Request) (found bool) {
	for _, call := range _f6.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	return
}

// AssertAuthenticateCalledWith calls t.Error if FakeAuthenticator.Authenticate was not called with the given values
func (_f7 *FakeAuthenticator) AssertAuthenticateCalledWith(t AuthenticatorTestingT, ident1 context.Context, ident2 *httpx.Request) {
	t.Helper()
	var found bool
	for _, call := range _f7.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeAuthenticator.Authenticate not called with expected parameters")
	}
}

// AuthenticateCalledOnceWith returns true if FakeAuthenticator.Authenticate was called exactly once with the given values
func (_f8 *FakeAuthenticator) AuthenticateCalledOnceWith(ident1 context.Context, ident2 *httpx.Request) bool {
	var count int
	for _, call := range _f8.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	return count == 1
}

// AssertAuthenticateCalledOnceWith calls t.Error if FakeAuthenticator.Authenticate was not called exactly once with the given values
func (_f9 *FakeAuthenticator) AssertAuthenticateCalledOnceWith(t AuthenticatorTestingT, ident1 context.Context, ident2 *httpx.Request) {
	t.Helper()
	var count int
	for _, call := range _f9.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeAuthenticator.Authenticate called %d times with expected parameters, expected one", count)
	}
}

// AuthenticateResultsForCall returns the result values for the first call to FakeAuthenticator.Authenticate with the given values
func (_f10 *FakeAuthenticator) AuthenticateResultsForCall(ident1 context.Context, ident2 *httpx.Request) (ident3 models.User, ident4 merry.Error, found bool) {
	for _, call := range _f10.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			ident3 = call.Results.Ident3
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}

func (_f11 *FakeAuthenticator) Challenge() (ident1 string) {
	if _f11.ChallengeHook == nil {
		panic("Authenticator.Challenge() called but FakeAuthenticator.ChallengeHook is nil")
	}

	invocation := new(AuthenticatorChallengeInvocation)
	_f11.ChallengeCalls = append(_f11.ChallengeCalls, invocation)

	ident1 = _f11.ChallengeHook()

	invocation.Results.Ident1 = ident1

	return
}

// SetChallengeStub configures Authenticator.Challenge to always return the given values
func (_f12 *FakeAuthenticator) SetChallengeStub(ident1 string) {
	_f12.ChallengeHook = func() string {
		return ident1
	}
}

// ChallengeCalled returns true if FakeAuthenticator.Challenge was called
func (f *FakeAuthenticator) ChallengeCalled() bool {
	return len(f.ChallengeCalls) != 0
}

// AssertChallengeCalled calls t.Error if FakeAuthenticator.Challenge was not called
func (f *FakeAuthenticator) AssertChallengeCalled(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.ChallengeCalls) == 0 {
		t.Error("FakeAuthenticator.Challenge not called, expected at least one")
	}
}

// ChallengeNotCalled returns true if FakeAuthenticator.Challenge was not called
func (f *FakeAuthenticator) ChallengeNotCalled() bool {
	return len(f.ChallengeCalls) == 0
}

// AssertChallengeNotCalled calls t.Error if FakeAuthenticator.Challenge was called
func (f *FakeAuthenticator) AssertChallengeNotCalled(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.ChallengeCalls) != 0 {
		t.Error("FakeAuthenticator.Challenge called, expected none")
	}
}

// ChallengeCalledOnce returns true if FakeAuthenticator.Challenge was called exactly once
func (f *FakeAuthenticator) ChallengeCalledOnce() bool {
	return len(f.ChallengeCalls) == 1
}

// AssertChallengeCalledOnce calls t.Error if FakeAuthenticator.Challenge was not called exactly once
func (f *FakeAuthenticator) AssertChallengeCalledOnce(t AuthenticatorTestingT) {
	t.Helper()
	if len(f.ChallengeCalls) != 1 {
		t.Errorf("FakeAuthenticator.Challenge called %d times, expected 1", len(f.ChallengeCalls))
	}
}

// ChallengeCalledN returns true if FakeAuthenticator.Challenge was called at least n times
func (f *FakeAuthenticator) ChallengeCalledN(n int) bool {
	return len(f.ChallengeCalls) >= n
}

// AssertChallengeCalledN calls t.Error if FakeAuthenticator.Challenge was called less than n times
func (f *FakeAuthenticator) AssertChallengeCalledN(t AuthenticatorTestingT, n int) {
	t.Helper()
	if len(f.ChallengeCalls) < n {
		t.Errorf("FakeAuthenticator.Challenge called %d times, expected >= %d", len(f.ChallengeCalls), n)
	}
}
