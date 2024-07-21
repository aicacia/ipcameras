package util

import (
	"fmt"
	"regexp"
	"time"
)

var NumericRegex = regexp.MustCompile(`\D+`)

type TimeRangeST struct {
	Start time.Time
	End   *time.Time
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetUniqueKey[T any](m map[string]T, key string) string {
	uniqueKey := key
	counter := 1
	for {
		if _, ok := m[uniqueKey]; !ok {
			break
		}
		uniqueKey = fmt.Sprintf("%s_%d", key, counter)
		counter++
	}
	return uniqueKey
}

func UniqueSlice[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}

func TruncateToMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}
