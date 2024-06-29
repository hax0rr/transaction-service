CREATE TABLE accounts
(
    id                  text        PRIMARY KEY,
    document_number     text        NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE transactions
(
    id                  text        PRIMARY KEY,
    account_id          text        NOT NULL,
    operation_type_id   SMALLINT    NOT NULL,
    amount              NUMERIC     NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_account_id FOREIGN KEY(account_id) references accounts(id)
);
