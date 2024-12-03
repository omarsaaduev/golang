package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

// SetupLogger создает и настраивает логгер.
func SetupLogger() *slog.Logger {
	// Получаем путь к корню проекта
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(currentFile), "../..") // Путь к корню проекта

	// Указываем путь к директории и файлу логов
	logDir := filepath.Join(projectRoot, "logs") // Папка logs в корне проекта
	logFile := filepath.Join(logDir, "app.log")  // Файл logs/app.log

	// Создаем директорию logs, если её нет
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return nil
	}

	// Открываем файл для записи логов
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	// Создаем хендлер для записи в файл
	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Создаем хендлер для вывода в stdout
	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Создаем логгер с хендлером для записи в файл
	logger := slog.New(fileHandler)

	// Создаем новый логгер с дополнительным хендлером для stdout
	// Мы можем использовать новый логгер с двух хендлеров
	multiLogger := slog.New(stdoutHandler)

	// Комбинируем два логгера в один
	logger = logger.With(multiLogger)

	return logger
}
