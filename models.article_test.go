package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != 3 {
		t.Fail()
	}

}
