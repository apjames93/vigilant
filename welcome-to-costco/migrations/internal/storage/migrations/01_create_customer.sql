-- Create table for storing users --
DROP TABLE IF EXISTS customers;

CREATE TABLE IF NOT EXISTS customers (
  id SERIAL NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
);

-- -- Create table for storing users --
-- DROP TABLE IF EXISTS users;

-- CREATE TABLE IF NOT EXISTS users (
--   id SERIAL NOT NULL PRIMARY KEY,
--   email TEXT NOT NULL UNIQUE,
--   encrypted_password TEXT NOT NULL,
--   verificationtoken TEXT,
--   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
-- );




-- cusotmer 
-- | id | name |

-- Projects 
-- | id | name | 

-- schedule 
-- | start_date | end_date | project_id

-- project_schedule_user_priority
-- | user_id | sechedule_id | priorty

-- User
-- | email | encrypted_password | verificationtoken

-- cusotmer_project_users
-- | user_id | project_id | customer_id |


