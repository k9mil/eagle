package cmd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Answer
	expectedError error
}

func TestSearch(t *testing.T) {
	tests := []Tests{
		{
			name: "basic-response-test",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"items":[{"answer_count":5,"score":10,"link":"https://stackoverflow.com/questions/32531854/how-to-initialize-nested-structure-array-in-golang","title":"How to initialize nested structure array in golang? [duplicate]"}]}`))
			})),
			response: &Answer{
				Items: []struct {
					AnswerCount int    `json:"answer_count"`
					Score       int    `json:"score"`
					Link        string `json:"link"`
					Title       string `json:"title"`
				}{
					{
						AnswerCount: 5,
						Score:       10,
						Link:        "https://stackoverflow.com/questions/32531854/how-to-initialize-nested-structure-array-in-golang",
						Title:       "How to initialize nested structure array in golang? [duplicate]",
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := apiCall(test.server.URL)

			if !reflect.DeepEqual(&resp, test.response) {
				t.Errorf("DEEP EQUAL FAILED: Expected: %v, got: %v\n", test.response, resp)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("EXPECTED ERROR FAILED: Expected: %v got: %v\n", test.expectedError, err)
			}
		})
	}
}
