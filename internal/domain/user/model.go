package user

type Model struct {
	UserId    string `json:"user_id" bson:"user_id"`
	IsTeacher bool   `json:"is_teacher" bson:"is_teacher"`
	GrantedBy string `json:"granted_by" bson:"granted_by"`
}
