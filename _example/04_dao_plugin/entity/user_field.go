// Code generated by eevee. DO NOT EDIT!

package entity

import "time"

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