/*
Copyright © 2023 M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package main

/*
   https://medium.com/@sumitsagar_20050/mastering-go-channels-from-beginner-to-pro-9c1eaba0da9e
*/

import (
	"fmt"
	"sync"
)

// TSubscriptions is the main structure for our simple pub-sub system.
type TSubscriptions[T any] struct {
	mtx sync.RWMutex

	// The index keys are the topic names,
	// the value is respective subscriber's channel
	subscriptions map[string][]chan T
}

// NewSubscriptions creates a new `TSubscriptions` instance and initialises
// its internal subscriptions map.
//
// The function creates a new instance of the `TSubscriptions` struct and
// initializes its subscriptions map. The code uses the generic syntax `[T any]`
// to define the type parameter `T` as any type.
//
// The function starts by creating a new instance of the `TSubscriptions`
// struct. It then initializes the subscriptions map by creating a new empty
// map and assigning it to the subscriptions field.
//
// The code returns a pointer to the `TSubscriptions` struct so that it can
// be used in other functions.
func NewSubscriptions[T any]() *TSubscriptions[T] {
	return &TSubscriptions[T]{
		subscriptions: make(map[string][]chan T),
	}
} // NewSubscriptions()

// Subscribe allows a subscriber to get updates for a specific topic.
//
// This method takes `aTopic` name as an argument and returns a channel
// that can be used to receive updates for that topic.
//
// The method starts by acquiring a lock on the mutex (`mtx`) using the 'Lock'
// method. This ensures that no other goroutine can modify the subscriptions
// map while we are updating it.
//
// Next, the code creates a new channel (`ch`) and appends it to the list of
// subscribers for the given topic (`ts.subscriptions[aTopic]`). This is done
// by using the `append` function to add the new channel to the slice of
// subscribers.
//
// Finally, the code returns the new channel so that the subscriber can
// receive updates on it.
func (ts *TSubscriptions[T]) Subscribe(aTopic string) <-chan T {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	ch := make(chan T, 1)
	ts.subscriptions[aTopic] = append(ts.subscriptions[aTopic], ch)

	return ch
} // Subscribe()

// Publish sends the given value to all subscribers of a specific topic.
//
// The method starts by acquiring a read lock on the mutex (`mtx`) using
// the `RLock` method. This ensures that no other goroutine can modify the
// subscriptions map while we are reading from it.
//
// Next, the code checks if there are any subscribers for the given topic by
// looking up the topic in the subscriptions map. If the topic is not found,
// the method returns without doing anything.
//
// If the topic is found, however, the code iterates over the list of
// subscribers for the topic and checks if any of them matches the
// unsubscribe channel (`aSubCh`). If a match is found, the code sends
// the given value to the subscriber's channel using the <- operator.
//
// After sending the value, the code releases the lock and returns.
func (ts *TSubscriptions[T]) Publish(aTopic string, value T) {
	ts.mtx.RLock()
	defer ts.mtx.RUnlock()

	// traverse the whole existing list:
	for _, subscriber := range ts.subscriptions[aTopic] {
		subscriber <- value
	}
} // Publish()

// Unsubscribe removes a specific subscriber from a topic and closes its
// channel.
//
// The method starts by acquiring a read lock on the mutex (`mtx`) using
// its `RLock` method.
// This ensures that no other goroutine can modify the subscriptions map
// while we are reading from it.
//
// Next, the method checks if there are any subscribers for the given
// `aTopic` by looking up the topic in the subscriptions map.
// If the topic is not found, the method returns without doing anything.
//
// If, however the topic is found, the code iterates over the list of
// subscribers for the topic and checks if any of them matches the
// unsubscribe channel (`aSubCh`).
// If a match is found, the method closes the subscriber's channel and
// removes the subscriber from the list.
//
// After the subscriber is removed, the method releases the lock and returns.
func (ts *TSubscriptions[T]) Unsubscribe(aTopic string, aSubCh <-chan T) {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	subscribers, found := ts.subscriptions[aTopic]
	if !found {
		return
	}

	for i, subscriber := range subscribers {
		if subscriber == aSubCh {
			close(subscriber)
			ts.subscriptions[aTopic] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
} // Unsubscribe()

// Main function demonstrating the publisher-subscriber functionality.
//
// The `main()` function demonstrates how to use the `TSubscriptions` type
// from the `Go` programming language. Specifically, it shows how to create
// a new instance of `TSubscriptions`, subscribe to a topic, publish a value
// to that topic, and unsubscribe from the topic.
//
// The function starts by creating a new instance of `TSubscriptions` and
// assigning it to a variable named `ts`. Then, it subscribes to a topic
// named "topic1" by calling the `Subscribe()` method on `ts` and assigning
// the result to a variable named `subscriber`.
//
// Next, the function creates a new goroutine (a lightweight thread) by
// calling the `go` function. Inside the goroutine, it calls the `Publish()`
// method on `ts`, passing "topic1" as the topic and the value 42 as the
// value to publish.
//
// Finally, the function waits for the subscriber to receive a value and
// stores it in a variable named `value`. Then, it prints the received `value`
// to the console.
//
// After printing the value, the code unsubscribes from the topic by calling
// the `Unsubscribe()` method on `ts` and passing "topic1" as the topic and
// `subscriber` as the unsubscribe channel.
func main() {
	ts := NewSubscriptions[int]()

	subscriber := ts.Subscribe("topic1")

	go func() {
		ts.Publish("topic1", 42)
	}()

	value := <-subscriber
	fmt.Println("Received value:", value) // Expected: Received value: 42
	ts.Unsubscribe("topic1", subscriber)
} // main()

//_EoF_
