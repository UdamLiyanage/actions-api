package main

import "go.mongodb.org/mongo-driver/mongo"

func checkError(err error) bool {
	if err != nil {
		return true
	}
	if err == mongo.ErrNoDocuments {
		return true
	}
	return false
}
