package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/pkg/errors"
)

type fileRepo struct {
	path string
}

func newFileRepo(path string) *fileRepo {
	return &fileRepo{path: path}
}

func (f *fileRepo) Read() ([]time.Time, error) {
	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		return []time.Time{}, nil
	}

	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file '%s'", f.path)
	}

	return decode(data)
}

func (f *fileRepo) Write(times []time.Time) error {
	data, err := encode(times)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(f.path, data, 0644); err != nil {
		return errors.Wrapf(err, "failed to write to '%s'", f.path)
	}

	return nil
}

func decode(data []byte) ([]time.Time, error) {
	var decoded map[string][]string
	if err := json.Unmarshal(data, &decoded); err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	dates := make([]string, 0, len(decoded))
	for d := range decoded {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	loc := time.Now().Location()
	times := []time.Time{}

	for _, d := range dates {
		for _, e := range decoded[d] {
			timeString := d + " " + e
			tm, err := time.ParseInLocation("2006-01-02 15:04", timeString, loc)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to parse entry %s", timeString)
			}
			times = append(times, tm)
		}
	}

	return times, nil
}

func encode(times []time.Time) ([]byte, error) {
	data := map[string][]string{}

	for _, t := range times {
		date := t.Format("2006-01-02")
		if _, exists := data[date]; !exists {
			data[date] = []string{}
		}
		data[date] = append(data[date], t.Format("15:04"))
	}

	encoded, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "json encode failed")
	}

	return encoded, nil
}
