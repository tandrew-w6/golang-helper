package converter

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertDateToDatetime(req time.Time) (time.Time, error) {
	dateTime, err := time.Parse(time.RFC3339, req.Format("2006-01-02")+"T16:59:59Z")
	if err != nil {
		return req, err
	}

	return dateTime, nil
}

func ConvertGoTimePointerToProtoTimePointer(value *time.Time) *timestamppb.Timestamp {
	var result *timestamppb.Timestamp

	if value != nil {
		result = timestamppb.New(*value)
	}

	return result
}

func ConvertGoTimeToProtoTimePointer(value time.Time) *timestamppb.Timestamp {
	result := timestamppb.New(value)

	return result
}

func ConvertPointerToTime(value *time.Time) time.Time {
	var result time.Time

	if value != nil {
		result = *value
	}

	return result
}

func ConvertTimeToPointer(value time.Time) *time.Time {
	var result *time.Time = &value

	return result
}

func ConvertPointerTimeToString(value *time.Time, layout string) string {
	result := ""

	if layout == "" {
		layout = time.RFC3339
	}

	if value != nil {
		t := *value
		result = t.Format(layout)
	}

	return result
}

func ConvertTimeToString(value time.Time, layout string) string {
	if layout == "" {
		layout = time.RFC3339
	}

	result := value.Format(layout)

	return result
}

func ConvertTimeToPointerString(value time.Time, layout string) *string {
	if layout == "" {
		layout = time.RFC3339
	}

	result := value.Format(layout)

	return &result
}