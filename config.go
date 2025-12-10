package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	UseMusic bool `json:"use_music"`
	UseLeaf  bool `json:"use_leaf"`
	Size     int  `json:"speed"`
}

const appDirName = "happy-go-christmas"
const configFileName = "config.json"

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appDirName, configFileName), nil
}

func CreateConfig() Config {
	print(Bold)
	var answer string
	var useMusic bool
	var useLeaf bool
	var size int

	fmt.Print("Use Music? (y/n): ")
	fmt.Scan(&answer)
	if answer != "n" {
		useMusic = true
	}

	fmt.Print("Use Leaf ( ⣾⣟⣽⣿⣯⣷ ) instead stars( * )? (y/n): ")
	fmt.Scan(&answer)
	if answer != "n" {
		useLeaf = true
	}

	for size < 5 {
		fmt.Print("Write the size of the tree(more than 5): ")
		fmt.Scan(&size)
	}

	err := SaveConfig(useMusic, useLeaf, size)
	if err != nil {
		path, _ := getConfigPath()
		fmt.Printf("Warning: could not save config to %s: %v\n", path, err)
	}

	return Config{
		UseMusic: useMusic,
		UseLeaf:  useLeaf,
		Size:     size,
	}
}

func LoadConfig() (Config, bool) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, false
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, false
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, false
	}

	return cfg, true
}

func SaveConfig(useMusic bool, useLeaf bool, size int) error {
	cfg := Config{
		UseMusic: useMusic,
		UseLeaf:  useLeaf,
		Size:     size,
	}

	path, err := getConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}

func handleConfig(cfg Config) {
	lenTree = cfg.Size
	marginBottom = lenTree / 5
	barkHight = lenTree / 10
	treeHight = (lenTree + 1) / 2
	marginSide = marginBottom * 7
	lenSide = treeHight + marginBottom + barkHight
	lenRow = marginSide + lenTree

	if cfg.UseMusic {
		playSound()
	}
}
