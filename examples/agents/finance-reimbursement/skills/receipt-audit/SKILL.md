---
name: receipt-audit
description: Audit receipts and reimbursement line items for completeness, duplicate risk, policy exceptions, and finance review blockers.
metadata:
  fastclaw:
    always: true
---

# Receipt Audit Skill

Use this skill when reviewing reimbursement evidence before submission or finance approval.

## Checks

Required evidence:
- Receipt or invoice is present.
- Merchant is present.
- Date is present.
- Amount and currency are present.
- Business purpose is present.
- Payment proof is present when required.
- Attendees are listed for meals or entertainment.
- Approval is present for exceptions or threshold breaches.

Duplicate risk:
- Same merchant + same date + same amount.
- Same invoice number.
- Multiple claims for one trip segment.
- Split expenses that appear to bypass approval thresholds.

Policy risk:
- Missing business purpose.
- Personal or mixed personal/business spend.
- Weekend, holiday, or late-night anomaly.
- Expense outside allowed category.
- Missing pre-approval.
- Amount exceeds policy threshold.

## Severity

- Blocker: cannot submit or approve without correction.
- Needs clarification: likely valid but missing explanation.
- Policy exception: requires authorized approval.
- Informational: finance should be aware but not blocked.

## Output

Return findings as a table:

- severity
- item_id
- issue
- evidence
- recommended_action
- owner

Also provide a short overall readiness status.
