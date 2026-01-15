CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    user_id UUID NOT NULL,
    start_date VARCHAR(7) NOT NULL,
    end_date VARCHAR(7),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_service_name ON subscriptions(service_name);
CREATE INDEX idx_subscriptions_start_end ON subscriptions(start_date, end_date);
