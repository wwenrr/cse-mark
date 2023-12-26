package configs

import "time"

var AdminChatIds []int64

var (
	MongoHost            = "localhost"
	MongoPort            = "27017"
	DbTransactionTimeout = 10 * time.Second

	DbMark = "mark-cse"

	DbSettings        = "mark-settings"
	DbSettingsUsers   = "users"
	DbSettingsCourses = "courses"

	FetchMaxAge = 6 * 30 * 24 * time.Hour // 6 months
)
