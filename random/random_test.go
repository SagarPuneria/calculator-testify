package rdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// randomMock is a mocked object that implements an random interface
// that describes an object that the code I am testing relies on.
type randomMock struct {
	mock.Mock
}

// getRandom is a method on randomMock that implements random interface
// and returns what the Mock object tells it to.
func (m *randomMock) getRandom(res, min int) (int, error) {
	args := m.Called(res, min)
	return args.Int(0), args.Error(1)
}

// Test_random_mock is used to test our rndMk object to
// make assertions for getRandom function code we are mocking.
func Test_random_mock(t *testing.T) {
	rndMk := new(randomMock) //OR &randomMock{}
	rndMk.On("getRandom", 9, 1).Return(3, nil)
	c := &Calc{Rnd: rndMk}
	min, max := 1, 10
	randomValue, err := c.Random(&min, &max)
	rndMk.AssertExpectations(t)
	assert.Nil(t, err, "Error is NOT nil")
	assert.Equal(t, 3, randomValue)
}

// Random will take two formal parameters
// should return any random value which lies within the range of two formal parameters
func Test_given_two_number_function_Random_should_return_random_number(t *testing.T) {
	r := &Calc{Rnd: &Calc{}}
	min, max := 1, 9
	_, err := r.Random(&min, &max)
	assert.Nil(t, err, "Error is NOT nil")
}
