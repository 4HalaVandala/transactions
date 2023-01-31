
CREATE TABLE transactions (
 id serial PRIMARY KEY NOT NULL,
 username varchar(100),
 amount int NOT NULL,
 TxType_id int NOT NULL,
 created_at timestamp without time zone NOT NULL,
 updated_at timestamp without time zone NOT NULL
);

CREATE TABLE transaction_type (
    id serial PRIMARY KEY NOT NULL,
    type varchar(255) NOT NULL
);

INSERT INTO transaction_type(id,type) VALUES ('1', 'deposit');
INSERT INTO transaction_type(id,type) VALUES ('2', 'withdrawal');

