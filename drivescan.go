package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// DriveInfo holds details about a drive.
type DriveInfo struct {
	Path      string
	Type      string
	FreeSpace uint64
	TotalSpace uint64
}

// getDriveType converts the integer drive type to a string.
func getDriveType(path string) string {
	driveType := syscall.GetDriveType(syscall.StringToUTF16Ptr(path))
	switch driveType {
	case syscall.DRIVE_UNKNOWN:
		return "Unknown"
	case syscall.DRIVE_NO_ROOT_DIR:
		return "Invalid Path"
	case syscall.DRIVE_REMOVABLE:
		return "Removable"
	case syscall.DRIVE_FIXED:
		return "Fixed"
	case syscall.DRIVE_REMOTE:
		return "Network"
	case syscall.DRIVE_CDROM:
		return "CD-ROM"
	case syscall.DRIVE_RAMDISK:
		return "RAM Disk"
	default:
		return "Unknown"
	}
}

// getDriveDetails gathers information about a drive.
func getDriveDetails(path string) (*DriveInfo, error) {
	var stat syscall.Statfs_t
	
	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, fmt.Errorf("could not retrieve stats for %s: %w", path, err)
	}

	freeSpace := stat.Bavail * uint64(stat.Bsize)
	totalSpace := stat.Blocks * uint64(stat.Bsize)
	driveType := getDriveType(path)

	return &DriveInfo{
		Path:      path,
		Type:      driveType,
		FreeSpace: freeSpace,
		TotalSpace: totalSpace,
	}, nil
}

// getDrives scans all drives and gathers details.
func getDrives() ([]DriveInfo, error) {
	drives := []DriveInfo{}

	// Iterate through drive letters A-Z.
	for drive := 'A'; drive <= 'Z'; drive++ {
		path := fmt.Sprintf("%c:\\", drive)
		if _, err := os.Stat(path); err == nil {
			driveInfo, err := getDriveDetails(path)
			if err == nil {
				drives = append(drives, *driveInfo)
			}
		}
	}

	return drives, nil
}

func main() {
	drives, err := getDrives()
	if err != nil {
		fmt.Printf("Error scanning drives: %v\n", err)
		return
	}

	for _, drive := range drives {
		fmt.Printf("Path: %s, Type: %s, Free Space: %d bytes, Total Space: %d bytes\n",
			drive.Path, drive.Type, drive.FreeSpace, drive.TotalSpace)
	}
}
