package cmd

import (
	"testing"
)

func TestValidateBrawlNum(t *testing.T) {

	tests := []struct {
		got  error
		want error
	}{
		{
			got:  validateBrawlNum(1),
			want: nil,
		},
		{
			got:  validateBrawlNum(25),
			want: nil,
		},
	}

	for _, test := range tests {
		if test.got != nil {
			t.Errorf("Expected nil got: %v", test.got)
		}
	}
}

func TestValidateBrawlNum_invalid(t *testing.T) {

	tests := []struct {
		got error
	}{
		{
			got: validateBrawlNum(26),
		},
		{
			got: validateBrawlNum(0),
		},
		{
			got: validateBrawlNum(-1),
		},
		{
			got: validateBrawlNum(9223372036854775807),
		},
	}

	for _, test := range tests {
		if test.got == nil {
			t.Errorf("Expected error got: nil")
		}
	}
}

func TestValidateBrawlHPool(t *testing.T) {

	tests := []struct {
		got  error
		want error
	}{
		{
			got:  validateBrawlHPool("all"),
			want: nil,
		},
		{
			got:  validateBrawlHPool("brawl"),
			want: nil,
		},
	}

	for _, test := range tests {
		if test.got != nil {
			t.Errorf("Expected nil got: %v", test.got)
		}
	}
}

func TestValidateBrawlHPool_invalid(t *testing.T) {

	tests := []struct {
		got error
	}{
		{
			got: validateBrawlHPool("blahl"),
		},
		{
			got: validateBrawlHPool("poo"),
		},
	}

	for _, test := range tests {
		if test.got == nil {
			t.Errorf("Expected error got: nil")
		}
	}
}
