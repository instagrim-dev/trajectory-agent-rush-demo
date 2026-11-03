package subscription

// HasAccess implements the owner-selected cancellation policy.
func HasAccess(paidUntil, now int64) bool { return now < paidUntil }
