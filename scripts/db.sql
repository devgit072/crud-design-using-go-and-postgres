CREATE TABLE employee (
  id SERIAL PRIMARY KEY,
  name TEXT,
  age INT,
  department TEXT,
  annual_salary INT,
  social_security_number TEXT,
  address TEXT,
  created_at timestamp,
  last_updated_at timestamp
);

