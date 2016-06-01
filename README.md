# tweetArchive: Simple tool for archiving tweets in a Couchbase store

## Features
- Downloads the full timelime of tweets offered by the twitter API for a user
- Subsequent loads only download the newest tweets not present in the store
- Supports loading tweets for a single user, as well as a batch mode for loading multiple users
- Ability to dump the stored tweets for a user to stdout

## Background
- I wanted an analysis of the tweets of the 2016 Presidential candidates, but was unable to find a simple archive of historical tweets that have since been dropped off the twitter API. I decided to start capturing them for current+future analysis.

## Current usage
- I'm currently running the batch mode as a periodic(every 30 mins) cron job on a VPS, grabbing the latest tweets for the Presidential candidates

## Latest version
- Currently in alpha/development.

## Requirements
- A Couchbase installation (hint: the Docker-based install works pretty great for dev purposes)
- Go installed on your system (godep strongly recommended)
- Twitter account with your API keys and secrets

## Getting started
- _Optional: Create a new Couchbase bucket to hold your tweets_
- Clone or download the repo
- Copy the supplied _tweetArchive.yml.sample_ to _$HOME/.tweetArchive.yml_
- Edit the _couchbase_ and _twitter_ sections _$HOME/tweetArchive.yml_ to reflect your environment
  - Configuration can also be supplied as env variables (e.g. COUCHBASE.URL, COUCHBASE.BUCKET, etc.)
- Run _godep go install_ in the root of the cloned repo to install
- Run _tweetArchive help_ from your $GOPATH/bin to get usage info
  - _tweetArchive load <username>_ will pull in all tweets for a user

## Usage modes
- Run _tweetArchive help_ for general help or _tweetArchive <command> --help_ for help about a specific command
- tweetArchive has 3 commands:
  - Batch command
    - _tweetArchive batch_ will load and persist all the tweets for the list of users configured in your yaml config in the _handles_ key
      - Can also be configured using the _HANDLES_ env variable
  - Load command
    - _tweetArchive load <handle>_ will load and persist all the tweets for the given twitter user
  - Query command
    - _tweetArchive query <handle>_ will query and dump to stdout all the tweets for the given twitter user (one tweet per line, newlines replaced with spaces)

## Versioning
  - This project uses [Semantic Versioning](http://semver.org/) to make release versions predictable
  - Versions consist of MAJOR.MINOR.PATCH
    - Different MAJOR versions are not guaranteed to be API compatible (e.g. existing commands have changed)
    - Incrementing MINOR versions within the same MAJOR version contain additional functionality, with existing calls being compatible (e.g. new commands added)
    - Different PATCH versions withing the same MAJOR.MINOR version are completely API compatible (e.g. bugfixes, no changes to commands themselves)

## Branches
  - *master* is the stable branch which releases are built from
  - other branches are not guaranteed to be stable and assumed to be under active feature development

## Changelog
See CHANGELOG.md for a full history of release changes

## License
Licensed under the MIT License (see LICENSE)
