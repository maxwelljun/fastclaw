---
name: reimbursement-intake
description: Convert reimbursement requests, receipt notes, and pasted expense details into a structured expense intake table with missing-field questions.
metadata:
  fastclaw:
    always: true
---

# Reimbursement Intake Skill

Use this skill when a user provides reimbursement details, receipts, pasted text, travel notes, or asks to prepare a claim.

## Inputs

- Free-form expense description.
- Receipt or invoice text.
- CSV or spreadsheet exports.
- Policy snippets supplied by the user.

## Output

Produce:

1. Expense intake table:
   - item_id
   - expense_date
   - merchant
   - category
   - amount
   - currency
   - tax_amount
   - invoice_number
   - payment_method
   - business_purpose
   - project_or_cost_center
   - evidence_source
   - confidence

2. Missing fields:
   - item_id
   - missing_field
   - why_needed
   - suggested_question

3. Submission readiness:
   - Ready for finance review
   - Needs requester action
   - Needs manager approval

## Procedure

1. Extract facts exactly as provided.
2. Do not infer merchant, amount, invoice number, or tax fields without evidence.
3. Normalize dates and currencies only when unambiguous.
4. If a value is inferred, mark confidence as "low" and add a clarification question.
5. Keep every question actionable and tied to an item_id.
