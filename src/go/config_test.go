package common

import "testing"

func TestNewHTConfigBase(t *testing.T) {
	cfg := NewHTConfigBase(8080, "/src", "/content", "/logs")
	if cfg.Port != 8080 {
		t.Errorf("expected port 8080, got %d", cfg.Port)
	}
	if cfg.SrcPath != "/src" {
		t.Errorf("expected src /src, got %s", cfg.SrcPath)
	}
	if cfg.ContentPath != "/content" {
		t.Errorf("expected content /content, got %s", cfg.ContentPath)
	}
	if cfg.LogPath != "/logs" {
		t.Errorf("expected log /logs, got %s", cfg.LogPath)
	}
}
