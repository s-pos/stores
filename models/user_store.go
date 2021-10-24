package models

import "time"

type UserStore struct {
	ID        int        `db:"id"`
	UserID    int        `db:"user_id"`
	RoleID    int        `db:"role_id"`
	StoreID   int        `db:"store_id"`
	Enabled   bool       `db:"enabled"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Deleted   bool       `db:"deleted"`
}

func NewUserStore() *UserStore {
	now := time.Now().UTC()

	return &UserStore{
		RoleID:    1,
		Enabled:   enabled,
		CreatedAt: now,
	}
}

func (us *UserStore) GetID() int {
	return us.ID
}

func (us *UserStore) SetUserID(userId int) {
	us.UserID = userId
}

func (us *UserStore) GetUserID() int {
	return us.UserID
}

func (us *UserStore) SetRoleID(roleId int) {
	us.RoleID = roleId
}

func (us *UserStore) GetRoleID() int {
	return us.RoleID
}

func (us *UserStore) SetStoreID(storeId int) {
	us.StoreID = storeId
}

func (us *UserStore) GetStoreID() int {
	return us.StoreID
}

func (us *UserStore) SetEnabled(enabled bool) {
	us.Enabled = enabled
}

func (us *UserStore) GetEnabled() bool {
	return us.Enabled
}

func (us *UserStore) IsEnabled() bool {
	return us.GetEnabled()
}

func (us *UserStore) SetCreatedAt(createdAt time.Time) {
	createdAt = createdAt.UTC()
	us.CreatedAt = createdAt
}

func (us *UserStore) GetCreatedAt() time.Time {
	return us.CreatedAt
}

func (us *UserStore) SetUpdatedAt(updatedAt time.Time) {
	updatedAt = updatedAt.UTC()
	us.UpdatedAt = &updatedAt
}

func (us *UserStore) GetUpdatedAt() time.Time {
	if us.UpdatedAt == nil {
		return time.Time{}
	}

	updatedAt := convertTimezone(*us.UpdatedAt)
	return updatedAt
}

func (us *UserStore) SetDeleted(deleted bool) {
	us.Deleted = deleted
}

func (us *UserStore) GetDeleted() bool {
	return us.Deleted
}
