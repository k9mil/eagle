package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

type Answer struct {
	Items []struct {
		AnswerCount int    `json:"answer_count"`
		Score       int    `json:"score"`
		Link        string `json:"link"`
		Title       string `json:"title"`
	}
}

var (
	Sort  string
	Title string
)

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
				log.Fatalln(err)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		Title = url.QueryEscape(args[0])
		if len(args) > 1 {
			search(Title, args[1])
		} else {
			search(Title, "votes")
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

func init() {
	searchCmd.Flags().StringVarP(&Title, "title", "t", "", "The title of the query. (required)")
	searchCmd.Flags().StringVarP(&Sort, "sort", "s", "", "The sort method to be used. (optional, default: votes)")
	rootCmd.AddCommand(searchCmd)
}

func search(title, sort string) {
	url := fmt.Sprintf("https://api.stackexchange.com/2.3/search?order=desc&sort=%s&intitle=%s&site=stackoverflow", sort, title)
	apiReturn, err := apiCall(url)

	if err != nil {
		log.Fatalln(err)
	}

	broadcastAnswer(apiReturn)
}

func apiCall(url string) (Answer, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return decodeJSON(body)
}

func decodeJSON(resp []byte) (Answer, error) {
	var rawData Answer

	err := json.Unmarshal(resp, &rawData)

	if err != nil {
		log.Fatalln(err)
	}

	return rawData, nil
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

func formatLink(Link string) string {
	standardURL := "https://stackoverflow.com/q/"
	re := regexp.MustCompile("[0-9]+")
	questionID := re.FindAllString(Link, -1)

	return standardURL + questionID[0]
}
