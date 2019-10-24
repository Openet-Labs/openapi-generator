/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// TypeHolderExample struct for TypeHolderExample
type TypeHolderExample struct {
	StringItem string `json:"string_item"`
	NumberItem float32 `json:"number_item"`
	FloatItem float32 `json:"float_item"`
	IntegerItem int32 `json:"integer_item"`
	BoolItem bool `json:"bool_item"`
	ArrayItem []int32 `json:"array_item"`
}

// GetStringItem returns the StringItem field value
func (o *TypeHolderExample) GetStringItem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StringItem
}

// SetStringItem sets field value
func (o *TypeHolderExample) SetStringItem(v string) {
	o.StringItem = v
}

// GetNumberItem returns the NumberItem field value
func (o *TypeHolderExample) GetNumberItem() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumberItem
}

// SetNumberItem sets field value
func (o *TypeHolderExample) SetNumberItem(v float32) {
	o.NumberItem = v
}

// GetFloatItem returns the FloatItem field value
func (o *TypeHolderExample) GetFloatItem() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FloatItem
}

// SetFloatItem sets field value
func (o *TypeHolderExample) SetFloatItem(v float32) {
	o.FloatItem = v
}

// GetIntegerItem returns the IntegerItem field value
func (o *TypeHolderExample) GetIntegerItem() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntegerItem
}

// SetIntegerItem sets field value
func (o *TypeHolderExample) SetIntegerItem(v int32) {
	o.IntegerItem = v
}

// GetBoolItem returns the BoolItem field value
func (o *TypeHolderExample) GetBoolItem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BoolItem
}

// SetBoolItem sets field value
func (o *TypeHolderExample) SetBoolItem(v bool) {
	o.BoolItem = v
}

// GetArrayItem returns the ArrayItem field value
func (o *TypeHolderExample) GetArrayItem() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ArrayItem
}

// SetArrayItem sets field value
func (o *TypeHolderExample) SetArrayItem(v []int32) {
	o.ArrayItem = v
}

type NullableTypeHolderExample struct {
	Value TypeHolderExample
	ExplicitNull bool
}

func (v NullableTypeHolderExample) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableTypeHolderExample) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

