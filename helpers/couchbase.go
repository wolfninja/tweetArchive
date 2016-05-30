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
	"fmt"
	"strconv"

	"github.com/chimeracoder/anaconda"
	"github.com/wolfninja/tweetArchive/configs"

	"github.com/spf13/viper"
	"gopkg.in/couchbase/gocb.v1"
)

func DoN1qlQuery(bucket *gocb.Bucket, query string, params []interface{}) (*uint64, error) {
	myQuery := gocb.NewN1qlQuery(query)
	rows, err := bucket.ExecuteN1qlQuery(myQuery, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var row map[string]string
	for rows.Next(&row) {
		val := row["val"]
		if val == "" {
			return nil, nil
		}
		i64, _ := strconv.ParseUint(val, 10, 64)
		return &i64, nil
	}
	return nil, nil
}

func InsertTweets(bucket *gocb.Bucket, tweets []anaconda.Tweet) ([]*TweetDoc, error) {
	docs := make([]*TweetDoc, len(tweets), len(tweets))
	for i, tweet := range tweets {
		tweetID := tweet.Id
		userName := tweet.User.ScreenName
		docID := fmt.Sprintf("%s_%d", userName, tweetID)
		doc := &TweetDoc{&tweet, "tweet"}
		bucket.Upsert(docID, doc, 0)
		docs[i] = doc
	}
	return docs, nil
}

func OpenBucket() (*gocb.Bucket, error) {
	cluster, err := gocb.Connect(viper.GetString(configs.CouchbaseURL))
	if err != nil {
		return nil, err
	}
	bucket, err := cluster.OpenBucket(viper.GetString(configs.CouchbaseBucket), viper.GetString(configs.CouchbasePassword))
	return bucket, err
}

type TweetDoc struct {
	*anaconda.Tweet
	Type string `json:"type"`
}
