package db

import (
	"database/sql"

	"github.com/freddysilber/nfl-looser-pool-api/models"
)

func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.OwnerId, &item.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	err := db.Conn.QueryRow(
		`INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`,
		item.Name,
		item.Description,
	).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	item.Id = id
	item.CreatedAt = createdAt
	return nil
}

func (db Database) GetItemById(itemId int) (models.Item, error) {
	item := models.Item{}
	row := db.Conn.QueryRow(
		`SELECT * FROM items WHERE id = $1;`,
		itemId,
	)
	switch err := row.Scan(&item.Id, &item.Name, &item.Description, &item.CreatedAt); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

func (db Database) DeleteItem(itemId int) error {
	_, err := db.Conn.Exec(`DELETE FROM items WHERE id = $1;`, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	err := db.Conn.QueryRow(
		`UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`,
		itemData.Name,
		itemData.Description,
		itemId,
	).Scan(&item.Id, &item.Name, &item.Description, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}
	return item, nil
}
