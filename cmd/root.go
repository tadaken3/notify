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
	"strings"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var message string
var token string


var RootCmd = &cobra.Command{
	Use:   "notify",
	Short:`line notify`,
	Long: `This application is a simple CLI tool. You can quickly notify to line`,

Run: func(cmd *cobra.Command, args []string) {
	  var endpoint string
		endpoint =  "https://notify-api.line.me/api/notify"
		data := url.Values{"message": {message}}
		req, err := http.NewRequest(
				"POST",
				endpoint,
		 		strings.NewReader(data.Encode()),
		)
	 if len(token) == 0 {
			token = viper.GetString("token")
		}
	 if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization" ,  "Bearer " + token)
		client := &http.Client{}
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
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&message, "message", "m", "Herro, world from Go","message")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.notify.yaml)")
	RootCmd.PersistentFlags().StringVarP(&token, "token", "l", "", "line access token")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".notify")
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
