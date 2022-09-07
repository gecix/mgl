package list

import (
	"testing"
)

func TestWrappedList(t *testing.T) {
	valuesComposite := []interface{}{make(chan int), []byte("test"), &struct{ n int }{n: 69}}
	ints := []int{1, 2, 3, 4}
	floats := []float32{1, 2, 3, 4}
	strings := []string{"kuy", "hee", "tad"}

	t.Run("SafeList", func(t *testing.T) {
		testSafeList(t, valuesComposite)
		testSafeList(t, ints)
		testSafeList(t, floats)
		testSafeList(t, strings)
	})

	t.Run("SetList", testSetList)
}

func testSafeList[T any](t *testing.T, values []T) {
	// Those list types wrapped in SafeList[T, BasicList[T]] also implement BasicList[T]
	basicStack := NewStack[T]()
	safeStack := WrapSafeList[T](basicStack)
	basicQueue := NewQueue[T]()
	safeQueue := WrapSafeList[T](basicQueue)
	anotherSafeQueue := NewSafeQueue[T]()
	tests := []BasicList[T]{safeStack, safeQueue, anotherSafeQueue}

	for _, l := range tests {
		testBasicList(t, values, l)
	}
}

func TestLists(t *testing.T) {
	// We can also create a composite key with non-comparable types
	valuesComposite := []interface{}{make(chan int), []byte("test"), &struct{ n int }{n: 69}}
	qComposite := NewQueue[interface{}]()
	sComposite := NewStack[interface{}]()
	testBasicList[interface{}](t, valuesComposite, qComposite)
	testBasicList[interface{}](t, valuesComposite, sComposite)

}

func testBasicList[T any](t *testing.T, values []T, basicList BasicList[T]) {
	for _, value := range values {
		basicList.Push(value)
		// t.Logf("BasicList.Push - value: %v, type: %s\n", value, reflect.TypeOf(value))
	}
	if listLen := basicList.Len(); listLen != len(values) {
		t.Fatal("unexpected basicList len")
	}
	for i, listLen := 0, basicList.Len(); i < listLen; i++ {
		p := basicList.Pop()
		if p == nil {
			t.Fatal("unexpected nil _here_ from BasicList.Pop")
		}
		// v := *p
		// t.Logf("BasicList.Pop - value: %v, type: %s\n", v, reflect.TypeOf(v))
	}
}
