package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var (
	Sort    string
	Title   string
	Results string
	rawData Answer
)

func init() {
	searchCmd.Flags().StringVarP(&Title, "title", "t", "", "The title of the query. (required)")
	searchCmd.Flags().StringVarP(&Sort, "sort", "s", "votes", "The sort method to be used. (optional)")
	searchCmd.Flags().StringVarP(&Results, "results", "r", "20", "The number of posts to be displayed. (optional)")
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Required command to search for your question.",
	Args: func(searchCmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a title")
		}

		ListOfOptions := []string{"votes", "activity", "creation", "relevance"}

		if len(args) > 1 {
			err := stringInSlice(args[1], ListOfOptions)

			if err != nil {
				return fmt.Errorf("searchCmd: An errror occured using stringInSlice(): %w", err)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		Title = url.QueryEscape(args[0])
		if len(args) < 2 {
			search(Title, "votes", "20")
		} else if len(args) < 3 {
			search(Title, args[1], "20")
		} else {
			search(Title, args[1], args[2])
		}
	},
}

func stringInSlice(sort string, ListOfOptions []string) error {
	for _, val := range ListOfOptions {
		if val == sort {
			return nil
		}
	}

	return errors.New("sort method not found")
}

func search(title, sort, results string) error {
	url := fmt.Sprintf("https://api.stackexchange.com/2.3/search?order=desc&sort=%s&intitle=%s&site=stackoverflow&pagesize=%s", sort, title, results)
	apiReturn, err := apiCall(url, rawData)

	if err != nil {
		return fmt.Errorf("search: An error occured: %w", err)
	}

	broadcastAnswer(apiReturn)
	return nil
}

func apiCall(url string, rawData Answer) (Answer, error) {
	resp, err := http.Get(url)

	if err != nil {
		return rawData, fmt.Errorf("apiCall: An error occured with the GET request: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return rawData, fmt.Errorf("apiCall: An error occured with ReadAll(): %w", err)
	}

	return decodeJSON(body, rawData)
}

func decodeJSON(resp []byte, rawData Answer) (Answer, error) {
	err := json.Unmarshal(resp, &rawData)

	if err != nil {
		return rawData, fmt.Errorf("decodeJSON: An error occured with unmarshalling the data: %w", err)
	}

	return rawData, nil
}

func formatLink(Link string) string {
	standardURL := "https://stackoverflow.com/q/"
	re := regexp.MustCompile("[0-9]+")
	questionID := re.FindAllString(Link, -1)

	return standardURL + questionID[0]
}

func broadcastAnswer(a Answer) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Title", "Answers", "Upvotes", "Link"})

	for _, item := range a.Items {
		t.AppendRows([]table.Row{{html.UnescapeString(item.Title), item.AnswerCount, item.Score, formatLink(item.Link)}})
	}

	t.Render()
}
