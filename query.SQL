-- Enable the fuzzy matching extension
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Create the GIN index on the UUID cast to TEXT
-- This is what makes the "Top 50" search take milliseconds instead of seconds
CREATE INDEX CONCURRENTLY idx_users_qr_code_trgm 
ON users USING gin ((qr_code::text) gin_trgm_ops);