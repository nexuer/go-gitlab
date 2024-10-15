package gitlab

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
	var zeroDate Date
	ds, err := json.Marshal(zeroDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("zero date: %s", string(ds))

	var zeroDatePtr *Date
	ds, err = json.Marshal(zeroDatePtr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("zero date ptr: %s", string(ds))

	now := time.Now()
	ds, err = json.Marshal(now)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("now time: %s", string(ds))

	nowDate := NewDate(now)
	ds, err = json.Marshal(nowDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("now date: %s", string(ds))
}

func TestDate_UnmarshalJSON(t *testing.T) {
	str := `null`
	var zeroDate Date
	err := json.Unmarshal([]byte(str), &zeroDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("date is zero: %t", zeroDate.IsZero())

	str = `"2020-02-03"`
	err = json.Unmarshal([]byte(str), &zeroDate)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("date: %s", zeroDate)

}
