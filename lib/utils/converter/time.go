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

func ConvertPointerToTime(value *time.Time) time.Time {
	var result time.Time

	if value != nil {
		result = *value
	}

	return result
}
