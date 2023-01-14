CREATE TABLE users(
    userid INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(150) UNIQUE NOT NULL
);

CREATE TABLE clip_stack (
    userid INT NOT NULL,
    clipid INT NOT NULL,
    message TEXT NOT NULL,
    secret BOOLEAN,
    PRIMARY KEY(userid, clipid),
    CONSTRAINT userid FOREIGN KEY(userid) REFERENCES users(userid)
);

CREATE TABLE passwords (
    userid INT NOT NULL,
    hash VARCHAR(150) NOT NULL,
    PRIMARY KEY(userid),
    CONSTRAINT userid FOREIGN KEY(userid) REFERENCES users(userid)
);