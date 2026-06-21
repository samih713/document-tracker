package db

const schemaSQL = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	role TEXT NOT NULL DEFAULT 'user',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS agents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	contact_no TEXT,
	email TEXT
);

CREATE TABLE IF NOT EXISTS documents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	vessel_name TEXT NOT NULL,
	blend TEXT,
	bl_date DATETIME,
	agent_id INTEGER NOT NULL,
	concession_holder TEXT,
	created_by_user_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	received_at DATETIME,
	is_closed INTEGER NOT NULL DEFAULT 0,

	FOREIGN KEY (agent_id) REFERENCES agents(id),
	FOREIGN KEY (created_by_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS document_events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	document_id INTEGER NOT NULL,
	event_type TEXT NOT NULL CHECK (
		event_type IN (
			'received',
			'revised',
			'signed_agent'
		)
	),
	note TEXT,
	next_action TEXT,
	updated_by_user_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY (document_id) REFERENCES documents(id),
	FOREIGN KEY (updated_by_user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_documents_agent
ON documents(agent_id);

CREATE INDEX IF NOT EXISTS idx_documents_blend
ON documents(blend);

CREATE INDEX IF NOT EXISTS idx_documents_events_document
ON document_events(document_id);

CREATE INDEX IF NOT EXISTS idx_document_events_type
ON document_events(event_type);
`
