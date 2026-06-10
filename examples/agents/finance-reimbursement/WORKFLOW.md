# WORKFLOW.md

## Workflow: Prepare Reimbursement

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
   - Check required fields.
   - Check receipt/invoice evidence.
   - Recompute totals.
   - Detect possible duplicates by amount, date, merchant, and invoice number.
   - Flag personal expenses, missing attendees, weekend/holiday anomalies, and policy exceptions.

5. Clarify
   - Ask only the questions needed to unblock submission.
   - Group questions by expense item.

6. Package
   - Produce expense table.
   - Produce evidence checklist.
   - Produce exception summary.
   - Produce final requester-facing submission draft.

7. Handoff
   - If ready, mark "Ready for finance review".
   - If not ready, mark "Needs requester action" with concrete missing items.

## Workflow: Finance Review

1. Load reimbursement packet and source evidence.
2. Verify totals, categories, invoice fields, approvals, and policy exceptions.
3. Produce findings grouped by severity:
   - Blocker
   - Needs clarification
   - Policy exception
   - Informational
4. Prepare a concise response to requester or manager.

## Default Expense Categories

- Travel - airfare
- Travel - hotel
- Travel - ground transportation
- Meals and entertainment
- Office supplies
- Software and subscriptions
- Training and conference
- Client expense
- Other
