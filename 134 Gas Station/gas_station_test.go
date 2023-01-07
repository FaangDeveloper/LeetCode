package main

import (
	"testing"
)

// https://leetcode.com/problems/gas-station/

func TestCanCompleteCircuit(t *testing.T) {
	firstResult := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	firstExpected := 3
	if firstResult != firstExpected {
		t.Errorf("result: %v, expected: %v", firstResult, firstExpected)
	}
	secondExample := canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	secondExpected := -1
	if secondExpected != secondExample {
		t.Errorf("result: %v, expected: %v", secondExample, secondExpected)
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank := 0
	startingTank := 0
	startingStation := 0
	for i := range gas {
		// if the total tank is positive or zero
		// then we can drive a circular route
		totalTank += gas[i] - cost[i]
		// if starting tank is negative
		// then we must try to start on next station
		startingTank += gas[i] - cost[i]
		if startingTank < 0 {
			startingStation = i + 1
			startingTank = 0
		}
	}
	if totalTank >= 0 {
		return startingStation
	}
	return -1
}

// Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// Output: 3

// Input: gas = [2,3,4], cost = [3,4,3]
// Output: -1
