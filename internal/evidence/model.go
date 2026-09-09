package evidence

import "time"

type ArchiveConfig struct {
	Root        string
	RuleVersion string
}

type PromiseSnapshot struct {
	GroupID    string
	Text       string
	CapturedAt time.Time
	Digest     string
}

type Evidence struct {
	ID         string
	GroupID    string
	Kind       string
	ObservedAt time.Time
	Digest     string
}

type VerificationReport struct {
	GroupID     string
	RuleVersion string
	Mismatches  []string
	EvidenceIDs []string
}
