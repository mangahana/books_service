CREATE TABLE books (
  id             SERIAL PRIMARY KEY,
  link           VARCHAR(50) UNIQUE NOT NULL,
  name           VARCHAR(50) NOT NULL,
  original_title VARCHAR(50) NOT NULL,
  description    TEXT NOT NULL,
  poster         TEXT NOT NULL,
  genres         TEXT[] NOT NULL,
  release_date   DATE NOT NULL,
  type_id        TEXT NOT NULL,
  authors_ids    TEXT[] NOT NULL,
  artists_ids    TEXT[] NOT NULL,
  formats_ids    TEXT[] NOT NULL,
  status_id      TEXT NOT NULL,
  owner_team_id  INTEGER NOT NULL DEFAULT 0,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE genres (
  id   TEXT UNIQUE NOT NULL,
  name VARCHAR(50) NOT NULL
);
INSERT INTO genres (id, name)
VALUES 
('aciton', 'Экшн');

CREATE TABLE types (
  id   TEXT UNIQUE NOT NULL,
  name VARCHAR(50) NOT NULL
);
INSERT INTO types (id, name)
VALUES 
('manga', 'Маңга'),
('manhwa', 'Манхуа'),
('manhua', 'Мәңхуә'),
('comics', 'Комикс');


CREATE TABLE formats (
  id   TEXT UNIQUE NOT NULL,
  NAME VARCHAR(50) NOT NULL
);
INSERT INTO formats (id, name)
VALUES 
('oneshot', 'Уаншот'),
('web-comic', 'Уеб комикс');


CREATE TABLE statuses (
  id   TEXT UNIQUE NOT NULL,
  name VARCHAR(50) NOT NULL
);
INSERT INTO statuses (id, name)
VALUES
('ongoing', 'Онгоиң'),
('completed', 'Аяқталған'),
('hiatus', 'Хиатус'),
('canceled', 'Доғарылған');


CREATE TABLE persons (
  id    SERIAL PRIMARY KEY,
  name  VARCHAR(50) NOT NULL
);


CREATE TABLE chapters (
  id             TEXT UNIQUE NOT NULL,
  book_id        INTEGER NOT NULL,
  team_id        INTEGER NOT NULL,
  team_name      TEXT NOT NULL,
  member_id      INTEGER NOT NULL ,
  volume         INTEGER NOT NULL DEFAULT 0,
  chapter_number VARCHAR(10) NOT NULL DEFAULT 0,
  name           TEXT NOT NULL DEFAULT '',
  pages           TEXT[] NOT NULL DEFAULT '{}',
  is_draft       BOOLEAN NOT NULL DEFAULT TRUE,
  created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_book FOREIGN KEY(book_id) REFERENCES books(id)
);