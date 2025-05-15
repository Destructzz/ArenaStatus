# PlayTracker

**PlayTracker** is a learning project built with Go that helps track how many times you've played with different people.

This app is created primarily for educational purposes — to explore Go language features and best practices — while also serving as a practical tool for gamers who want to log and analyze their co-players over time.

## How It Works

The app uses **OCR (Optical Character Recognition)** to automatically read player nicknames from match screenshots. This allows you to:

- Avoid manual input
- Quickly gather player data after each game
- Build statistics based on real screenshots

OCR is handled via [Tesseract OCR](https://github.com/tesseract-ocr/tesseract) integrated into the Go application.

## Features

- Detect player names from screenshots using OCR
- Track how many times you've played with each player
- Clean and minimalistic codebase written in Go
- Great for analyzing match history and spotting frequent teammates
- Helps build useful statistics for games with matchmaking or random players

## Purpose

- Learn Go through building a real-world application
- Practice working with OCR, image processing, file I/O, and data structures
- Create something genuinely useful along the way
