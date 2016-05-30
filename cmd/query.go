// Copyright © 2016 Nick Ball <nick@wolfninja.com>
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
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wolfninja/tweetArchive/configs"
	"github.com/wolfninja/tweetArchive/helpers"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query [Twitter name]",
	Short: "Query for all tweets for the given twitter handle",
	Long: `
Query and output for all tweets for the given twitter handle.`,
	RunE: Query,
}

func init() {
	RootCmd.AddCommand(queryCmd)
}

func doQuery(handle string) error {
	bucket, err := helpers.OpenBucket()
	if err != nil {
		return err
	}
	defer bucket.Close()

	log.Printf("Running for %s", handle)
	tweets, err := helpers.GetTweets(bucket, handle)
	if err != nil {
		return err
	}
	for _, tweet := range tweets {
		fmt.Println(tweet)
	}
	return nil
}

func Query(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("twitter username needs to be provided")
	}
	if !viper.IsSet(configs.CouchbaseURL) {
		return errors.New("couchbase.url must be provided")
	}
	if !viper.IsSet(configs.CouchbaseBucket) {
		return errors.New("couchbase.bucket must be provided")
	}
	if !viper.IsSet(configs.CouchbasePassword) {
		return errors.New("couchbase.password must be provided")
	}

	userName := args[0]
	return doQuery(userName)
}
