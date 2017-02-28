package monitor

import "testing"

func BenchmarkReportCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReportCall(
			"foo.bar.test_report_call",
			"",
			"bar",
			SUCC,
			1,
		)
	}
}
