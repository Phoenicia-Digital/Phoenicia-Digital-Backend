CREATE TABLE contact_info (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(20) UNIQUE,
    CONSTRAINT email_or_phone_not_null CHECK (
        email IS NOT NULL OR phone_number IS NOT NULL
    )
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    message TEXT
);
