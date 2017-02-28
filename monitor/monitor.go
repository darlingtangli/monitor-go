/**
* @file monitor.go
* @brief golang wrapper for c monitor api
* @author litang
* @version 1.0
* @date 2017-02-27
 */
package monitor

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lmonitor
#include <report.h>
*/
import "C"

const (
	SUCC      int = 0
	FAILED    int = 1
	EXCEPTION int = 2
)

func ReportCall(metric, caller, callee string, status int, costUSec uint64) {
	C.moni_report_call(
		C.CString(metric),
		C.CString(caller),
		C.CString(callee),
		C.moni_call_status_t(status),
		C.uint64_t(costUSec),
	)
}

func ReportIncr(metric string, step uint64) {
	C.moni_report_incr(
		C.CString(metric),
		C.uint64_t(step),
	)
}

func ReportStatics(metric string, value uint64) {
	C.moni_report_statics(
		C.CString(metric),
		C.uint64_t(value),
	)
}

func ReportAvg(metric string, value uint64) {
	C.moni_report_avg(
		C.CString(metric),
		C.uint64_t(value),
	)
}

func ReportMin(metric string, value uint64) {
	C.moni_report_min(
		C.CString(metric),
		C.uint64_t(value),
	)
}

func ReportMax(metric string, value uint64) {
	C.moni_report_max(
		C.CString(metric),
		C.uint64_t(value),
	)
}
