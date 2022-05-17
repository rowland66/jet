//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Film = newFilmTable("dvds", "film", "")

type filmTable struct {
	postgres.Table

	//Columns
	FilmID          postgres.ColumnInteger
	Title           postgres.ColumnString
	Description     postgres.ColumnString
	ReleaseYear     postgres.ColumnInteger
	LanguageID      postgres.ColumnInteger
	RentalDuration  postgres.ColumnInteger
	RentalRate      postgres.ColumnFloat
	Length          postgres.ColumnInteger
	ReplacementCost postgres.ColumnFloat
	Rating          postgres.ColumnString
	LastUpdate      postgres.ColumnTimestamp
	SpecialFeatures postgres.ColumnString
	Fulltext        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FilmTable struct {
	filmTable

	EXCLUDED filmTable
}

// AS creates new FilmTable with assigned alias
func (a FilmTable) AS(alias string) *FilmTable {
	return newFilmTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FilmTable with assigned schema name
func (a FilmTable) FromSchema(schemaName string) *FilmTable {
	return newFilmTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FilmTable with assigned table prefix
func (a FilmTable) WithPrefix(prefix string) *FilmTable {
	return newFilmTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FilmTable with assigned table suffix
func (a FilmTable) WithSuffix(suffix string) *FilmTable {
	return newFilmTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFilmTable(schemaName, tableName, alias string) *FilmTable {
	return &FilmTable{
		filmTable: newFilmTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newFilmTableImpl("", "excluded", ""),
	}
}

func newFilmTableImpl(schemaName, tableName, alias string) filmTable {
	var (
		FilmIDColumn          = postgres.IntegerColumn("film_id")
		TitleColumn           = postgres.StringColumn("title")
		DescriptionColumn     = postgres.StringColumn("description")
		ReleaseYearColumn     = postgres.IntegerColumn("release_year")
		LanguageIDColumn      = postgres.IntegerColumn("language_id")
		RentalDurationColumn  = postgres.IntegerColumn("rental_duration")
		RentalRateColumn      = postgres.FloatColumn("rental_rate")
		LengthColumn          = postgres.IntegerColumn("length")
		ReplacementCostColumn = postgres.FloatColumn("replacement_cost")
		RatingColumn          = postgres.StringColumn("rating")
		LastUpdateColumn      = postgres.TimestampColumn("last_update")
		SpecialFeaturesColumn = postgres.StringColumn("special_features")
		FulltextColumn        = postgres.StringColumn("fulltext")
		allColumns            = postgres.ColumnList{FilmIDColumn, TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, LastUpdateColumn, SpecialFeaturesColumn, FulltextColumn}
		mutableColumns        = postgres.ColumnList{TitleColumn, DescriptionColumn, ReleaseYearColumn, LanguageIDColumn, RentalDurationColumn, RentalRateColumn, LengthColumn, ReplacementCostColumn, RatingColumn, LastUpdateColumn, SpecialFeaturesColumn, FulltextColumn}
	)

	return filmTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		FilmID:          FilmIDColumn,
		Title:           TitleColumn,
		Description:     DescriptionColumn,
		ReleaseYear:     ReleaseYearColumn,
		LanguageID:      LanguageIDColumn,
		RentalDuration:  RentalDurationColumn,
		RentalRate:      RentalRateColumn,
		Length:          LengthColumn,
		ReplacementCost: ReplacementCostColumn,
		Rating:          RatingColumn,
		LastUpdate:      LastUpdateColumn,
		SpecialFeatures: SpecialFeaturesColumn,
		Fulltext:        FulltextColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
