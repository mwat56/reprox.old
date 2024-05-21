/*
Copyright © 2023 M.Watermann, 10247 Berlin, Germany

	    All rights reserved
	EMail : <support@mwat.de>
*/
package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSubscriptions(t *testing.T) {
	var (
		//T  = any
		T1 = "string"
	)
	tests := []struct {
		name string
		want *TSubscriptions[T]
	}{
		{"1", T1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubscriptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubscriptions() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestNewSubscriptions()

func TestNewSubscriptionsXX[T any](t *testing.t) {
	// Arrange
	var expectedSubscriptions = make(map[string][]chan T)

	// Act
	var actualSubscriptions = NewSubscriptions[T]()

	// Assert
	if assert.Equal(t, expectedSubscriptions, actualSubscriptions.subscriptions) {
		t.Log("Success: NewSubscriptions initialized the subscriptions map correctly")
	} else {
		t.Error("Failure: NewSubscriptions did not initialize the subscriptions map correctly")
	}
}
