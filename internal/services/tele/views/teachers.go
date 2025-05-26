package views

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/domain/entities"
)

func RenderTeacherProfile(courses []*entities.CourseSettingsModel) string {
	t := table.NewWriter()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Course", "Cnt", "Till"})
	for _, c := range courses {
		t.AppendRow(table.Row{c.Course, c.RecordCnt, data.CourseUpdateTill(c).Format("0106")})
	}

	return t.Render()
}
