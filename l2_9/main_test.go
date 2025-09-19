package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	// Структура для описания одного тестового случая
	type testCase struct {
		name        string // Имя теста для понятного вывода
		input       string // Входные данные
		expected    string // Ожидаемый результат
		expectError bool   // Ожидается ли ошибка
	}

	// Набор тестовых случаев
	testCases := []testCase{
		// --- Основные случаи ---
		{
			name:     "Standard case",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			name:     "No numbers",
			input:    "abcd",
			expected: "abcd",
		},
		{
			name:        "Invalid: starts with digit",
			input:       "45",
			expectError: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character repetitions",
			input:    "a3b2c1",
			expected: "aaabbc",
		},
		{
			name:     "Repetition of 1",
			input:    "a1b1c1d1",
			expected: "abcd",
		},
		{
			name:     "Repetition of 0",
			input:    "a0b2c0d3",
			expected: "bbddd",
		},
		{
			name:     "Character at the end",
			input:    "a2b3c",
			expected: "aabbbc",
		},

		// --- Случаи с escape-последовательностями ---
		{
			name:     "Escaped digits",
			input:    "qwe\\4\\5",
			expected: "qwe45",
		},
		{
			name:     "Digit after escaped digit",
			input:    "qwe\\45",
			expected: "qwe44444",
		},
		{
			name:     "Escaped backslash",
			input:    "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
		},
		{
			name:     "Complex escapes",
			input:    "a2\\3b1\\45c0", // -> aa3b44444
			expected: "aa3b44444",
		},

		// --- Некорректные строки (ожидаем ошибку) ---
		{
			name:        "Starts with a digit",
			input:       "45",
			expectError: true,
		},
		{
			name:        "Only digits",
			input:       "123",
			expectError: true,
		},
		{
			name:        "Invalid sequence (digit after digit)",
			input:       "a23b",
			expectError: true,
		},
		{
			name:        "Ends with an escape character",
			input:       "abc\\",
			expectError: true,
		},
	}

	// Запуск тестов
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Выполняем тестируемую функцию
			result, err := unpack(tc.input)

			// Проверяем наличие ошибки
			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none")
				}
				// Если ошибка ожидалась и получена, тест пройден
				return
			}

			// Проверяем, что ошибки не было, когда её не ждали
			if err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
				return
			}

			// Сравниваем фактический результат с ожидаемым
			if result != tc.expected {
				t.Errorf("expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
