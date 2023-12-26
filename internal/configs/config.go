package configs

import "time"

var AdminChatIds []int64

var (
	MongoHost            = "localhost"
	MongoPort            = "27017"
	MongoDb              = "cse-mark"
	DbTransactionTimeout = 10 * time.Second
)
