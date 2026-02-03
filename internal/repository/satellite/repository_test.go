package satellite

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/BigDwarf/testci/internal/model"
)

//func GetFixture(t *testing.T, name string) []byte {
//	dir, err := os.Getwd()
//	os.Open(path.Join(dir, t.Name(), ".fixture")
//}
//
//func GetGolden(t *testing.T, name string) []byte {
//	dir, err := os.Getwd()
//	os.Open(path.Join(dir, t.Name(), ".golden")
//}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tests := map[string]struct {
		input        *model.Satellite
		expectedSql  string
		expectedArgs []driver.Value
		expectedErr  error
	}{
		"insert moon": {
			input: &model.Satellite{
				Name: "moon",
			},
			expectedSql:  "INSERT INTO satellite",
			expectedArgs: []driver.Value{"moon"},
		},
		"insert europa": {
			input: &model.Satellite{
				Name: "europa",
			},
			expectedSql:  "INSERT INTO satellite",
			expectedArgs: []driver.Value{"europa"},
		},
	}

	satelliteRepository := NewRepository(db)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.expectedErr != nil {
				mock.ExpectQuery(tt.expectedSql).WillReturnError(tt.expectedErr)
			}

			mock.ExpectExec(tt.expectedSql).
				WithArgs(tt.expectedArgs...).WillReturnResult(sqlmock.NewResult(1, 1))

			err = satelliteRepository.Create(t.Context(), model.Satellite{Name: tt.input.Name})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when inserting satellite", err)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})

	}
}
