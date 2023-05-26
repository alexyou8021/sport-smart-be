package controllers

import "github.com/alexyou8021/sport-smart-be/entities"

func GetEvents() []entities.Event {
	return []entities.Event{
		{
			Id:       "1",
			Name:     "Practice 1",
			Location: "Location 1",
		},
		{
			Id:       "2",
			Name:     "Game 2",
			Location: "Location 2",
		},
		{
			Id:       "3",
			Name:     "Tournament 3",
			Location: "Location 3",
		},
	}
}

func GetMembers() []entities.Member {
	return []entities.Member{
		{
			Id:   "1",
			Name: "Coach",
		},
		{
			Id:   "2",
			Name: "Captain",
		},
		{
			Id:   "3",
			Name: "Player",
		},
	}
}
