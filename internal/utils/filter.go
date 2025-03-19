package utils

import "time"

func Filter[T any](ss []T, predicate func(T) bool) (ret []T) {
	for _, s := range ss {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return
}

func FilterHour12(dt time.Time) bool {
	return 7 < dt.Hour() && dt.Hour() < 13
}

func FilterHour16(dt time.Time) bool {
	return 11 < dt.Hour() && dt.Hour() < 17
}

func FilterHour20(dt time.Time) bool {
	return 15 < dt.Hour() && dt.Hour() < 22
}

func Avg(vals []float64) (ret float64) {
	for _, val := range vals {
		ret += val
	}

	return ret / float64(len(vals))
}

func MapTemps[T any, TO any](ss []T, mapper func(item T) TO) []TO {
	result := make([]TO, len(ss))
	for i, field := range ss {
		result[i] = mapper(field)
	}
	return result
}
