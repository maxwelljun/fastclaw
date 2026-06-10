# TOOLS.md

## Tool Rules

- Use file tools to inspect uploaded receipts, CSV exports, policy documents, and prepared reimbursement packets.
- Use sandbox execution for calculations, CSV normalization, duplicate detection, and report generation.
- Use skills before writing one-off scripts when an available skill matches the task.
- Use web search only when the user asks for public information or when a public source is required. Do not search for internal policy unless the user provides the URL and confirms access is appropriate.

## Data Handling

- Do not expose full bank card numbers, national IDs, tax IDs, home addresses, or personal phone numbers in summaries.
- Redact sensitive identifiers in generated reports unless the user explicitly asks for an internal finance-ready artifact.
- Never fabricate missing invoice fields.
- Keep original values and normalized values separate.

## Calculation Rules

- Preserve currency codes.
- Do not convert currencies unless an exchange rate source and date are provided.
- If rounding is needed, state the rule used.
- Recompute totals from line items and flag mismatches.

## Output Artifacts

Preferred artifacts:
- `reimbursement_packet.md`
- `expense_table.csv`
- `finance_review_checklist.md`
- `exceptions.md`
- `clarification_questions.md`
