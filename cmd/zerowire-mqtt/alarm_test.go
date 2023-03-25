package main

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	tests := []struct {
		a    alarm
		want string
		err  string
	}{
		{alarm{"01000000000000"}, "disarmed", ""},
		{alarm{"01000000010000"}, "pending", ""},
		{alarm{"01000100000000"}, "armed_home", ""},
		{alarm{"01000001000000"}, "armed_away", ""},
		{alarm{}, "", "invalid zerowire alarm status"},
	}

	for _, foo := range tests {
		s, e := foo.a.GetStatus()
		if s != foo.want {
			t.Errorf("wanted %s got %s", foo.want, s)
		}
		if e != nil && e.Error() != foo.err {
			t.Errorf("wanted %s got %s", foo.err, e)
		}
	}
}
