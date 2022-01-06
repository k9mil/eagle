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
			name: "api-call-test",
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
		{
			name: "decode-json-test",
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

	t.Run(tests[0].name, func(t *testing.T) {
		defer tests[0].server.Close()
		resp, err := apiCall(tests[0].server.URL)

		if !reflect.DeepEqual(&resp, tests[0].response) {
			t.Errorf("DEEP EQUAL FAILED: Expected: %v, got: %v\n", tests[0].response, resp)
		}

		if !errors.Is(err, tests[0].expectedError) {
			t.Errorf("EXPECTED ERROR FAILED: Expected: %v got: %v\n", tests[0].expectedError, err)
		}
	})

	t.Run(tests[1].name, func(t *testing.T) {
		defer tests[0].server.Close()
		byteResponse := []byte(`{"items":[{"answer_count":5,"score":10,"link":"https://stackoverflow.com/questions/32531854/how-to-initialize-nested-structure-array-in-golang","title":"How to initialize nested structure array in golang? [duplicate]"}]}`)
		resp, err := decodeJSON(byteResponse)

		if !reflect.DeepEqual(&resp, tests[1].response) {
			t.Errorf("DEEP EQUAL FAILED: Expected: %v, got: %v\n", tests[1].response, resp)
		}

		if !errors.Is(err, tests[1].expectedError) {
			t.Errorf("EXPECTED ERROR FAILED: Expected: %v got: %v\n", tests[1].expectedError, err)
		}
	})
}

func TestStringInSlice(t *testing.T) {
	sampleSort := "activity"
	sampleList := []string{"votes", "activity", "creation", "relevance"}

	got := stringInSlice(sampleSort, sampleList)

	if got != nil {
		t.Errorf("got %s want %s given: %s", got, "nil", sampleSort)
	}
}

func TestFormatLink(t *testing.T) {
	testURL := "https://stackoverflow.com/questions/7560832/how-to-center-a-button-within-a-div"
	got := formatLink(testURL)
	want := "https://stackoverflow.com/q/7560832"

	if got != want {
		t.Errorf("got %s want %s given: %s", got, want, testURL)
	}
}
