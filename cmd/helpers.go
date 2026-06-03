package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"os"
	"path/filepath"
)

func LoadKeep(keepRoot string) ([]Link, error) {
	var allLinks []Link

	keepFile := filepath.Join(keepRoot, "keep.yaml")
	data, err := os.ReadFile(keepFile)
	if err != nil {
		return nil, fmt.Errorf(" Failed to read keep.yaml: %w", err)
	}

	var keep Keep
	if err := yaml.Unmarshal(data, &keep); err != nil {
		return nil, fmt.Errorf(" Failed to parse keep.yaml: %w", err)
	}
	allLinks = append(allLinks, keep.Links...)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf(" Failed to get hostname: %w", err)
	}

	hSpecFile := filepath.Join(keepRoot, "hSpecs", hostname+".yaml")
	if _, err := os.Stat(hSpecFile); err == nil {
		hData, err := os.ReadFile(hSpecFile)
		if err != nil {
			return nil, fmt.Errorf(" Failed to read hSpec: %w", err)
		}

		var hSpec Keep
		if err := yaml.Unmarshal(hData, &hSpec); err != nil {
			return nil, fmt.Errorf(" Failed to parse hSpec: %w", err)
		}
		color.Green("󰌨 Applying hSpec for host: " + hostname)

		allLinks = append(allLinks, hSpec.Links...)
	}

	return allLinks, nil
}

func SaveState(path string, st State) error {
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return fmt.Errorf(" Failed to marshal state: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf(" Failed to write state file: %w", err)
	}

	return nil
}
