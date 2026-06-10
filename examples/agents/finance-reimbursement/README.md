# Finance Reimbursement Agent

This example defines a finance reimbursement agent that helps employees prepare compliant reimbursement packages and helps finance reviewers check completeness, policy alignment, and exception handling.

## Files

- `agent.json` - Suggested FastClaw agent configuration.
- `IDENTITY.md` - Agent role, operating scope, and escalation boundaries.
- `SOUL.md` - Tone, principles, and decision posture.
- `AGENTS.md` - Collaboration protocol for requester, manager, finance, and approver personas.
- `TOOLS.md` - Tool usage rules and safety constraints.
- `WORKFLOW.md` - End-to-end reimbursement workflow.
- `USER.md` - Per-user profile template.
- `MEMORY.md` - Long-term memory template.
- `HEARTBEAT.md` - Optional periodic follow-up checks.
- `skills/` - Agent-private skills for reimbursement work.

## Suggested Install

Create or select an agent, then upload the files:

```bash
fastclaw agents init finance-reimbursement
fastclaw agents files put finance-reimbursement IDENTITY.md examples/agents/finance-reimbursement/IDENTITY.md
fastclaw agents files put finance-reimbursement SOUL.md examples/agents/finance-reimbursement/SOUL.md
fastclaw agents files put finance-reimbursement AGENTS.md examples/agents/finance-reimbursement/AGENTS.md
fastclaw agents files put finance-reimbursement TOOLS.md examples/agents/finance-reimbursement/TOOLS.md
fastclaw agents files put finance-reimbursement HEARTBEAT.md examples/agents/finance-reimbursement/HEARTBEAT.md
fastclaw agents files put finance-reimbursement agent.json examples/agents/finance-reimbursement/agent.json
```

Copy `skills/*` into the agent's private skills directory if you want the skills to be available only to this agent.
