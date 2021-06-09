package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CompoundInterest(t *testing.T) {
	initialInvestment := EUR(1000)
	returnRate := Percent(5)
	years := 5

	subj := compoundInterest(initialInvestment, returnRate, years)

	assert.Equal(t, EUR(1276.28).String(), subj.String())
}

func Test_MonthlyCompoundInterestStarting(t *testing.T) {
	initialInvestment := EUR(1000)
	returnRate := Percent(3)

	cases := []struct {
		m   time.Month
		exp EUR
	}{
		{
			m:   time.January,
			exp: EUR(1030),
		},
		{
			m:   time.March,
			exp: EUR(1025),
		},
		{
			m:   time.June,
			exp: EUR(1017.5),
		},
		{
			m:   time.November,
			exp: EUR(1005),
		},
		{
			m:   time.December,
			exp: EUR(1002.5),
		},
	}

	for _, c := range cases {
		t.Run(c.m.String(), func(t *testing.T) {
			// when
			subj := monthlyCompoundInterestStarting(initialInvestment, returnRate, c.m)

			// then
			assert.Equal(t, c.exp.String(), subj.String())
		})
	}
}

func Test_MonthlyCompoundInterestUntil(t *testing.T) {
	initialInvestment := EUR(1000)
	returnRate := Percent(3)

	cases := []struct {
		m   time.Month
		exp EUR
	}{
		{
			m:   time.January,
			exp: EUR(1002.5),
		},
		{
			m:   time.March,
			exp: EUR(1007.5),
		},
		{
			m:   time.June,
			exp: EUR(1015),
		},
		{
			m:   time.November,
			exp: EUR(1027.5),
		},
		{
			m:   time.December,
			exp: EUR(1030),
		},
	}

	for _, c := range cases {
		t.Run(c.m.String(), func(t *testing.T) {
			// when
			subj := monthlyCompoundInterestUntil(initialInvestment, returnRate, c.m)

			// then
			assert.Equal(t, c.exp.String(), subj.String())
		})
	}
}

func Test_YearShare(t *testing.T) {
	t.Run("full year", func(t *testing.T) {
		assert.Equal(t, float64(1), yearShare(time.January))
	})

	t.Run("partial year near start", func(t *testing.T) {
		assert.Equal(t, 3.0/12.0, yearShare(time.March))
	})
	t.Run("partial year", func(t *testing.T) {
		assert.Equal(t, 0.5, yearShare(time.June))
	})

	t.Run("partial year near end", func(t *testing.T) {
		assert.Equal(t, 11.0/12.0, yearShare(time.November))
	})
}
