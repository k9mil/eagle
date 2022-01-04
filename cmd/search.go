package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Required command to search for your question.",
	Run: func(cmd *cobra.Command, args []string) {
		search()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

type Answer struct {
	AnswerCount int    `json:"answer_count"`
	Score       int    `json:"score"`
	Link        string `json:"link"`
	Title       string `json:"title"`
}

func search() {
	resp, err := http.Get("https://api.stackexchange.com/2.3/posts?order=asc&sort=creation&site=stackoverflow")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
