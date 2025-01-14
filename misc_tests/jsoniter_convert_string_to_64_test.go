package misc_tests

import (
	"testing"

	jsoniter "github.com/sandwich-go/jsoniter"
	"github.com/stretchr/testify/require"
)

func Test_convert_read_uint64_invalid(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, ",")
	iter.ReadUint64()
	should.NotNil(iter.Error)
}

func Test_convert_read_int64_with_quote(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"123456789098765\"")
	v := iter.ReadUint64()
	should.Nil(iter.Error)
	should.Equal(v, uint64(123456789098765))

	iter = jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"123456789098765L\"")
	v = iter.ReadUint64()
	should.Nil(iter.Error)
	should.Equal(v, uint64(123456789098765))

	iter = jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"123456789098765\"")
	v2 := iter.ReadInt64()
	should.Nil(iter.Error)
	should.Equal(v2, int64(123456789098765))

	iter = jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"123456789098765L\"")
	v2 = iter.ReadInt64()
	should.Nil(iter.Error)
	should.Equal(v2, int64(123456789098765))

	iter = jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"-123456789098765\"")
	v2 = iter.ReadInt64()
	should.Nil(iter.Error)
	should.Equal(v2, int64(-123456789098765))

	iter = jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "\"-123456789098765L\"")
	v2 = iter.ReadInt64()
	should.Nil(iter.Error)
	should.Equal(v2, int64(-123456789098765))
}

func Test_read_int64_array_with_quote(t *testing.T) {
	should := require.New(t)
	input := `["123",456,"789"]`
	val := make([]int64, 0)
	err := jsoniter.ConfigConvertStringTo64.UnmarshalFromString(input, &val)
	should.Nil(err)
	should.Equal(3, len(val))
}

func Test_read_float_as_interface_with_quote(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, `12.3`)
	should.Equal(float64(12.3), iter.Read())
}

func Test_read_float64_cursor_with_quote(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigConvertStringTo64, "[1.23456789\n,2,3]")
	should.True(iter.ReadArray())
	should.Equal(1.23456789, iter.Read())
	should.True(iter.ReadArray())
	should.Equal(float64(2), iter.Read())
}

func Test_read_float64_array_with_quote(t *testing.T) {
	should := require.New(t)
	input := `["123.0",4560,"789.453"]`
	val := make([]float64, 0)
	err := jsoniter.ConfigConvertStringTo64.UnmarshalFromString(input, &val)
	should.Nil(err)
	should.Equal(3, len(val))
	should.Equal([]float64{123.0, 4560, 789.453}, val)
}
