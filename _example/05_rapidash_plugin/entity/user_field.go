// Code generated by eevee. DO NOT EDIT!

package entity

import (
	"go.knocknote.io/rapidash"
	"golang.org/x/xerrors"
	"time"
)

type UserField struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"userID"`
	FieldID   uint64    `json:"fieldID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserFields []*UserField

func (e UserFields) IDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.ID)
	}
	return values
}

func (e UserFields) UserIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.UserID)
	}
	return values
}

func (e UserFields) FieldIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.FieldID)
	}
	return values
}

func (e UserFields) CreatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.CreatedAt)
	}
	return values
}

func (e UserFields) UpdatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.UpdatedAt)
	}
	return values
}

func (e *UserField) Struct() *rapidash.Struct {
	return rapidash.NewStruct("user_fields").
		FieldUint64("id").
		FieldUint64("user_id").
		FieldUint64("field_id").
		FieldTime("created_at").
		FieldTime("updated_at")
}

func (e *UserField) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.Uint64("user_id", e.UserID)
	enc.Uint64("field_id", e.FieldID)
	enc.Time("created_at", e.CreatedAt)
	enc.Time("updated_at", e.UpdatedAt)
	if err := enc.Error(); err != nil {
		return xerrors.Errorf("failed to encode: %w", err)
	}
	return nil
}

func (e *UserFields) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return xerrors.Errorf("failed to encode: %w", err)
		}
	}
	return nil
}

func (e *UserField) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.UserID = dec.Uint64("user_id")
	e.FieldID = dec.Uint64("field_id")
	e.CreatedAt = dec.Time("created_at")
	e.UpdatedAt = dec.Time("updated_at")
	if err := dec.Error(); err != nil {
		return xerrors.Errorf("failed to decode: %w", err)
	}
	return nil
}

func (e *UserFields) DecodeRapidash(dec rapidash.Decoder) error {
	decLen := dec.Len()
	values := make(UserFields, decLen)
	for i := 0; i < decLen; i++ {
		var v UserField
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return xerrors.Errorf("failed to decode: %w", err)
		}
		values[i] = &v
	}
	*e = values
	return nil
}