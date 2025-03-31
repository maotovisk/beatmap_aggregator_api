# Beatmap Aggregator API – Golang

A lightweight API built to store osu! beatmaps in a SQLite database.  
This serves as a boilerplate for future Go projects — and possibly a base for a frontend later on.  
It's also my personal playground while learning and experimenting with Golang.

## Tech Stack

- **Go** 1.21
- **GORM** – ORM for Go
- **SQLite3** – Embedded database
- **Chi** – Lightweight HTTP router

## Features & Roadmap

- [x] **Add beatmaps** – `POST /beatmap` with a beatmap URL
- [x] **List beatmapsets** – `GET /beatmap`
- [ ] **JWT Authentication** – Basic token-based auth (coming soon)
