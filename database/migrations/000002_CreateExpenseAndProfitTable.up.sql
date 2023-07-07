CREATE TABLE IF NOT EXISTS finance.expense_or_profit (
    id                  SERIAL PRIMARY KEY,
    user_id             INT NOT NULL,
    expense_or_profit   SMALLINT NOT NULL,
    value_to_register   BIGINT NOT NULL,
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP
);

ALTER TABLE finance.expense_or_profit ADD CONSTRAINT fk_expenseprofit_userid FOREIGN KEY (user_id) REFERENCES finance.user (id);