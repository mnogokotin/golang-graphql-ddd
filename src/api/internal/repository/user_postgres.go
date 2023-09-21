package repository

import (
	"context"
)

type UsersRepo struct {
	db int8 //*postgres.Postgres
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(offersCollection),
	}
}

func (r *UserRepo) GetHistory(ctx context.Context) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.Source, &e.Destination, &e.Original, &e.User)
		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}
