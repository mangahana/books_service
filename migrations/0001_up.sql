CREATE TABLE books (
  id             SERIAL PRIMARY KEY,
  link           VARCHAR(50) UNIQUE NOT NULL,
  name           VARCHAR(50) NOT NULL,
  original_title VARCHAR(50) NOT NULL,
  description    TEXT NOT NULL,
  poster         TEXT NOT NULL,
  genres         INTEGER[] NOT NULL,
  release_date   DATE NOT NULL,
  type_id        INTEGER NOT NULL,
  authors_ids    INTEGER[] NOT NULL,
  artists_ids    INTEGER[] NOT NULL,
  status_id      INTEGER NOT NULL,
  owner_team_id  INTEGER NOT NULL DEFAULT 0,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE genres (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE types (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE statuses (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE persons (
  id    SERIAL PRIMARY KEY,
  name  VARCHAR(50) NOT NULL
);

CREATE TABLE chapters (
  id             TEXT UNIQUE NOT NULL,
  book_id        INTEGER NOT NULL,
  team_id        INTEGER NOT NULL,
  volume         INTEGER,
  chapter_number VARCHAR(10),
  is_draft       BOOLEAN NOT NULL DEFAULT FALSE,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pages (
  chapter_id  TEXT NOT NULL,
  page_number INTEGER NOT NULL,
  image       TEXT NOT NULL
);