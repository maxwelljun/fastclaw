package store

import (
	"context"
	"testing"
)

func TestDeleteAgentRemovesCurrentSchemaConfigRows(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	ctx := context.Background()

	const ownerID = "u_owner"
	const agentID = "agt_delete"
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO agents (id, user_id, name, config) VALUES (?, ?, 'delete-me', '{}')`,
		agentID, ownerID); err != nil {
		t.Fatalf("seed agent: %v", err)
	}
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO agent_files (agent_id, user_id, filename, content) VALUES (?, ?, 'SOUL.md', 'x')`,
		agentID, ownerID); err != nil {
		t.Fatalf("seed agent file: %v", err)
	}
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO sessions (user_id, agent_id, session_key) VALUES (?, ?, 's1')`,
		ownerID, agentID); err != nil {
		t.Fatalf("seed session: %v", err)
	}
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO cron_jobs (id, user_id, agent_id, name, schedule, message, channel, chat_id)
		 VALUES ('cj_delete', ?, ?, 'job', '* * * * *', 'msg', 'web', 'chat')`,
		ownerID, agentID); err != nil {
		t.Fatalf("seed cron job: %v", err)
	}
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO configs (id, kind, scope, scope_id, name, data)
		 VALUES ('cfg_agent', 'setting', 'agent', ?, 'agents.defaults', '{}')`,
		agentID); err != nil {
		t.Fatalf("seed agent config: %v", err)
	}
	if _, err := db.db.ExecContext(ctx,
		`INSERT INTO configs (id, kind, scope, scope_id, name, data)
		 VALUES ('cfg_system', 'setting', 'system', '', 'agents.defaults', '{}')`); err != nil {
		t.Fatalf("seed system config: %v", err)
	}

	if err := db.DeleteAgent(ctx, agentID); err != nil {
		t.Fatalf("DeleteAgent: %v", err)
	}

	assertCount := func(table, where string, args ...any) {
		t.Helper()
		var n int
		if err := db.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM `+table+` WHERE `+where, args...).Scan(&n); err != nil {
			t.Fatalf("count %s: %v", table, err)
		}
		if n != 0 {
			t.Fatalf("%s rows left = %d", table, n)
		}
	}
	assertCount("agents", "id = ?", agentID)
	assertCount("agent_files", "agent_id = ?", agentID)
	assertCount("sessions", "agent_id = ?", agentID)
	assertCount("cron_jobs", "agent_id = ?", agentID)
	assertCount("configs", "scope = 'agent' AND scope_id = ?", agentID)

	var systemConfigs int
	if err := db.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM configs WHERE id = 'cfg_system'`).Scan(&systemConfigs); err != nil {
		t.Fatalf("count system config: %v", err)
	}
	if systemConfigs != 1 {
		t.Fatalf("system config count = %d, want 1", systemConfigs)
	}
}
