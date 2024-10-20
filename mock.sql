INSERT INTO statuses (name) VALUES ('completed'), ('ongoing'), ('preview');

INSERT INTO types (name) VALUES ('Манга'), ('Манхва'), ('Комикс');

INSERT INTO genres (name) VALUES ('action'), ('drama'), ('dark'), ('fantasy');

INSERT INTO persons (name) VALUES ('Hajime Isayama');



INSERT INTO books (link, name, original_title, description, poster, genres, release_date, type_id, authors_ids, artists_ids, status_id)
VALUES ('shingeki-no-kyojin', 'Attack on Titan', 'Shingeki no Kyojin', 'Several hundred years ago, humans were nearly exterminated by Titans. Titans are typically several stories tall, seem to have no intelligence, devour human beings and, worst of all, seem to do it for pleasure rather than as a food source.', 'https://mangadex.org/covers/304ceac3-8cdb-4fe7-acf7-2b6ff7a60613/7b328381-b785-40b2-83b2-13e0f940a335.jpg.512.jpg',
'{1,2,3,4}', '11-25-2009', 1, '{1}', '{1}', 1);

INSERT INTO chapters (id, book_id, team_id, team_name, volume, chapter_number, name, is_draft)
VALUES ('3fe6be4c-c1cb-476c-a17c-c8bd15db3604', 1, 1, 'Dala Comics', 1, 1, 'Episode 1: To You, 2,000 Years From Now', false);

INSERT INTO chapters (id, book_id, team_id, team_name, volume, chapter_number, name, is_draft)
VALUES ('c3b3ee6f-8617-4873-b036-84021e18264d', 1, 1, 'Dala Comics', 1, 2, 'Episode 2: That Day', false);

INSERT INTO chapters (id, book_id, team_id, team_name, volume, chapter_number, name, is_draft)
VALUES ('4d4f2b48-5d4d-4403-b8c1-a415a1b2d1d7', 1, 1, 'Dala Comics', 1, 3, 'Episode 3: Night of the Disbanding Ceremony', false);

INSERT INTO chapters (id, book_id, team_id, team_name, volume, chapter_number, name, is_draft)
VALUES ('689d2957-13df-48a6-8bb8-36e8e20b6d36', 1, 1, 'Dala Comics', 1, 4, 'Episode 4: Their First Battle', false);