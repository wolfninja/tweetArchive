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

package helpers

import (
	"net/url"
	"strconv"

	"github.com/chimeracoder/anaconda"
	"github.com/wolfninja/tweetArchive/configs"

	"github.com/spf13/viper"
)

func OpenTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(viper.GetString(configs.TwitterConsumerKey))
	anaconda.SetConsumerSecret(viper.GetString(configs.TwitterConsumerSecret))
	return anaconda.NewTwitterApi(viper.GetString(configs.TwitterAccessToken), viper.GetString(configs.TwitterAccessSecret))
}

func QueryTweets(api *anaconda.TwitterApi, username string, count int64, maxID *uint64, sinceID *uint64) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("screen_name", username)
	v.Set("count", strconv.FormatInt(count, 10))
	v.Set("exclude_replies", "false")
	v.Set("include_rts", "true")
	if maxID != nil {
		v.Set("max_id", strconv.FormatUint(*maxID, 10))
	}
	if sinceID != nil {
		v.Set("since_id", strconv.FormatUint(*sinceID, 10))
	}

	tweets, err := api.GetUserTimeline(v)
	return tweets, err
}
