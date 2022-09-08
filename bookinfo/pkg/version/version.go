package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// Info contains versioning information.
type Info struct {
	Module       string `json:"module"`
	Version      string `json:"version"`
	Branch       string `json:"branch"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// Pretty returns a pretty output representation of Info
func (info Info) Pretty() string {
	str, _ := json.MarshalIndent(info, "", "    ")
	return string(str)
}

// String returns the marshalled json string of Info
func (info Info) String() string {
	str, _ := json.Marshal(info)
	return string(str)
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in version/var.go
	return Info{
		Module:       module,
		Version:      version,
		Branch:       branch,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
