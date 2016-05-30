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

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load [Twitter name]",
	Short: "Load + store tweets for the given twitter user",
	Long: `
Load and store all new tweets for the given twitter user.
All new tweets accessible are loaded and stored in the data store`,
	RunE: Load,
}

func init() {
	RootCmd.AddCommand(loadCmd)
}

func doLoad(handle string) error {
	fmt.Printf("Loading %s\n", handle)
	bucket, err := helpers.OpenBucket()
	if err != nil {
		return err
	}
	defer bucket.Close()

	api := helpers.OpenTwitterAPI()
	defer api.Close()

	count, err := helpers.LoadAndStoreTweets(bucket, api, handle)
	if err == nil {
		fmt.Printf("Loaded %d tweets for %s\n", *count, handle)
	}
	return err
}

func Load(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("twitter username needs to be provided")
	}
	for _, config := range configs.AllRequired {
		if !viper.IsSet(config) {
			return errors.New(config + " must be provided")
		}
	}

	userName := args[0]
	return doLoad(userName)
}
