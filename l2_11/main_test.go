package main

import (
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Основной случай из ТЗ",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "Пустой срез",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Нет анаграмм",
			input:    []string{"один", "два", "три", "четыре"},
			expected: map[string][]string{},
		},
		{
			name:  "Слова в разном регистре",
			input: []string{"Пятак", "пятка", "Тяпка"},
			expected: map[string][]string{
				"Пятак": {"Пятак", "Тяпка", "пятка"},
			},
		},
		{
			name:  "Несколько групп анаграмм",
			input: []string{"кот", "ток", "кто", "рост", "сорт", "торс"},
			expected: map[string][]string{
				"кот":  {"кот", "кто", "ток"},
				"рост": {"рост", "сорт", "торс"},
			},
		},
		{
			name:     "Только одиночные слова",
			input:    []string{"слово"},
			expected: map[string][]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findAnagram(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("findAnagram(%v) = %v, ожидалось %v", tc.input, result, tc.expected)
			}
		})
	}
}
