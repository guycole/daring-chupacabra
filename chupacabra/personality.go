// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

/*
type userType struct {
	name string
}

const maxUsers = 10

type userArrayType [maxUsers]*userType

var userArray userArrayType

func newUser(name string) (*userType, error) {
	if len(name) < 1 {
		return nil, errors.New("bad name")
	}

	result := userType{name: name}
	return &result, nil
}

func userAdd(name string) (int, error) {
	for ndx := 0; ndx < maxUsers; ndx++ {
		if (userArray[ndx] != nil) && (userArray[ndx].name == name) {
			return -1, errors.New("duplicate name")
		}
	}

	candidate, err := newUser(name)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	for ndx := 0; ndx < maxUsers; ndx++ {
		if userArray[ndx] == nil {
			userArray[ndx] = candidate
			return ndx, nil
		}
	}

	return -1, errors.New("userArray full")
}

func userDelete(name string) int {
	for ndx := 0; ndx < maxUsers; ndx++ {
		if userArray[ndx] == nil {
			userArray[ndx] = nil
			return ndx
		}
	}

	return -1
}
*/
