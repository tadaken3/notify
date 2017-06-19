// Copyright © 2017 tadaken3 <k.tanaka6057@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var channel string

// slackCmd represents the slack command
var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "notify to slact",
	Long:  "notify to slact",
	Run: func(cmd *cobra.Command, args []string) {
		var apiUrl string
		baseUrl := "https://slack.com/api/chat.postMessage?token="
		message = url.QueryEscape(message)

		if len(token) == 0 {
			token = viper.GetString("slack_token")
		}
		
		if len(channel) == 0 {
			roomId = viper.GetString("channel")
		}
		
		
		apiUrl = baseUrl + token + "&channel=" + channel + "&text=" + message
		req, err := http.Get(apiUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("status:", req.Status)

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(body))

	},
}

func init() {
	RootCmd.AddCommand(slackCmd)
	slackCmd.Flags().StringVarP(&channel, "channel", "c", "general", "slack channel")
}
