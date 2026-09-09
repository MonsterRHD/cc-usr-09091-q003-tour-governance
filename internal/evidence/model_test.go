package evidence

import "testing"

func TestEvidenceRequiresDigest(t *testing.T) {
	e := Evidence{ID: "e-1", Digest: "sha256:abc"}
	if e.Digest == "" {
		t.Fatal("证据摘要不能为空")
	}
}
