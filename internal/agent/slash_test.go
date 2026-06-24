package agent

import (
	"strings"
	"testing"

	"github.com/fastclaw-ai/fastclaw/internal/bus"
)

func TestSlashStartUsesDisplayName(t *testing.T) {
	a := &Agent{name: "agt_123", displayName: "DClaw"}
	res := a.handleSlashCommand(bus.InboundMessage{Text: "/start"})
	if !res.handled {
		t.Fatal("expected handled=true")
	}
	if !strings.Contains(res.reply, "I'm DClaw") {
		t.Fatalf("reply = %q, want display name", res.reply)
	}
	if strings.Contains(res.reply, "agt_123") {
		t.Fatalf("reply leaked internal agent id: %q", res.reply)
	}
	if strings.Contains(res.reply, "Just send me a message") {
		t.Fatalf("reply includes extra usage hint: %q", res.reply)
	}
}

func TestSlashStartFallsBackToAgentID(t *testing.T) {
	a := &Agent{name: "agt_123"}
	res := a.handleSlashCommand(bus.InboundMessage{Text: "/start"})
	if !res.handled {
		t.Fatal("expected handled=true")
	}
	if !strings.Contains(res.reply, "I'm agt_123") {
		t.Fatalf("reply = %q, want agent id fallback", res.reply)
	}
}

func TestSlashRequiresAdmin(t *testing.T) {
	tests := []struct {
		name     string
		cmd      string
		peerKind string
		want     bool
	}{
		{name: "new in dm", cmd: "/new", peerKind: "dm", want: false},
		{name: "reset in dm", cmd: "/reset", peerKind: "dm", want: false},
		{name: "new with legacy empty peer kind", cmd: "/new", want: false},
		{name: "new in group", cmd: "/new", peerKind: "group", want: true},
		{name: "reset in group", cmd: "/reset", peerKind: "group", want: true},
		{name: "model in dm", cmd: "/model", peerKind: "dm", want: true},
		{name: "read command", cmd: "/status", peerKind: "group", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := bus.InboundMessage{PeerKind: tt.peerKind}
			if got := slashRequiresAdmin(tt.cmd, msg); got != tt.want {
				t.Fatalf("slashRequiresAdmin(%q, peer=%q) = %v, want %v", tt.cmd, tt.peerKind, got, tt.want)
			}
		})
	}
}
