package data

import "thuanle/cse-mark/internal/services/db"

func AppendQueryHistory(chatId int64, course string, studentId string) error {
	return db.Instance().AppendUserHistory(chatId, course, studentId)
}

func ClearStudentQueries(chatId int64) error {
	return db.Instance().ClearUserHistory(chatId)
}
