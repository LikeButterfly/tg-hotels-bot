package states

// можно хранить состояния в mongodb

import (
	"log"
	"sync"
)

// Перечисление всех возможных состояний бота
type StateType int

const (
	// SelectCity
	WaitCityName StateType = iota
	SelectCity

	// SelectDates
	SelectDateIn
	SelectDateOut
	IsDateCorrect
)

// Управляет состояниями пользователей
type StateManager struct {
	mu     sync.RWMutex // Защита - чтобы избежать гонок данных в многопоточной среде
	states map[int64]StateType
}

// Создаёт менеджер состояний
func NewStateManager() *StateManager {
	return &StateManager{
		states: make(map[int64]StateType),
	}
}

// Устанавливает состояние для пользователя
func (sm *StateManager) SetState(userID int64, state StateType) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.states[userID] = state
	log.Printf("User %d switched to state: %d\n", userID, state)
}

// Получает текущее состояние пользователя
func (sm *StateManager) GetState(userID int64) (StateType, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	state, exists := sm.states[userID]
	return state, exists
}

// Сбрасывает состояние пользователя
func (sm *StateManager) ClearState(userID int64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.states, userID)
}
