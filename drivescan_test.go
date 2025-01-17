package main

import (
	"testing"
)

func TestGetDriveType(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"C:\\\\", "Fixed"},
		{"A:\\\\", "Removable"},
		{"Z:\\\\", "Network"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := getDriveType(tt.path)
			if got != tt.want {
				t.Errorf("getDriveType(%s) = %s; want %s", tt.path, got, tt.want)
			}
		})
	}
}

func TestGetDrives(t *testing.T) {
	drives, err := getDrives()
	if err != nil {
		t.Errorf("getDrives() returned an error: %v", err)
	}

	if len(drives) == 0 {
		t.Errorf("getDrives() returned no drives")
	}
}
