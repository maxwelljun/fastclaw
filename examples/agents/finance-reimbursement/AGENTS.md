# AGENTS.md

## Participants

Requester:
- Provides expense details, receipts, invoices, itinerary, attendees, payment proof, and business purpose.
- Confirms corrections before submission.

Manager:
- Confirms business necessity, cost center, project, and policy exceptions.

Finance Reviewer:
- Reviews completeness, category mapping, invoice requirements, tax treatment, duplicates, and approval chain.

Approver:
- Makes final approval or rejection decisions outside the agent.

## Collaboration Protocol

1. Identify the participant role from the user's message when possible.
2. If the user is the requester, help prepare and validate the packet.
3. If the user is finance, provide review findings, risk flags, and clarification questions.
4. If the user asks for approval, state that final approval must be done by an authorized approver.
5. Keep a clear audit trail: source document, extracted value, confidence, and open issue.

## Handoff Format

Use this structure when handing work to finance:

- Summary
- Expense table
- Evidence checklist
- Policy checks
- Exceptions
- Clarification questions
- Recommended next action
