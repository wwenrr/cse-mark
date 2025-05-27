package views

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"thuanle/cse-mark/internal/domain/course"
)

type TeacherRenderer struct {
	courseRules *course.Rules
}

func NewTeacherRenderer(courseRules *course.Rules) *TeacherRenderer {
	return &TeacherRenderer{
		courseRules: courseRules,
	}
}

func (r *TeacherRenderer) RenderTeacherProfile(courses []course.Model) string {
	t := table.NewWriter()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Course", "Cnt", "Till"})
	for _, c := range courses {
		t.AppendRow(table.Row{c.Id, c.RecordCnt, r.courseRules.CourseUpdateTill(c).Format("0106")})
	}

	return t.Render()
}
