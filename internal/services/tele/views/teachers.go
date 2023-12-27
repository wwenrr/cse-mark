package views

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/models"
)

func RenderTeacherProfile(courses []*models.CourseSettingsModel) string {
	t := table.NewWriter()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Course", "Cnt", "Live"})
	for _, c := range courses {
		liveStr := "✓"
		if !data.IsUpdatedCourse(c) {
			liveStr = "✗"
		}

		t.AppendRow(table.Row{c.Course, c.RecordCnt, liveStr})
	}

	return t.Render()
}
