---
name: reimbursement-report
description: Generate finance-ready reimbursement packets, review checklists, exception summaries, and requester clarification drafts.
metadata:
  fastclaw:
    always: true
---

# Reimbursement Report Skill

Use this skill after intake and audit are complete.

## Report Types

Requester packet:
- Summary
- Expense table
- Missing items
- Clarification questions
- Submission draft

Finance review packet:
- Summary
- Expense table
- Evidence checklist
- Policy checks
- Duplicate checks
- Exceptions
- Recommended next action

Manager approval note:
- Business purpose
- Total amount
- Exceptions requiring approval
- Decision options

## Format Rules

- Use Markdown by default.
- Use CSV when the user asks for spreadsheet-ready output.
- Keep amounts in original currencies unless conversion is explicitly requested.
- Include "Assumptions" only when assumptions were made.
- Include "Not reviewed" for any area where the required source was unavailable.

## Standard Status Labels

- Ready for finance review
- Needs requester action
- Needs manager approval
- Needs finance decision
- Not enough information
