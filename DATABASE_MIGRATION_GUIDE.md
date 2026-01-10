# Database Migration Guide

## Overview

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database schema
versioning and migrations. All migrations are automatically applied when the application starts up.

## Migration System Architecture

- **Database**: PostgreSQL
- **Migration Library**: `golang-migrate/migrate/v4`
- **ORM**: GORM
- **Migration Path**: `pkg/database/migrations/`
- **Auto-execution**: Migrations run automatically on application startup

## Migration File Structure

Migration files are stored in `pkg/database/migrations/` and follow this naming convention:

```
{version}_{name}.up.sql    # Forward migration
{version}_{name}.down.sql  # Rollback migration
```

### Example:

```
0012_new_model.up.sql      # Creates the new model
0012_new_model.down.sql    # Drops the new model
```

## How to Create a New Model Migration

### Step 1: Determine the Next Version Number

Look at the existing migrations in `pkg/database/migrations/` to find the highest version number:

```bash
ls pkg/database/migrations/ | grep -E '^[0-9]+_' | sort -V | tail -1
```

For example, if the latest migration is `0011_mock_data.up.sql`, your new migration should be
`0012_`.

### Step 2: Create Migration Files

Create both UP and DOWN migration files:

```bash
# Navigate to the migrations directory
cd pkg/database/migrations/

# Create the migration files (replace 0012 with your version number)
touch 0012_new_model.up.sql
touch 0012_new_model.down.sql
```

### Step 3: Write the UP Migration (`0012_new_model.up.sql`)

The UP migration should create your new model/table with all necessary constraints, indexes, and
triggers.

**Example UP migration:**

```sql
-- Create new_model table
CREATE TABLE IF NOT EXISTS new_models (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    -- Foreign key constraints
    CONSTRAINT fk_new_models_user_id 
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_new_models_user_id ON new_models(user_id);
CREATE INDEX IF NOT EXISTS idx_new_models_status ON new_models(status);
CREATE INDEX IF NOT EXISTS idx_new_models_created_at ON new_models(created_at);

-- Create trigger for auto-updating updated_at timestamp
DROP TRIGGER IF EXISTS new_models_set_updated_at ON new_models;
CREATE TRIGGER new_models_set_updated_at
    BEFORE UPDATE ON new_models
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- Add any initial data if needed
INSERT INTO new_models (name, description, user_id) VALUES 
    ('Default Model', 'Default model description', 1)
ON CONFLICT DO NOTHING;
```

### Step 4: Write the DOWN Migration (`0012_new_model.down.sql`)

The DOWN migration should completely reverse the UP migration:

```sql
-- Drop trigger
DROP TRIGGER IF EXISTS new_models_set_updated_at ON new_models;

-- Drop indexes
DROP INDEX IF EXISTS idx_new_models_created_at;
DROP INDEX IF EXISTS idx_new_models_status;
DROP INDEX IF EXISTS idx_new_models_user_id;

-- Drop table
DROP TABLE IF EXISTS new_models;
```

## Migration Best Practices

### 1. **Always Create Both UP and DOWN Files**

- UP file: Contains the forward migration
- DOWN file: Contains the rollback logic
- Both files must be created even if DOWN is just dropping tables

### 2. **Use IF EXISTS/IF NOT EXISTS**

```sql
-- Good
CREATE TABLE IF NOT EXISTS my_table (...);
DROP TABLE IF EXISTS my_table;

-- Avoid
CREATE TABLE my_table (...);  -- Will fail if table exists
```

### 3. **Handle Data Carefully**

- When dropping columns, consider data loss
- When adding NOT NULL columns, provide DEFAULT values
- Use `ON CONFLICT DO NOTHING` for initial data inserts

### 4. **Include Proper Constraints**

```sql
-- Foreign keys
CONSTRAINT fk_table_column FOREIGN KEY (column) REFERENCES other_table(id)

-- Check constraints
CONSTRAINT chk_status CHECK (status IN ('active', 'inactive', 'pending'))

-- Unique constraints
CONSTRAINT uk_table_column UNIQUE (column)
```

### 5. **Add Performance Indexes**

