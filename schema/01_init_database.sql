CREATE TYPE txType AS ENUM ('deposit', 'withdrawal')

CREATE TABLE transactions (
 id uint PRIMARY KEY
 tx_Type txType
 amount uint
 created_at timestamp without time zone
 updated_at timestamp without time zone
);