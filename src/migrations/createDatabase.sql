CREATE TABLE if not exists authors (
    id SERIAL PRIMARY KEY,
    FIO CHAR NOT NULL NOT NULL,
    pseudonym CHAR,
    specialization VARCHAR NOT NULL
);
CREATE TABLE books(
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    ISBN CHAR NOT NULL UNIQUE,
    genre char,
    author_id INT,
    member_id INT,
    FOREIGN KEY (author_id) REFERENCES authors(id),
    FOREIGN KEY (member_id) REFERENCES members(ID)
);

CREATE TABLE IF NOT EXISTS members(
    id SERIAL PRIMARY KEY,
    FIO VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS subscriptions(
    id SERIAL PRIMARY KEY,
    book_id int,
    member_id int,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);