```sql
-- Single column index
CREATE INDEX idx_table_column ON table_name(column);

-- Composite index
CREATE INDEX idx_table_col1_col2 ON table_name(col1, col2);

-- Partial index
CREATE INDEX idx_table_status_active ON table_name(status) WHERE status = 'active';
```

### 6. **Use the Updated_at Trigger Pattern**

Most tables should include the automatic `updated_at` trigger:

```sql
-- In your table creation
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

-- After table creation
DROP TRIGGER IF EXISTS table_set_updated_at ON table_name;
CREATE TRIGGER table_set_updated_at
    BEFORE UPDATE ON table_name
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
```

## Testing Your Migration

### 1. **Test the Migration Locally**

Before committing, test your migration:

```bash
# Run the application (migrations run automatically)
make run-local

# Check logs to ensure migration succeeded
# Look for messages like "migrate up: no change" or successful migration logs
```

### 2. **Test Rollback (Optional)**

If you need to test rollback capabilities:

```sql
-- Connect to your database and manually test
-- This is advanced usage and not typically needed for normal development
```

### 3. **Verify Database State**

Connect to your PostgreSQL database and verify:

```sql
-- Check if table was created
\dt new_models

-- Check table structure
\d new_models

-- Check indexes
\di new_models*

-- Check triggers
\dt+ new_models
```

## Migration Execution Flow

1. **Automatic Execution**: When you run `make run-local`, migrations execute automatically
2. **Version Tracking**: golang-migrate tracks which migrations have been applied
3. **Sequential Execution**: Migrations run in numerical order (0001, 0002, 0003, etc.)
4. **Idempotent**: Re-running the app won't re-execute already applied migrations

## Common Migration Scenarios

### Adding a New Table

```sql
-- UP
CREATE TABLE IF NOT EXISTS new_table (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- DOWN
DROP TABLE IF EXISTS new_table;
```

### Adding a Column to Existing Table

```sql
-- UP
ALTER TABLE existing_table 
ADD COLUMN IF NOT EXISTS new_column VARCHAR(255) DEFAULT '';

-- DOWN
ALTER TABLE existing_table 
DROP COLUMN IF EXISTS new_column;
```

### Modifying a Column

```sql
-- UP
ALTER TABLE existing_table 
ALTER COLUMN existing_column TYPE TEXT;

-- DOWN
ALTER TABLE existing_table 
ALTER COLUMN existing_column TYPE VARCHAR(255);
```

### Adding Foreign Key Relationship

```sql
-- UP
ALTER TABLE child_table 
ADD CONSTRAINT fk_child_parent 
FOREIGN KEY (parent_id) REFERENCES parent_table(id);

-- DOWN
ALTER TABLE child_table 
DROP CONSTRAINT IF EXISTS fk_child_parent;
```

## Troubleshooting

### Migration Failed

- Check the error logs when running the application
- Verify SQL syntax in your migration files
- Ensure proper foreign key references exist
- Check for data conflicts (e.g., NOT NULL constraints on existing data)

### Rolling Back

- Migrations automatically rollback on failure
- Manual rollback is advanced and should be avoided in production

### Schema Conflicts

- Always pull latest code before creating new migrations
- If conflicts occur, create a new migration rather than modifying existing ones

## File Naming Examples

```
0012_users_add_email_verification.up.sql
0012_users_add_email_verification.down.sql

0013_products_create_table.up.sql
0013_products_create_table.down.sql

0014_orders_add_status_index.up.sql
0014_orders_add_status_index.down.sql
```

## Integration with GORM Models

After creating your database migration, create the corresponding GORM model:

```go
// internal/domain/newmodel/models.go
type NewModel struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `gorm:"not null" json:"name"`
    Description string    `json:"description"`
    Status      string    `gorm:"default:'active'" json:"status"`
    UserID      uint      `gorm:"not null" json:"user_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    
    // Relationships
    User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
```

## Summary

1. **Find the next version number** by checking existing migrations
2. **Create both UP and DOWN migration files** with proper naming
3. **Write SQL migrations** following best practices
4. **Test locally** by running the application
5. **Commit and deploy** - migrations run automatically

Remember: Never modify existing migration files once they're in production. Always create new
migrations for schema changes.