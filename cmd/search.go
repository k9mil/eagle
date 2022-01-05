package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		if len(args) != 2 {
			return errors.New("requires a sort & title argument")
		}

		ListOfOptions := []string{"votes", "activity", "creation", "relevance"}
		err := stringInSlice(args[0], ListOfOptions)

		if err != nil {
			log.Fatalln(err)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		search(args[0], args[1])
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
	searchCmd.Flags().StringVar(&Sort, "sort", "s", "The sort method to be used.")
	searchCmd.Flags().StringVar(&Title, "title", "t", "The title of the query.")
	rootCmd.AddCommand(searchCmd)
}

func search(sort, title string) {
	url := fmt.Sprintf("https://api.stackexchange.com/2.3/search?order=desc&sort=%s&intitle=%s&site=stackoverflow", sort, title)
	apiReturn := apiCall(url)
	broadcastAnswer(apiReturn)
}

func apiCall(url string) Answer {
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

func decodeJSON(resp []byte) Answer {
	var rawData Answer

	err := json.Unmarshal(resp, &rawData)

	if err != nil {
		log.Fatalln(err)
	}

	return rawData
}

func broadcastAnswer(a Answer) {
	for _, item := range a.Items {
		fmt.Printf("Title: %+v\n", item.Title)
		fmt.Printf("Amount of Answers: %+v\n", item.AnswerCount)
		fmt.Printf("Upvotes: %+v\n\n", item.Score)
		fmt.Printf("Link: %+v\n", item.Link)
	}
}
