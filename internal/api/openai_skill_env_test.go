package api

import (
	"net/http"
	"testing"
)

func TestExtractSkillEnvHeadersNormalizesAndFilters(t *testing.T) {
	h := http.Header{}
	h.Set("X-Fastclaw-Skill-Env-DC-API-KEY", " key ")
	h.Set("X-Fastclaw-Skill-Env-DC.SECRET_KEY", "secret")
	h.Set("X-Fastclaw-Skill-Env-FASTCLAW_STORAGE_DSN", "dsn")
	h.Set("X-Fastclaw-Skill-Env-1BAD", "bad")

	got := extractSkillEnvHeaders(h)
	if got["DC_API_KEY"] != "key" {
		t.Fatalf("DC_API_KEY = %q", got["DC_API_KEY"])
	}
	if got["DC_SECRET_KEY"] != "secret" {
		t.Fatalf("DC_SECRET_KEY = %q", got["DC_SECRET_KEY"])
	}
	if _, ok := got["FASTCLAW_STORAGE_DSN"]; ok {
		t.Fatal("blocked FASTCLAW env should not be extracted")
	}
	if _, ok := got["1BAD"]; ok {
		t.Fatal("invalid env key should not be extracted")
	}
}
