CREATE TABLE accounts
(
    account_id      INT     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    main_balance    INT     NOT NULL,
    reserve_balance INT     NOT NULL,
)

CREATE TABLE transactions
(
    transaction_id INT          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    account_id     INT          NOT NULL
	service_id     INT          NULL,
	order_id       INT          NULL,
	amount         INT          NOT NULL,
    trans_type     TEXT         NULL
    description    TEXT         NULL,
    FOREIGN KEY    (account_id)    REFERENCES accounts (account_id)
)