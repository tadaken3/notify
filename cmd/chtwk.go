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
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var roomId string

var chtwkCmd = &cobra.Command{
	Use:   "chtwk",
	Short: "notify to chatwork",
	Long: "notify to chatwork",
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(roomId) == 0 {
			roomId = viper.GetString("room_id")
		}
		
		apiUrl := "https://api.chatwork.com/"
		resource := "/v1/rooms/" + roomId + "/messages"

		u, _ := url.ParseRequestURI(apiUrl)
		u.Path = resource
		urlStr := fmt.Sprintf("%v", u) 
		
		data := url.Values{}
		data.Set("body", message)
		//fmt.Println(data.Encode())

		req, err := http.NewRequest(
				"POST",
				urlStr,
				bytes.NewBufferString(data.Encode()),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		if len(token) == 0 {
			token = viper.GetString("chatwork_access_token")
		}
		
		client := &http.Client{}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("X-ChatWorkToken", token)

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
	RootCmd.AddCommand(chtwkCmd)
	chtwkCmd.Flags().StringVarP(&roomId, "roomId", "r", "", "chatwork room id")
}
