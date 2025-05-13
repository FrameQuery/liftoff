package gcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run executes a gcloud command and streams output
func Run(args ...string) error {
	cmd := exec.Command("gcloud", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ServiceExists returns true if the Cloud Run service already exists in that region.
// ServiceExists returns true if a Cloud Run service with that name exists.
func ServiceExists(service, region, project string) bool {
	// Use `gcloud run services list` with a filter and a value-only format
	args := []string{
		"run", "services", "list",
		"--platform=managed",
		fmt.Sprintf("--region=%s", region),
		fmt.Sprintf("--project=%s", project),
		"--filter", fmt.Sprintf("metadata.name=%s", service),
	}
	cmd := exec.Command("gcloud", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	// If it exactly matches the service name, it exists
	return !strings.Contains(string(out), "Listed 0 items.")
}

// Deploy wraps `gcloud run deploy`, using --no-traffic only if the service already exists.
func Deploy(service, image, region, project string, envVars []string) error {
	fmt.Printf("🛰️  Deploying %s to %s (--no-traffic)\n", service, region)
	args := []string{"run", "deploy", service,
		"--image=" + image,
		"--region=" + region,
		"--platform=managed",
		"--project=" + project,
		"--no-allow-unauthenticated",
	}
	if len(envVars) > 0 {
		args = append(args, "--set-env-vars="+strings.Join(envVars, ","))
	}
	serviceExists := ServiceExists(service, region, project)
	if serviceExists {
		args = append(args, "--no-traffic")
	} else {
		fmt.Println("⚠️  Service does not exist, first deployment will be without --no-traffic")
	}

	return Run(args...)
}

// UpdateTraffic wraps `gcloud run services update-traffic`
func UpdateTraffic(service, region string, toRevisions string, project string) error {
	fmt.Printf("🚦 Setting traffic: %s -> %s in %s\n", service, toRevisions, region)
	return Run("run", "services", "update-traffic", service,
		fmt.Sprintf("--to-revisions=%s", toRevisions),
		fmt.Sprintf("--region=%s", region),
		"--platform=managed",
		fmt.Sprintf("--project=%s", project),
	)
}

// RunOutput is like Run but captures stdout (so you can parse it).
func RunOutput(args ...string) (string, error) {
	cmd := exec.Command("gcloud", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// GetServingRevision returns the name of the revision currently getting traffic.
func GetServingRevision(service, region, project string) (string, error) {
	args := []string{
		"run", "services", "describe", service,
		"--platform=managed",
		fmt.Sprintf("--region=%s", region),
		fmt.Sprintf("--project=%s", project),
		"--format=json(status.traffic)",
	}
	return RunOutput(args...)
}

func getRevisionWithHighestTraffic(jsonOutput string) (string, error) {
	var response struct {
		Status struct {
			Traffic []struct {
				Percent      int    `json:"percent"`
				RevisionName string `json:"revisionName"`
			} `json:"traffic"`
		} `json:"status"`
	}

	if err := json.Unmarshal([]byte(jsonOutput), &response); err != nil {
		return "", err
	}

	var maxPercent int
	var highestRevision string

	for _, traffic := range response.Status.Traffic {
		if traffic.Percent > maxPercent {
			maxPercent = traffic.Percent
			highestRevision = traffic.RevisionName
		}
	}

	return highestRevision, nil
}

// GetLatestRevision returns the name of the latest‐created revision for the service.
func GetLatestRevision(service, region, project string) (string, error) {
	args := []string{
		"run", "services", "describe", service,
		"--platform=managed",
		fmt.Sprintf("--region=%s", region),
		fmt.Sprintf("--project=%s", project),
		"--format=value(status.latestCreatedRevisionName)",
	}
	return RunOutput(args...)
}

// SplitTraffic splits traffic in one region between the latest and the current revision.
func SplitTraffic(service, region string, pct int, project string) error {
	data, err := GetServingRevision(service, region, project)
	if err != nil {
		return fmt.Errorf("failed to get serving revision: %w", err)
	}
	oldRev, err := getRevisionWithHighestTraffic(data)
	if err != nil {
		return fmt.Errorf("failed to parse JSON from revision: %w", err)
	}
	newRev, err := GetLatestRevision(service, region, project)
	if err != nil {
		return fmt.Errorf("failed to get latest revision: %w", err)
	}

	toRevisions := fmt.Sprintf("%s=%d,%s=%d", newRev, pct, oldRev, 100-pct)
	return UpdateTraffic(service, region, toRevisions, project)
}

// SplitTrafficAcrossRegions applies the same traffic split to multiple regions.
func SplitTrafficAcrossRegions(service string, regions []string, pct int, project string) error {
	for _, region := range regions {
		if err := SplitTraffic(service, region, pct, project); err != nil {
			return err
		}
	}
	return nil
}
