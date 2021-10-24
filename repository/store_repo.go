package repository

import "spos/stores/models"

func (r *repo) FindStoreByDomain(domain string) (*models.Store, error) {
	var (
		store models.Store
		err   error
	)
	query := `select id, name, domain from stores where domain = $1`

	err = r.db.Get(&store, query, domain)
	return &store, err
}

func (r *repo) FindStoreByName(name string) (*models.Store, error) {
	var (
		store models.Store
		err   error
	)
	query := `select id, name, domain from stores where LOWER(name) = $1`

	err = r.db.Get(&store, query, name)
	return &store, err
}

func (r *repo) InsertStore(store *models.Store) (*models.Store, error) {
	var (
		tx  = r.db.MustBegin()
		err error
	)

	query := `insert into stores
				(owner_id, ref_shop_id, ref_user_id, ref_shop_status, name,
				domain, logo, description, enabled, type, source, created_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
				returning id`

	err = tx.QueryRowx(
		query,
		store.GetOwnerID(),
		store.RefShopID,
		store.RefUserID,
		store.RefShopStatus,
		store.GetName(),
		store.GetDomain(),
		store.GetLogo(),
		store.Description,
		store.GetEnabled(),
		store.GetType(),
		store.GetSource(),
		store.GetCreatedAt(),
	).StructScan(store)

	if err != nil {
		tx.Rollback()
		return store, err
	}

	err = tx.Commit()
	return store, err
}
