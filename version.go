package sapera

import (
	"fmt"
)

// GoSaperaVersion of this package, for display purposes.
// Change this variable on a new package release.
const GoSaperaVersion = "0.1.0"

// Version returns the current Golang package version.
func Version() string {
	return GoSaperaVersion
}

// SDKVersion returns the current Sapera SDK version.
func SDKVersion() string {
	info := GetVersionInfo()
	defer info.Delete()

	return fmt.Sprintf("%d.%d.%d.%d", info.GetMajor(), info.GetMinor(), info.GetRevision(), info.GetBuild())
}
