package common

import "testing"

func TestHTDateDefault(t *testing.T) {
	d := HTDate{}
	if d.DateType != "" {
		t.Errorf("expected empty DateType, got %q", d.DateType)
	}
}

func TestHTTextDefault(t *testing.T) {
	tt := HTText{}
	if tt.Text != "" {
		t.Errorf("expected empty Text, got %q", tt.Text)
	}
}

func TestClassTemplateFileDefault(t *testing.T) {
	ctf := ClassTemplateFile{}
	if ctf.Version != 0 {
		t.Errorf("expected version 0, got %d", ctf.Version)
	}
}
