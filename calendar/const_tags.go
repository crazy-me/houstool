package calendar

type TimeFormat string

const (
	FMT     TimeFormat = "2006-01-02 15:04:05"
	FMT_BAR TimeFormat = "2006/01/02 15:04:05"
	YEAR    TimeFormat = "2006"
	Month   TimeFormat = "01"
	DAY     TimeFormat = "02"
	Hour    TimeFormat = "15"
)

func (t TimeFormat) ToString() string {
	switch t {
	case FMT:
		return "2006-01-02 15:04:05"
	case FMT_BAR:
		return "2006/01/02 15:04:05"
	case YEAR:
		return "2006"
	case Month:
		return "01"
	case DAY:
		return "02"
	case Hour:
		return "15"
	default:
		return "Unknown"
	}
}
