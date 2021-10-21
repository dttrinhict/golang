DROP TYPE IF EXISTS golang.rolenum;

CREATE TYPE golang.rolenum AS ENUM (
'ADMIN',
'STUDENT',
'TRAINER',
'SALE',
'EMPLOYER',
'AUTHOR',
'EDITOR',
'MAINTAINER');


DROP TABLE IF EXISTS golang.users;
CREATE TABLE golang.users (
                            id text PRIMARY KEY,
                            name text NOT NULL,
                            email text,
                            mobile text,
                            int_roles int[],
                            enum_roles test.rolenum[]
);

DROP TABLE IF EXISTS golang.roles;
CREATE TABLE golang.roles (
                            id int PRIMARY KEY,
                            name text NOT NULL UNIQUE
);
INSERT INTO golang.roles (id, name) VALUES (1, 'ADMIN');
INSERT INTO golang.roles (id, name) VALUES (2, 'STUDENT');
INSERT INTO golang.roles (id, name) VALUES (3, 'TRAINER');
INSERT INTO golang.roles (id, name) VALUES (4, 'SALE');
INSERT INTO golang.roles (id, name) VALUES (5, 'EMPLOYER');
INSERT INTO golang.roles (id, name) VALUES (6, 'AUTHOR');
INSERT INTO golang.roles (id, name) VALUES (7, 'EDITOR');
INSERT INTO golang.roles (id, name) VALUES (8, 'MAINTAINER');

DROP TABLE IF EXISTS test.user_role;
CREATE TABLE test.user_role (
                                user_id text REFERENCES golang.users(id),
                                role_id int REFERENCES golang.roles(id)
);

CREATE UNIQUE INDEX user_idx ON golang.user_role (user_id, role_id);