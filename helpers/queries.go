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
	"strings"

	"gopkg.in/couchbase/gocb.v1"
)

func GetMaxTweetID(bucket *gocb.Bucket, handle string) (*uint64, error) {
	var params []interface{}
	params = append(params, strings.ToLower(handle))
	query := "SELECT id_str AS val FROM tweets WHERE LOWER(`user`.screen_name) = $1 ORDER BY id DESC LIMIT 1"
	//query := "SELECT MAX(id_str) AS val FROM tweets WHERE LOWER(`user`.screen_name) = $1"
	return DoN1qlQuery(bucket, query, params)
}

func GetMinTweetID(bucket *gocb.Bucket, handle string) (*uint64, error) {
	var params []interface{}
	params = append(params, strings.ToLower(handle))
	query := "SELECT id_str as val FROM tweets WHERE LOWER(`user`.screen_name) = $1 ORDER BY id ASC LIMIT 1"
	//query := "SELECT MIN(id_str) AS val FROM tweets WHERE LOWER(`user`.screen_name) = $1"
	return DoN1qlQuery(bucket, query, params)
}

func GetTweets(bucket *gocb.Bucket, handle string) ([]string, error) {
	var params []interface{}
	params = append(params, strings.ToLower(handle))
	queryTxt := "SELECT text FROM tweets WHERE LOWER(`user`.screen_name) = $1"
	query := gocb.NewN1qlQuery(queryTxt)
	rows, err := bucket.ExecuteN1qlQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var row map[string]string
	var texts = make([]string, 0)
	for rows.Next(&row) {
		text := row["text"]
		texts = append(texts, text)
	}
	return texts, nil
}
