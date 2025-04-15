package account

import (
    "vocabulary/api/account"
    "vocabulary/internal/logic/users"
)

type ControllerV1 struct {
    users *users.Users
}

func NewV1() account.IAccountV1 {
    return &ControllerV1{
        users: users.New(),
    }
}