CREATE TABLE IF NOT EXISTS students (
    nuid integer PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    book_id serial PRIMARY KEY,
    title varchar NOT NULL,
    author varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS checked_out_books (
    checkout_id serial PRIMARY KEY,
    book_id integer NOT NULL references books(book_id),
    nuid integer NOT NULL references students(nuid),
    expected_return_date timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS holds (
    hold_id serial PRIMARY KEY,
    book_id integer NOT NULL references books(book_id),
    nuid integer NOT NULL references students(nuid),
    hold_creation_date timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS liked_books (
    like_id serial PRIMARY KEY,
    book_id integer NOT NULL references books(book_id),
    nuid integer NOT NULL references students(nuid)
);
