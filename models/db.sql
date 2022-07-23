-- DROP TABLE clip_stack;
-- DROP TABLE users;
-- DROP DATABASE cshare;
-- CREATE DATABASE cshare;
CREATE TABLE  users(
    user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(150) UNIQUE NOT NULL
);
CREATE TABLE  clip_stack
(
    user_id INT NOT NULL ,
    clip_id INT NOT NULL,
    message TEXT NOT NULL,
    secret BOOLEAN,
    PRIMARY KEY(user_id, clip_id),
    CONSTRAINT user_id
        FOREIGN KEY(user_id)
            REFERENCES users(user_id)
);
-- -- Test Data users
-- INSERT INTO users(username) VALUES ('koko');
-- INSERT INTO users(username) VALUES ('serpent');
--
-- INSERT INTO users(username) VALUES ('serpent'); -- expected error: duplicate
-- -- Test Data clip_stack
-- INSERT INTO clip_stack( user_id,clip_id, message, secret) VALUES (1, 1, 'Welcome to the board', FALSE);
-- INSERT INTO clip_stack( user_id,clip_id, message, secret) VALUES (1, 2, 'Welcome to the SQL', TRUE);
-- -- Check constrains data
-- INSERT INTO clip_stack( user_id,clip_id, message, secret) VALUES (1, 1, 'Welcome to the SQL', TRUE);
-- -- expected error due to presence of composite primary key
-- INSERT INTO clip_stack( user_id,clip_id, message, secret) VALUES (3, 1, 'Welcome to the SQL', TRUE);
-- -- expected error due to no data in foreign key
--
-- SELECT * FROM users;
-- SELECT * FROM clip_stack;
