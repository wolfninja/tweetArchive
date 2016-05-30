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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wolfninja/tweetArchive/configs"
	"github.com/wolfninja/tweetArchive/helpers"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch load new tweets for all users",
	Long: `
Batch loads all new tweets for the handles under
the 'handles' config.`,
	RunE: Batch,
}

func init() {
	RootCmd.AddCommand(batchCmd)
}

func doBatch(handles []string) error {
	bucket, err := helpers.OpenBucket()
	if err != nil {
		return err
	}
	defer bucket.Close()

	api := helpers.OpenTwitterAPI()
	defer api.Close()

	for _, handle := range handles {
		fmt.Printf("Loading tweets for %s\n", handle)
		count, err := helpers.LoadAndStoreTweets(bucket, api, handle)
		if err == nil {
			fmt.Printf("Loaded %d tweets for %s\n", *count, handle)
		}
	}
	return err
}

func Batch(cmd *cobra.Command, args []string) error {
	if !viper.IsSet(configs.CouchbaseURL) {
		return errors.New("couchbase.url must be provided")
	}
	if !viper.IsSet(configs.CouchbaseBucket) {
		return errors.New("couchbase.bucket must be provided")
	}
	if !viper.IsSet(configs.CouchbasePassword) {
		return errors.New("couchbase.password must be provided")
	}
	if !viper.IsSet(configs.CouchbaseURL) {
		return errors.New("couchbase.url must be provided")
	}
	if !viper.IsSet(configs.Handles) {
		return errors.New("handles must be provided")
	}

	return doBatch(viper.GetStringSlice(configs.Handles))
}
