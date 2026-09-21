# FAQ Chatbot

A simple, typo-tolerant Q&A chatbot using Levenshtein distance for FAQ matching.

## [English]

This chatbot matches user questions against a predefined set of FAQ entries based on keyword similarity using the Levenshtein distance algorithm.

### Features
- **Levenshtein Distance:** Implements a memory-efficient 1D slice approach for typo-tolerant string matching.
- **FAQ Matching:** Matches user input against `faq.txt` based on keyword similarity.

### Getting Started
1. **Build:** `go build -o chatbot main.go levenshtein.go`
2. **Run:** `./chatbot`
3. **Usage:** Enter your question when prompted. Type `exit` to quit.

### FAQ Data Format (`faq.txt`)
Each line should be `Question|Answer`. Example:
`Когда будет репитиция?|Репетиция идет прямо сейчас.`

### Tests
Run `go test .` to verify the functionality.

---

## [Русский]

Это простой чат-бот для вопросов и ответов, который использует расстояние Левенштейна для сопоставления вопросов пользователя с предопределенным списком FAQ.

### Функции
- **Расстояние Левенштейна:** Реализован эффективный алгоритм с использованием 1D-массива для исправления опечаток.
- **Сопоставление FAQ:** Сравнивает ввод пользователя с вопросами в файле `faq.txt` на основе схожести ключевых слов.

### Начало работы
1. **Сборка:** `go build -o chatbot main.go levenshtein.go`
2. **Запуск:** `./chatbot`
3. **Использование:** Введите вопрос после приглашения. Введите `exit`, чтобы выйти.

### Формат данных FAQ (`faq.txt`)
Каждая строка должна быть в формате `Вопрос|Ответ`. Пример:
`Когда будет репитиция?|Репетиция идет прямо сейчас.`

### Тестирование
Запустите `go test .` для проверки работоспособности.
