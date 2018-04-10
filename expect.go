package expect

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func ObjectsToBeEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected objects to be equal, but weren't\nwanted: %+v\ngot   : %+v", want, got)
	}
}

// func ToBeNil(t *testing.T, label string, obj interface{}) {
// 	if obj != nil {
// 		t.Errorf("expected %s to be nil, got %s", label, obj)
// 	}
// }

func ToBeNotNil(t *testing.T, label string, obj interface{}) {
	if obj == nil {
		t.Errorf("expected %s to be not nil", label)
	}
}

func StringToBeBlank(t *testing.T, label, s string) {
	if s != "" {
		t.Errorf("expected %s to be blank, got %s", label, s)
	}
}

func StringNotBlank(t *testing.T, label, s string) {
	if s == "" {
		t.Errorf("expected %s to not be blank", label)
	}
}

func StringEqual(t *testing.T, label, want, got string) {
	if want != got {
		t.Errorf("%s not equal. wanted %s, got %s", label, want, got)
	}
}

func StringNotEqual(t *testing.T, label, want, got string) {
	if want == got {
		t.Errorf("expected %s to not be %s, but it was", label, want)
	}
}

func TimeNotEmpty(t *testing.T, label string, tt time.Time) {
	if tt.IsZero() {
		t.Errorf("expected %s to be set, but was empty", label)
	}
}

func ToBeTrue(t *testing.T, label string, val bool) {
	if !val {
		t.Errorf("expected %s to be true, but was false", label)
	}
}

func ToBeFalse(t *testing.T, label string, val bool) {
	if val {
		t.Errorf("expected %s to be false, but was true", label)
	}
}

func ErrorToEqual(t *testing.T, want string, err error) {
	got := err.Error()
	if want != got {
		t.Errorf("expected error to be %s, but got %s", want, got)
	}
}

func ErrorToBeNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected error to be nil, but got %v", err)
	}
}

func ErrorToBeNotNil(t *testing.T, err error) {
	if err == nil {
		t.Fatal("expected error, but got nil")
	}
}

func IntToEqual(t *testing.T, label string, want, got int) {
	if want != got {
		t.Errorf("expected %s to be %d, but got %d", label, want, got)
	}
}

func Int64ToEqual(t *testing.T, label string, want, got int64) {
	if want != got {
		t.Errorf("expected %s to be %d, but got %d", label, want, got)
	}
}

func UInt8ToEqual(t *testing.T, label string, want, got uint8) {
	if want != got {
		t.Errorf("expected %s to be %d, but got %d", label, want, got)
	}
}

func StringSliceToContain(t *testing.T, label, want string, gots []string) {
	for _, got := range gots {
		if want == got {
			return
		}
	}
	t.Errorf("expected %s to contain %s, but didn't", label, want)
}

func HTTPStatusToEqual(t *testing.T, want, got int) {
	if want != got {
		t.Fatalf(
			"expected status %d - %s, but got %d - %s",
			want, http.StatusText(want),
			got, http.StatusText(got),
		)
	}
}
