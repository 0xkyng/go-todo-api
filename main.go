package main

type todo struct {
	ID          string
	Item        string
	Completed   bool
}

var todos = []todo {
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "1", Item: "Watch Footbal", Completed: false},
}