package db

import (
	"database/sql"
)

// GetCategories gibt eine Map von CategoryID zu CategoryName zurück.
func GetCategories(db *sql.DB) (map[int]string, error) {
	result := make(map[int]string)
	rows, err := db.Query("SELECT CategoryID, CategoryName FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		result[id] = name
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// GetProductsInCategory gibt eine Map von ProductID zu ProductName für eine bestimmte Kategorie zurück.
func GetProductsInCategory(db *sql.DB, categoryID int) (map[int]string, error) {
	result := make(map[int]string)
	stmt, err := db.Prepare("SELECT ProductID, ProductName FROM products WHERE CategoryID = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		result[id] = name
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTotalOrderQuantity gibt die Gesamtmenge der Bestellungen für ein bestimmtes Produkt zurück.
func GetTotalOrderQuantity(db *sql.DB, productID int) (int, error) {
	var total int

	// TODO Gesamtmenge aus der Tabelle OrderDetails ermitteln

	return total, nil
}
