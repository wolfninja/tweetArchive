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
	"log"

	"github.com/chimeracoder/anaconda"
	"gopkg.in/couchbase/gocb.v1"
)

func LoadAndStoreTweets(bucket *gocb.Bucket, api *anaconda.TwitterApi, handle string) (*uint32, error) {
	var total uint32

	count := 200
	//Check old tweets
	log.Printf("Checking old tweets for %s", handle)
	for count > 0 {
		minID, err := GetMinTweetID(bucket, handle)
		if err != nil {
			return nil, err
		}

		var maxID *uint64
		if minID != nil {
			var i = *minID - 1
			maxID = &i
		}

		tweets, err := QueryTweets(api, handle, 200, maxID, nil)
		if err != nil {
			return nil, err
		}
		log.Printf("Loaded %d old tweets", len(tweets))

		count = len(tweets)

		if count > 0 {
			docs, err := InsertTweets(bucket, tweets)
			if err != nil {
				return nil, err
			}
			total = total + uint32(len(docs))
			log.Printf("Inserted %d old docs", len(docs))
		}
	}

	//Check new tweets
	log.Printf("Checking new tweets for %s", handle)
	count = 200
	for count > 0 {
		maxID, err := GetMaxTweetID(bucket, handle)
		if err != nil {
			return nil, err
		}

		tweets, err := QueryTweets(api, handle, 200, nil, maxID)
		if err != nil {
			return nil, err
		}
		log.Printf("Loaded %d new tweets", len(tweets))

		count = len(tweets)

		if count > 0 {
			docs, err := InsertTweets(bucket, tweets)
			if err != nil {
				return nil, err
			}
			total = total + uint32(len(docs))
			log.Printf("Inserted %d new docs", len(docs))
		}
	}

	return &total, nil
}
