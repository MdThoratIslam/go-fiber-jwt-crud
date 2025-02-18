package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

// ✅ লগ ফাইল এবং কনসোলে লগিং সেটআপ
func init() {
	// লগ ফাইল তৈরি / ওপেন করা (Append Mode)
	// %s দিয়ে তারিখ ফরম্যাট করা হয়েছে (2006-01-02) এবং ফাইলের নাম তৈরি করা হয়েছে
	fileName := fmt.Sprintf("log/app/app_%s.log", time.Now().Format("02-01-2006"))
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("❌ Could not open log file:", err)
	}

	// ✅ লগ ফাইল এবং কনসোলে একসাথে লগ প্রিন্ট করা
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// ✅ লগ লেভেল সেট করা (INFO এর নিচের লগ গুলো প্রিন্ট হবে না)
	log.SetLevel(log.LevelInfo)

	// ✅ লগিং শুরু
	log.Info("🚀 Logger initialized successfully!")
}

// ✅ Custom Error log function
func Error(msg string, err error) {
	if err != nil {
		log.Errorf("%s: %s", msg, err.Error())
	} else {
		log.Error(msg)
	}
}

// ✅ Custom Success log function
func Success(msg string) {
	log.Info(msg)
}

// ✅ Custom Warning log function
func Warn(msg string) {
	log.Warn(msg)
}

// ✅ Custom Debug log function
func Debug(msg string) {
	log.Debug(msg)
}
