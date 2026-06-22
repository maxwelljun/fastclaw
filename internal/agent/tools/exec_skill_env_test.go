package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveSkillEnvMergesRequestEnvForDeclaredRequirements(t *testing.T) {
	skillsDir := filepath.Join(t.TempDir(), "skills")
	skillDir := filepath.Join(skillsDir, "deepcoin-trade")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
name: deepcoin-trade
metadata:
  openclaw:
    requires:
      env: ["DC_API_KEY", "DC_SECRET_KEY"]
---

Run.`
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	got := resolveSkillEnv(
		"python "+filepath.Join(skillDir, "main.py"),
		func(skillName string) map[string]string {
			if skillName != "deepcoin-trade" {
				t.Fatalf("skillName = %q", skillName)
			}
			return map[string]string{
				"DC_API_KEY":   "from-config",
				"OTHER_CONFIG": "keep",
			}
		},
		map[string]string{
			"DC_API_KEY":    "from-header",
			"DC_SECRET_KEY": "secret-from-header",
			"UNDECLARED":    "drop",
		},
		[]string{skillsDir},
	)

	if got["DC_API_KEY"] != "from-header" {
		t.Fatalf("request env should override configured env, got %q", got["DC_API_KEY"])
	}
	if got["DC_SECRET_KEY"] != "secret-from-header" {
		t.Fatalf("declared request env missing, got %q", got["DC_SECRET_KEY"])
	}
	if got["OTHER_CONFIG"] != "keep" {
		t.Fatalf("configured env should be preserved, got %q", got["OTHER_CONFIG"])
	}
	if _, ok := got["UNDECLARED"]; ok {
		t.Fatal("undeclared request env should not be injected")
	}
}
