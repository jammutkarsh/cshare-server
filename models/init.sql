CREATE TABLE  users(
    userID INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(150) UNIQUE NOT NULL
);
CREATE TABLE  clip_stack
(
    userID INT NOT NULL ,
    clipID INT NOT NULL,
    message TEXT NOT NULL,
    secret BOOLEAN,
    PRIMARY KEY(userID, clipID),
    CONSTRAINT userID
        FOREIGN KEY(userID)
            REFERENCES users(userID)
);

CREATE TABLE  passwords
(
    userID INT NOT NULL,
    hash VARCHAR(150) NOT NULL,
    PRIMARY KEY(userID),
    CONSTRAINT userID
        FOREIGN KEY(userID)
            REFERENCES users(userID)
);