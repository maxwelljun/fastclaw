package api

import (
	"net/http"
	"testing"

	"github.com/fastclaw-ai/fastclaw/internal/agent"
	"github.com/fastclaw-ai/fastclaw/internal/auth"
	"github.com/fastclaw-ai/fastclaw/internal/config"
)

type recordingResolver struct {
	seenUserID string
}

func (r *recordingResolver) UserSpaceFor(userID string) (*UserSpaceView, error) {
	r.seenUserID = userID
	return &UserSpaceView{}, nil
}

func (r *recordingResolver) LocalAgentManager() *agent.Manager { return nil }

func (r *recordingResolver) IsCloudMode() bool { return false }

func TestUserSpaceForAppUserResolvesOwnerSpace(t *testing.T) {
	resolver := &recordingResolver{}
	s := &Server{resolver: resolver}
	req, err := http.NewRequest(http.MethodPost, "/v1/chat/completions", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := auth.WithIdentity(req.Context(), auth.Identity{
		UserID:      "u_app",
		Role:        "app_user",
		AuthMethod:  "apikey",
		OwnerUserID: "u_owner",
	})
	req = req.WithContext(config.WithUserID(ctx, "u_app"))

	if _, err := s.userSpaceFor(req); err != nil {
		t.Fatal(err)
	}
	if resolver.seenUserID != "u_owner" {
		t.Fatalf("UserSpaceFor userID = %q, want owner", resolver.seenUserID)
	}
}
