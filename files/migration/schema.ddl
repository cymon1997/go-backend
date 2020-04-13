
-- Generate UUID = md5(random()::text || clock_timestamp()::text)::uuid

CREATE TABLE article (
    id UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
    title VARCHAR(20) NOT NULL,
    description VARCHAR(40),
    content TEXT,
    create_time TIMESTAMP DEFAULT now(),
    create_by varchar(20),
    update_time TIMESTAMP DEFAULT now(),
    update_by varchar(20)
);