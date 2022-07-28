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