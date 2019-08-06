package mysql

import (
	"github.com/go-jet/jet/internal/jet"
	"time"
)

var Bool = jet.Bool
var Int = jet.Int
var Float = jet.Float
var String = jet.String

var Date = func(year int, month time.Month, day int) DateExpression {
	return CAST(jet.Date(year, month, day)).AS_DATE()
}

var DateT = func(t time.Time) DateExpression {
	return CAST(jet.DateT(t)).AS_DATE()
}
var Time = func(hour, minute, second int, milliseconds ...int) TimeExpression {
	return CAST(jet.Time(hour, minute, second, milliseconds...)).AS_TIME()
}

var TimeT = func(t time.Time) TimeExpression {
	return CAST(jet.TimeT(t)).AS_TIME()
}
var DateTime = func(year int, month time.Month, day, hour, minute, second int, milliseconds ...int) DateTimeExpression {
	return CAST(jet.Timestamp(year, month, day, hour, minute, second, milliseconds...)).AS_DATETIME()
}

var DateTimeT = func(t time.Time) DateTimeExpression {
	return CAST(jet.TimestampT(t)).AS_DATETIME()
}
var Timestamp = func(year int, month time.Month, day, hour, minute, second int, milliseconds ...int) TimestampExpression {
	return CAST(jet.Timestamp(year, month, day, hour, minute, second, milliseconds...)).AS_TIMESTAMP()
}
var TimestampT = func(t time.Time) TimestampExpression {
	return CAST(jet.TimestampT(t)).AS_TIMESTAMP()
}
