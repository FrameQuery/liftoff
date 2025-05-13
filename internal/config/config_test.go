// === internal/config/config_test.go ===
package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadServiceSpecificDefaults(t *testing.T) {
	// Set up a temporary directory for config
	tmpDir, err := ioutil.TempDir("", "liftoff_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Override os.TempDir via TMPDIR env
	os.Setenv("TMPDIR", tmpDir)

	// Prepare config file
	cfgPath := filepath.Join(tmpDir, "liftoff_config.json")
	configJSON := `{
      "services": {
        "svc1": {
          "project": "proj-123",
          "regions": ["r1","r2"],
          "percentages": [5,50,100],
          "intervals": [10,20]
        }
      }
    }`
	if err := ioutil.WriteFile(cfgPath, []byte(configJSON), 0644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	// Reset viper and set the target service
	viper.Reset()
	viper.Set("service", "svc1")

	// Call Load
	if err := Load(); err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	// Verify merged values
	if got := viper.GetString("project"); got != "proj-123" {
		t.Errorf("expected project=proj-123, got %q", got)
	}

	wantRegions := []string{"r1", "r2"}
	if got := viper.GetStringSlice("regions"); !equalStringSlice(got, wantRegions) {
		t.Errorf("expected regions=%v, got %v", wantRegions, got)
	}

	wantPcts := []int{5, 50, 100}
	if got := viper.GetIntSlice("percentages"); !equalIntSlice(got, wantPcts) {
		t.Errorf("expected percentages=%v, got %v", wantPcts, got)
	}

	wantInt := []int{10, 20}
	if got := viper.GetIntSlice("intervals"); !equalIntSlice(got, wantInt) {
		t.Errorf("expected intervals=%v, got %v", wantInt, got)
	}
}

func TestLoadNoConfig(t *testing.T) {
	// Empty temp dir
	tmpDir, err := ioutil.TempDir("", "liftoff_test_empty")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	os.Setenv("TMPDIR", tmpDir)

	// Remove any pre-existing config
	cfgPath := filepath.Join(tmpDir, "liftoff_config.json")
	os.Remove(cfgPath)

	viper.Reset()
	viper.Set("service", "nonexistent")
	// Should not error even if config missing
	if err := Load(); err != nil {
		t.Errorf("Load() error with no config: %v", err)
	}
}

// Helpers
func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
