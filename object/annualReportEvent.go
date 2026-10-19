package object

// AnnualReportEvent is a small reference to an annual report, attached to a
// NotificationGroup when its type is "annual_report".
// https://docs.joinmastodon.org/entities/NotificationGroup/#annual_report
type AnnualReportEvent struct {
	Year string `json:"year"`
}
