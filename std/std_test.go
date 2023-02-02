package std

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/fmyxyz/tx/test"

	_ "github.com/go-sql-driver/mysql"
)

func TestWithTx(t *testing.T) {

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	//db = db.Debug()
	Register(db)

	//resetAll(db)
	test.Update88 = update88
	test.Update99 = update99
	test.Update = update
	for _, tt := range test.Tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Fc(context.Background())
			eq88 := getValId(db, 88) == 88
			eq99 := getValId(db, 99) == 99
			resetAll(db)
			if eq88 != tt.Eq88 || eq99 != tt.Eq99 || (err != nil) != tt.WantErr {
				t.Errorf("WithTx() {eq88 = %v, want %v} {eq99 = %v, want %v} {error = %v, want has error %v}", eq88, tt.Eq88, eq99, tt.Eq99, err, tt.WantErr)
			}
			if err != nil {
				t.Log(err)
			}
		})
	}

	//resetAll(db)

}

func openDB() (db *sql.DB, err error) {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getValId(db *sql.DB, id int) (val int) {
	dest := &test.Item{
		ID: id,
	}
	rows, err := db.Query("select id,qty from item where id=?", id)
	if err != nil {
		return 0
	}
	if rows.Next() {
		err := rows.Scan(&dest.ID, &dest.Qty)
		if err != nil {
			return 0
		}
	}
	return dest.Qty
}

func reset(db *sql.DB, id int) error {
	_, err := db.Exec("update item set qty=? where id=?", id+1, id)
	if err != nil {
		return err
	}
	return nil
}

func resetAll(db *sql.DB) error {
	log.Println("resetAll")
	reset(db, 88)
	reset(db, 99)
	return nil
}

func update(ctx context.Context, id, num int) error {
	db := FromContext(ctx)
	_, err := db.Exec("update item set qty=? where id=?", num, id)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func update88(ctx context.Context) error {
	db := FromContext(ctx)

	_, err := db.Exec("update item set qty=? where id=?", 88, 88)
	if err != nil {
		return err
	}
	return nil
}

func update99(ctx context.Context) error {
	db := FromContext(ctx)

	_, err := db.Exec("update item set qty=? where id=?", 99, 99)
	if err != nil {
		return err
	}
	return nil
}
