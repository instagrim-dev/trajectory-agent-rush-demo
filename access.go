package subscription

func HasAccess(paidUntil, now int64) bool { return now < paidUntil }

