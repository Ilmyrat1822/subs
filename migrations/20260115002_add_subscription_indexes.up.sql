-- Index for filtering by user
CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id
ON subscriptions (user_id);

-- Composite index for user + service
CREATE INDEX IF NOT EXISTS idx_subscriptions_user_service
ON subscriptions (user_id, service_name);

-- Optional: index for date filtering
CREATE INDEX IF NOT EXISTS idx_subscriptions_start_date
ON subscriptions (start_date);
