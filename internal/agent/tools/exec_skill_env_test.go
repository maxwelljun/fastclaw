package tools

import (
	"context"
	"encoding/json"
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

func TestResolveRegistrySkillEnvUsesLoadedSkillForPlainCommand(t *testing.T) {
	skillsDir := filepath.Join(t.TempDir(), "skills")
	skillDir := filepath.Join(skillsDir, "deepcoin-portfolio")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
name: deepcoin-portfolio
metadata:
  openclaw:
    requires:
      env: ["DC_API_KEY", "DC_SECRET_KEY", "DC_PASSPHRASE"]
---

Run.`
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	r := NewRegistry(t.TempDir(), t.TempDir())
	r.SetRequestSkillEnv(map[string]string{
		"DC_API_KEY":    "key",
		"DC_SECRET_KEY": "secret",
		"DC_PASSPHRASE": "pass",
		"UNDECLARED":    "drop",
	})
	RegisterLoadSkill(r, []string{skillsDir})
	rawArgs, err := json.Marshal(map[string]string{"name": "deepcoin-portfolio"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := r.GetFunc("load_skill")(context.Background(), rawArgs); err != nil {
		t.Fatal(err)
	}

	got := resolveRegistrySkillEnv("echo $DC_API_KEY", r, nil, []string{skillsDir})
	if got["DC_API_KEY"] != "key" || got["DC_SECRET_KEY"] != "secret" || got["DC_PASSPHRASE"] != "pass" {
		t.Fatalf("declared request env not injected for loaded skill: %#v", got)
	}
	if _, ok := got["UNDECLARED"]; ok {
		t.Fatal("undeclared request env should not be injected")
	}
}

func TestResolveRegistrySkillEnvInjectsReferencedDeclaredRequestEnv(t *testing.T) {
	skillsDir := filepath.Join(t.TempDir(), "skills")
	skillDir := filepath.Join(skillsDir, "deepcoin-portfolio")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
name: deepcoin-portfolio
metadata:
  openclaw:
    requires:
      env: ["DC_API_KEY", "DC_SECRET_KEY", "DC_PASSPHRASE"]
---

Run.`
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	r := NewRegistry(t.TempDir(), t.TempDir())
	r.SetRequestSkillEnv(map[string]string{
		"DC_API_KEY":    "key",
		"DC_SECRET_KEY": "secret",
		"DC_PASSPHRASE": "pass",
		"UNDECLARED":    "drop",
	})

	got := resolveRegistrySkillEnv(
		`python3 - <<'EOF'
import os
print(os.environ["DC_API_KEY"], os.environ["DC_SECRET_KEY"])
EOF`,
		r,
		nil,
		[]string{skillsDir},
	)
	if got["DC_API_KEY"] != "key" || got["DC_SECRET_KEY"] != "secret" {
		t.Fatalf("referenced declared request env not injected: %#v", got)
	}
	if _, ok := got["DC_PASSPHRASE"]; ok {
		t.Fatal("declared but unreferenced env should not be injected by command fallback")
	}
	if _, ok := got["UNDECLARED"]; ok {
		t.Fatal("undeclared request env should not be injected")
	}
}

func TestResolveRegistrySkillEnvInjectsDeclaredEnvForWorkspaceScript(t *testing.T) {
	skillsDir := filepath.Join(t.TempDir(), "skills")
	skillDir := filepath.Join(skillsDir, "deepcoin-portfolio")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
name: deepcoin-portfolio
metadata:
  openclaw:
    requires:
      env: ["DC_API_KEY", "DC_SECRET_KEY", "DC_PASSPHRASE"]
---

Run.`
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	r := NewRegistry(t.TempDir(), t.TempDir())
	r.SetRequestSkillEnv(map[string]string{
		"DC_API_KEY":    "key",
		"DC_SECRET_KEY": "secret",
		"DC_PASSPHRASE": "pass",
		"UNDECLARED":    "drop",
	})

	got := resolveRegistrySkillEnv("python3 /workspace/query_balance.py", r, nil, []string{skillsDir})
	if got["DC_API_KEY"] != "key" || got["DC_SECRET_KEY"] != "secret" || got["DC_PASSPHRASE"] != "pass" {
		t.Fatalf("declared request env not injected for workspace script: %#v", got)
	}
	if _, ok := got["UNDECLARED"]; ok {
		t.Fatal("undeclared request env should not be injected")
	}
}
