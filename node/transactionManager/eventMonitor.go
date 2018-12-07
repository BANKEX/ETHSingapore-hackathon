package transactionManager

type EventMonitor struct {
	transactionManager *TransactionManager
}

func NewEventMonitor(m *TransactionManager) *EventMonitor {
	result := EventMonitor{
		transactionManager: m,
	}

	// start watchers here

	return &result
}

// todo monitor deposit events, forward to transaction manager
// todo monitor withdraw events, forward to transaction manager
// todo if we need to send some challenges from the operator, this is the place to do it
