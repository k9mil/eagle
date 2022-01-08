package cmd

import "net/http/httptest"

type Answer struct {
	Items []struct {
		AnswerCount int    `json:"answer_count"`
		Score       int    `json:"score"`
		Link        string `json:"link"`
		Title       string `json:"title"`
	}
}

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Answer
	expectedError error
}
