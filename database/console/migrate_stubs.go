package console

type MysqlStubs struct{}

// CreateUp Create up migration content.
func (receiver MysqlStubs) CreateUp() string {
	return `CREATE TABLE DummyTable (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_DummyTable_created_at (created_at),
  KEY idx_DummyTable_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = DummyDatabaseCharset;
`
}

// CreateDown Create down migration content.
func (receiver MysqlStubs) CreateDown() string {
	return `DROP TABLE IF EXISTS DummyTable;
`
}

// UpdateUp Update up migration content.
func (receiver MysqlStubs) UpdateUp() string {
	return `ALTER TABLE DummyTable ADD column VARCHAR(255);
`
}

// UpdateDown Update down migration content.
func (receiver MysqlStubs) UpdateDown() string {
	return `ALTER TABLE DummyTable DROP COLUMN column;
`
}

type PostgresqlStubs struct{}

// CreateUp Create up migration content.
func (receiver PostgresqlStubs) CreateUp() string {
	return `CREATE TABLE DummyTable (
  id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_DummyTable_created_at ON DummyTable (created_at);
CREATE INDEX idx_DummyTable_updated_at ON DummyTable (updated_at);
`
}

// CreateDown Create down migration content.
func (receiver PostgresqlStubs) CreateDown() string {
	return `DROP TABLE IF EXISTS DummyTable;
`
}

// UpdateUp Update up migration content.
func (receiver PostgresqlStubs) UpdateUp() string {
	return `ALTER TABLE DummyTable ADD column TEXT NOT NULL;
`
}

// UpdateDown Update down migration content.
func (receiver PostgresqlStubs) UpdateDown() string {
	return `ALTER TABLE DummyTable DROP COLUMN column;
`
}

type SqliteStubs struct{}

// CreateUp Create up migration content.
func (receiver SqliteStubs) CreateUp() string {
	return `CREATE TABLE DummyTable (
  id BIGINT PRIMARY KEY AUTOINCREMENT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);

CREATE INDEX idx_DummyTable_created_at ON DummyTable (created_at);
CREATE INDEX idx_DummyTable_updated_at ON DummyTable (updated_at);
`
}

// CreateDown Create down migration content.
func (receiver SqliteStubs) CreateDown() string {
	return `DROP TABLE IF EXISTS DummyTable;
`
}

// UpdateUp Update up migration content.
func (receiver SqliteStubs) UpdateUp() string {
	return `ALTER TABLE DummyTable ADD column TEXT;
`
}

// UpdateDown Update down migration content.
func (receiver SqliteStubs) UpdateDown() string {
	return `ALTER TABLE DummyTable DROP COLUMN column;
`
}

type SqlserverStubs struct{}

// CreateUp Create up migration content.
func (receiver SqlserverStubs) CreateUp() string {
	return `CREATE TABLE DummyTable (
  id BIGINT NOT NULL IDENTITY(1,1),
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX idx_DummyTable_created_at ON DummyTable (created_at);
CREATE INDEX idx_DummyTable_updated_at ON DummyTable (updated_at);
`
}

// CreateDown Create down migration content.
func (receiver SqlserverStubs) CreateDown() string {
	return `DROP TABLE IF EXISTS DummyTable;
`
}

// UpdateUp Update up migration content.
func (receiver SqlserverStubs) UpdateUp() string {
	return `ALTER TABLE DummyTable ADD column VARCHAR(255);
`
}

// UpdateDown Update down migration content.
func (receiver SqlserverStubs) UpdateDown() string {
	return `ALTER TABLE DummyTable DROP COLUMN column;
`
}
