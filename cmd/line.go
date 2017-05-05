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
	"log"
	"os"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lineCmd = &cobra.Command{
	Use:   "line",
	Short: "notify to line",
	Long: "notify to line",
	Run: func(cmd *cobra.Command, args []string) {
		var apiUrl string
		apiUrl =  "https://notify-api.line.me/api/notify"
		data := url.Values{"message": {message}}
		
		if len(token) == 0 {
			token = viper.GetString("line_access_token")
		}
		
		req, err := http.NewRequest(
				"POST",
				apiUrl,
		 		strings.NewReader(data.Encode()),
		)
	 	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		client := &http.Client{}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization" ,  "Bearer " + token)

		resp, err := client.Do(req)
		if err != nil {
			 	log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
		}
		fmt.Printf("%s\n", body)
		fmt.Println(resp.Status)
	},
}

func init() {
	RootCmd.AddCommand(lineCmd)

}
