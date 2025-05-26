package entities

type CourseSettingsModel struct {
	Course    string `json:"course" bson:"course"`
	Link      string `json:"link" bson:"link"`
	ById      int64  `json:"by_id" bson:"by_id"`
	ByUser    string `json:"by_user,omitempty" bson:"by_user,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
	RecordCnt int64  `json:"record_cnt" bson:"record_cnt"`
}

type UserSettingsModel struct {
	UserId    string            `json:"user_id" bson:"user_id"`
	IsTeacher bool              `json:"is_teacher" bson:"is_teacher"`
	GrantedBy string            `json:"granted_by" bson:"granted_by"`
	Queries   map[string]string `json:"queries,omitempty" bson:"queries,omitempty"`
}
