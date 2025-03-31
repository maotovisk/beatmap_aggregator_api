# Beatmap Aggregator API – Golang

This is a lightweight API I built to save osu! beatmaps into a SQLite database. It's a boilerplate setup I plan to use for future Go projects — and maybe build a frontend on top of it later.

## Tech Stack

- **Go** 1.21
- **GORM** – ORM for Go
- **SQLite3** – Embedded database
- **Chi** – Lightweight HTTP router

## Features & Roadmap

- [x] **Add beatmaps** – `POST /beatmap` with a beatmap URL
- [x] **List beatmapsets** – `GET /beatmap`
- [ ] **JWT Authentication** – Basic token-based auth (coming soon)
