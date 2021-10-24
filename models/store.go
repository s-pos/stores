package models

import "time"

type Store struct {
	ID            int        `db:"id"`
	OwnerID       int        `db:"owner_id"`
	RefShopID     *int       `db:"ref_shop_id"`
	RefUserID     *int       `db:"ref_user_id"`
	RefShopStatus *string    `db:"ref_shop_status"`
	Name          string     `db:"name"`
	Domain        *string    `db:"domain"`
	Logo          string     `db:"logo"`
	Description   *string    `db:"description"`
	Enabled       bool       `db:"enabled"`
	Type          string     `db:"type"`
	Source        string     `db:"source"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	Deleted       bool       `db:"deleted"`
}

func NewStore() *Store {
	// let time now is UTC +0
	now := time.Now().UTC()

	return &Store{
		Enabled:   enabled,
		Type:      Offline,
		Source:    Offline,
		CreatedAt: now,
	}
}

func (s *Store) GetId() int {
	return s.ID
}

func (s *Store) SetOwnerID(ownerId int) {
	s.OwnerID = ownerId
}

func (s *Store) GetOwnerID() int {
	return s.OwnerID
}

func (s *Store) SetRefShopID(refShopId int) {
	s.RefShopID = &refShopId
}

func (s *Store) GetRefShopID() int {
	if s.RefShopID == nil {
		return defaultInt
	}
	return *s.RefShopID
}

func (s *Store) SetRefUserID(refUserId int) {
	s.RefUserID = &refUserId
}

func (s *Store) GetRefUserID() int {
	if s.RefUserID == nil {
		return defaultInt
	}
	return *s.RefUserID
}

func (s *Store) SetRefShopStatus(refShopStatus string) {
	s.RefShopStatus = &refShopStatus
}

func (s *Store) GetRefShopStatus() string {
	if s.RefShopStatus == nil {
		return defaultString
	}
	return *s.RefShopStatus
}

func (s *Store) SetName(name string) {
	s.Name = name
}

func (s *Store) GetName() string {
	return s.Name
}

func (s *Store) SetDomain(domain string) {
	s.Domain = &domain
}

func (s *Store) GetDomain() string {
	if s.Domain == nil {
		return defaultString
	}
	return *s.Domain
}

func (s *Store) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Store) GetLogo() string {
	return s.Logo
}

func (s *Store) SetDescription(description string) {
	s.Description = &description
}

func (s *Store) GetDescription() string {
	if s.Description == nil {
		return defaultString
	}
	return *s.Description
}

func (s *Store) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

func (s *Store) GetEnabled() bool {
	return s.Enabled
}

func (s *Store) IsEnabled() bool {
	return s.GetEnabled()
}

func (s *Store) SetType(types string) {
	s.Type = types
}

func (s *Store) GetType() string {
	return s.Type
}

func (s *Store) SetSource(sourcee string) {
	s.Source = sourcee
}

func (s *Store) GetSource() string {
	return s.Source
}

func (s *Store) SetCreatedAt(CreatedAt time.Time) {
	CreatedAt = CreatedAt.UTC()
	s.CreatedAt = CreatedAt
}

func (s *Store) GetCreatedAt() time.Time {
	return convertTimezone(s.CreatedAt)
}

func (s *Store) SetUpdatedAt(UpdatedAt time.Time) {
	UpdatedAt = UpdatedAt.UTC()
	s.UpdatedAt = &UpdatedAt
}

func (s *Store) GetUpdatedAt() time.Time {
	if s.UpdatedAt == nil {
		return time.Time{}
	}

	updatedAt := convertTimezone(*s.UpdatedAt)
	return updatedAt
}

func (s *Store) SetDeleted(deleted bool) {
	s.Deleted = deleted
}

func (s *Store) GetDeleted() bool {
	return s.Deleted
}
