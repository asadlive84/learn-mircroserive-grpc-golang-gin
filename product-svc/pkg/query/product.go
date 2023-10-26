package query

import "database/sql"

func (q *QueryInit) InsertProduct(u Product) error {
	const query = `INSERT INTO 
	product 
	(
		name, 
		description,
		created_by
	) 
	VALUES 
	(
		:name, 
		:description, 
		:created_by
	)`
	stmt, err := q.Db.PrepareNamed(query)
	if err != nil {
		return err

	}
	_, err = stmt.Exec(&u)

	if err != nil {
		return err
	}
	return nil
}

func (q *QueryInit) GetProductByID(id string) (*Product, error) {
	var product Product
	qc := "SELECT * FROM product WHERE id=$1"
	if err := q.Db.Get(&product, qc, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return &product, nil
}
