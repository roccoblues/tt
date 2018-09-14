package main

import (
	"testing"
	"time"
)

func Test_format(t *testing.T) {
	testTime1, err := time.Parse("2006-01-02 15:04", "2018-09-01 10:00")
	if err != nil {
		t.Fatal(err)
	}
	testTime2 := testTime1.Add(time.Hour * 2)
	testTime3 := testTime1.Add(time.Hour * 4)
	testTime4 := testTime1.Add(time.Hour * 24)
	testTime5 := testTime1.Add(time.Hour * 24 * 8)

	days := []*day{
		&day{
			Date: testTime1,
			Entries: []*entry{
				&entry{
					Start: testTime1,
					End:   testTime2,
				},
				&entry{
					Start: testTime3,
				},
			},
		},
		&day{
			Date: testTime4,
			Entries: []*entry{
				&entry{
					Start: testTime4,
				},
			},
		},
		&day{
			Date: testTime5,
			Entries: []*entry{
				&entry{
					Start: testTime5,
				},
			},
		},
	}

	tests := []struct {
		name string
		days []*day
		want string
	}{
		{
			name: "works",
			days: days,
			want: "01.09.2018  2.00  10:00-12:00 14:00-\n02.09.2018  0.00  10:00-\n\n09.09.2018  0.00  10:00-\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.days); got != tt.want {
				t.Errorf("format() = %v, want %v", got, tt.want)
			}
		})
	}
}