package platform

import "runtime"

// IsTermux returns true if the platform is running under Termux (Android)
func IsTermux() bool {
	return runtime.GOOS == "linux" && isTermuxEnv()
}

func isTermuxEnv() bool {
	return getenv("TERMUX_VERSION") != "" ||
		getenv("PREFIX") == "/data/data/com.termux/files/usr"
}

// IsKaliLinux returns true if running on Kali Linux
func IsKaliLinux() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	return getDistro() == "kali"
}

func getDistro() string {
	return ""
}

func getenv(key string) string {
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
		return arch
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
