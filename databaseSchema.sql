CREATE TABLE file_tags (
    tag VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    PRIMARY KEY (tag, file_path)
);
/