package configs

const (
	CouchbaseBucket       = "couchbase.bucket"
	CouchbaseURL          = "couchbase.url"
	CouchbasePassword     = "couchbase.password"
	Handles               = "handles"
	TwitterConsumerKey    = "twitter.consumer.key"
	TwitterConsumerSecret = "twitter.consumer.secret"
	TwitterAccessToken    = "twitter.access.token"
	TwitterAccessSecret   = "twitter.access.secret"
)

var (
	CouchbaseRequired = []string{CouchbaseBucket, CouchbasePassword, CouchbaseURL}
	TwitterRequired   = []string{TwitterAccessSecret, TwitterAccessToken, TwitterConsumerKey, TwitterConsumerSecret}
	AllRequired       = append(TwitterRequired, CouchbaseRequired...)
)
