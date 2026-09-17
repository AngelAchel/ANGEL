package platform

import (
	"os"
	"runtime"
	"strings"
)

// IsTermux returns true if the platform is running under Termux (Android)
func IsTermux() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	return isTermuxEnv()
}

func isTermuxEnv() bool {
	if os.Getenv("TERMUX_VERSION") != "" {
		return true
	}
	if os.Getenv("PREFIX") == "/data/data/com.termux/files/usr" {
		return true
	}
	if _, err := os.Stat("/data/data/com.termux"); err == nil {
		return true
	}
	return false
}

// IsKaliLinux returns true if running on Kali Linux
func IsKaliLinux() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	return getDistro() == "kali"
}

func getDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	content := string(data)
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "ID=") {
			return strings.Trim(line[len("ID="):], `"`)
		}
	}
	return ""
}

// GetPlatform returns the current platform name
func GetPlatform() string {
	switch {
	case IsTermux():
		return "termux"
	case IsKaliLinux():
		return "kali"
	case runtime.GOOS == "darwin":
		return "macos"
	case runtime.GOOS == "windows":
		return "windows"
	default:
		return runtime.GOOS
	}
}

// GetArchitecture returns the platform-specific architecture
func GetArchitecture() string {
	arch := runtime.GOARCH
	if IsTermux() {
		if arch == "arm64" {
			return "arm64-v8a"
		}
		return "arm"
	}
	return arch
}

// GetBinaryExtension returns the file extension for the current platform
func GetBinaryExtension() string {
	if GetPlatform() == "windows" {
		return ".exe"
	}
	return ""
}

// FileExists checks if a file exists on the filesystem
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetEnv returns the value of an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}

// IsAndroid returns true if running on Android
func IsAndroid() bool {
	return IsTermux() || os.Getenv("ANDROID_ROOT") != ""
}
