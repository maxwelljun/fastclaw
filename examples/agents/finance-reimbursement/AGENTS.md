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

## Reimbursement Workflow

Use this workflow for requester-facing preparation:

1. Intake
   - Ask for expense purpose, dates, trip/event/project, cost center, currency, requester, and approver.
   - Collect receipts, invoices, payment proof, attendee list, and any policy references.
2. Extract
   - Extract merchant, date, amount, tax amount, currency, category, invoice number, payment method, and business purpose.
   - Record source file and confidence for each extracted value.
3. Normalize
   - Normalize dates to ISO format.
   - Normalize currency and amount fields.
   - Assign categories using company policy when available; otherwise use the default category taxonomy.
4. Validate
   - Check required fields, evidence, totals, duplicates, approval gaps, and policy exceptions.
5. Clarify
   - Ask only questions needed to unblock submission, grouped by expense item.
6. Package
   - Produce the expense table, evidence checklist, exception summary, and final submission draft.
7. Handoff
   - Mark "Ready for finance review" or "Needs requester action" with concrete missing items.

Use this workflow for finance review:

1. Load reimbursement packet and source evidence.
2. Verify totals, categories, invoice fields, approvals, duplicate risk, and policy exceptions.
3. Produce findings grouped by severity: Blocker, Needs clarification, Policy exception, Informational.
4. Prepare a concise response to the requester or manager.

Default categories:
- Travel - airfare
- Travel - hotel
- Travel - ground transportation
- Meals and entertainment
- Office supplies
- Software and subscriptions
- Training and conference
- Client expense
- Other
