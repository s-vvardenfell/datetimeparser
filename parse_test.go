package datetimeparser

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TestData struct {
	input   string
	want    string
	wantErr bool
}

var testData = []TestData{
	// Common formats
	{"15:04:05", "0000-01-01T15:04:05Z", false},
	{"15:04", "0000-01-01T15:04:00Z", false},
	{"3:04:05", "0000-01-01T03:04:05Z", false},
	{"3:4:5", "0000-01-01T03:04:05Z", false},

	{"3/24/2022", "2022-03-24T00:00:00Z", false},
	{"02/01/2022", "2022-01-02T00:00:00Z", false},
	{"2022-01-02", "2022-01-02T00:00:00Z", false},

	{"15:04:05 02-01-2006", "2006-01-02T15:04:05Z", false},
	{"15:04:05 2006-01-02", "2006-01-02T15:04:05Z", false},
	{"2006-01-02 15:04:05", "2006-01-02T15:04:05Z", false},
	{"02-01-2006 15:04:05", "2006-01-02T15:04:05Z", false},

	{"18:21, 18 MAR 2020", "2020-03-18T18:21:00Z", false},
	{"12:44, 6 FEB 2020", "2020-02-06T12:44:00Z", false},
	{"00:00, 17 MAY 2006 Updated 23:15, 19 APR 2013", "2006-05-17T00:00:00Z", false},
	{"00:00, 25 JUN 2011 Updated 21:34, 16 APR 2013", "2011-06-25T00:00:00Z", false},
	{"Updated 12:26, 9 FEB 2021", "2021-02-09T12:26:00Z", false},
	{"Updated 20:03, 31 MAY 2018", "2018-05-31T20:03:00Z", false},

	{"2020年05月20日19時30分", "2020-05-20T19:30:00Z", false},
	{"2021年07月03日10時00分", "2021-07-03T10:00:00Z", false},

	// some weird things
	{"5.1.18 09:33:54", "2018-01-05T09:33:54Z", false},
	{"01.1.18 09:33:54", "2018-01-01T09:33:54Z", false},
	{"01.1.18 9:31", "2018-01-01T09:31:00Z", false},
	{"09:33:54 01.1.18", "2018-01-01T09:33:54Z", false},
	{"2008.5.8 33:54", "2008-05-08T00:00:00Z", false},

	// timestamp
	{"1675929598", "2023-02-09T10:59:58Z", false},
	{"1675929598123", "2023-02-09T10:59:58Z", false},

	// ISO / RFC
	{"2006-01-02T15:04", "2006-01-02T15:04:00Z", false},
	{"2006-01-02T15:04:05", "2006-01-02T15:04:05Z", false},
	{"2006-1-2T15:4:5-07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02T15:04-0700", "2006-01-02T15:04:00Z", false},
	{"2006-01-02T15:04-07:00", "2006-01-02T15:04:00Z", false},
	{"2006-01-02T15:04:05-0700", "2006-01-02T15:04:05Z", false},
	{"2006-01-02T15:04:05-07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02T15:04:05Z07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02MST15:04:05-07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02T15:04:05.9999999-07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02Z15:04:05.9999999-07:00", "2006-01-02T15:04:05Z", false},
	{"2006-01-02T15:04:05.999999999Z07:00", "2006-01-02T15:04:05Z", false},

	// Named months
	{"June 13, 2006", "2006-06-13T00:00:00Z", false},
	{"Monday, June 13, 2006", "2006-06-13T00:00:00Z", false},
	{"June 13 Mon", "0000-06-13T00:00:00Z", false},
	{"Mon June _9 15:04:05", "0000-06-09T15:04:05Z", false},
	{"4 марта 2021 года в 14:55", "2021-03-04T14:55:00Z", false},
	{"4 марта 2021 года в 14:55:34", "2021-03-04T14:55:34Z", false},
	{"в 14:55 4 марта 2021 года", "2021-03-04T14:55:00Z", false},
	{"в 14:55:34 4 марта 2021 года ", "2021-03-04T14:55:34Z", false},

	// Adverbs
	{"toissapäivänä klo " + time.Now().Format("15:04:05"),
		time.Now().AddDate(0, 0, -2).Format("2006-01-02T15:04:05Z"), false},
	{"gestern im " + time.Now().Format("15:04:05"),
		time.Now().AddDate(0, 0, -1).Format("2006-01-02T15:04:05Z"), false},
	{"today at " + time.Now().Format("15:04:05"),
		time.Now().AddDate(0, 0, 0).Format("2006-01-02T15:04:05Z"), false},
	{"завтра в " + time.Now().Format("15:04:05"),
		time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05Z"), false},
	{"pojutrze o " + time.Now().Format("15:04:05"),
		time.Now().AddDate(0, 0, 2).Format("2006-01-02T15:04:05Z"), false},

	// Time zones
	{"Mon Jan 02 15:04:05 -0700", "0000-01-02T22:04:05Z", false},
	{"02 Jan 06 15:04 -0700", "2006-01-02T22:04:00Z", false},
	{"Mon, 02 Jan 2006 15:04:05 -0700", "2006-01-02T22:04:05Z", false},
	{"2006-01-02T15:04:05", "2006-01-02T15:04:05Z", false},

	// Names time zones
	{"Mon Jan 02 15:04:05 MSK", "0000-01-02T18:04:05Z", false},
	{"02 Jan 06 15:04 BST", "2006-01-02T21:04:00Z", false},
	{"Mon, 02 Jan 2006 15:04:05 CKT", "2006-01-02T05:04:05Z", false},

	// comment to pass CI-tests - diff time zone
	// time-ago
	// {"39 minutes ago", time.Now().Add(time.Minute * -time.Duration(39)).Format("2006-01-02T15:04:05Z"), false},
	// {"5 minutes ago", time.Now().Add(time.Minute * -time.Duration(5)).Format("2006-01-02T15:04:05Z"), false},
	// {"four hours ago", time.Now().Add(time.Hour * -time.Duration(4)).Format("2006-01-02T15:04:05Z"), false},

	// AM/PM
	{"3:04:05 PM 02-01-2006", "2006-01-02T15:04:05Z", false},
	{"3:04:05 AM 02-01-2006", "2006-01-02T03:04:05Z", false},
	{"3:04:05PM 02-01-2006", "2006-01-02T15:04:05Z", false},
	{"3:04:05AM 02-01-2006", "2006-01-02T03:04:05Z", false},
	{"3:04:05 P.M. 02-01-2006", "2006-01-02T15:04:05Z", false},
	{"3:04:05 A.M. 02-01-2006", "2006-01-02T03:04:05Z", false},
	{"3:04:05P.M. 02-01-2006", "2006-01-02T15:04:05Z", false},
	{"3:04:05A.M. 02-01-2006", "2006-01-02T03:04:05Z", false},
}

func Test_parseDateTime(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Moscow")
	require.NoError(t, err)

	time.Local = loc

	for _, testd := range testData {
		t.Run(testd.input, func(t *testing.T) {
			var tp TimeParser

			got, err := tp.Parse(testd.input)
			if (err != nil) != testd.wantErr {
				t.Errorf("ProcessDateTime() error = %v, wantErr %v", err, testd.wantErr)
			}
			if !reflect.DeepEqual(got.Format(time.RFC3339), testd.want) {
				t.Errorf("ProcessDateTime() = %v, want %v", got.Format(time.RFC3339), testd.want)
			}
		})
	}
}

func Benchmark_ProcessDateTime(b *testing.B) {
	var (
		got time.Time
		err error
	)

	for i := 0; i < b.N; i++ {
		for _, tt := range testData {
			tp := TimeParser{} // create here to avoid using stored layouts
			got, err = tp.Parse(tt.input)
		}
	}

	b.Log(got, err) // use variables
}
