package course

type Model struct {
	Id         string `json:"course" bson:"course"`
	Link       string `json:"link" bson:"link"`
	ByTeleId   int64  `json:"by_id" bson:"by_id"`
	ByTeleUser string `json:"by_user,omitempty" bson:"by_user,omitempty"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at"`
	RecordCnt  int64  `json:"record_cnt" bson:"record_cnt"`
}
