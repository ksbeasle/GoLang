package mysql

import (
	"database/sql"
	"io/ioutil"
	"testing"
)

func createTestDB(t *testing.T) (*sql.DB, func()) {
	//multiStatements set to true since we are calling multiple SQL statements in the setup.sql/teardown.sql files
	db, err := sql.Open("mysql", "test_web1:pass@/test_videogames?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	//Lets read the script
	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	//Lets execute the setup.sql file
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		//Lets read the teardown.sql script now
		script, err := ioutil.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		//Exectute the script
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	}

}
