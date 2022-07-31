// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"testing"
)

/*
func TestNewUser(t *testing.T) {
	name := "Testaroo"

}
*/

func TestNewQueueBadName(t *testing.T) {
	queue, err := newQueue("")
	if err == nil {
		t.Fatalf("empty queue should generate error")
	}
	log.Println(queue)

	/*
		    if user != "" || err == nil {
		        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
			}
	*/
}
