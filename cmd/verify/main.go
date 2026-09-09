package main

import (
	"fmt"

	"tour-promise-verifier/internal/evidence"
)

func main() {
	config := evidence.ArchiveConfig{Root: "data/evidence", RuleVersion: "promise-2026-01"}
	fmt.Printf("承诺核验归档目录：%s，规则：%s\n", config.Root, config.RuleVersion)
}
