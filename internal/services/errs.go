package services

import "errors"

// бизнес-ошибки лучше хранить в пакете, где они рождаются
// ошибка должна нести смысл бизнес-правила, HTTP-статуc добавит - это уровень хендлера (409, 400, 500).
var (
	// 409 Conflict
	ErrEmailExists = errors.New("email alrady exists")
	
	ErrInvalidRole = errors.New("invalid role")

)
