// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreSQLCreateTablesSecret - создаёт таблицы для схемы secret (для админки и авторизации)
func PostgreSQLCreateTablesSecret(dbc *pgxpool.Pool) {

	// Админка

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "users",
			CreateStatement: `CREATE TABLE secret.users
			(
				id uuid NOT NULL,
				role character varying(50) NOT NULL,
				email character varying(50) NOT NULL,
				phone character varying(50),
				name character varying(150),
				isadmin boolean,
				confirmed boolean,
				disabled boolean,
				totp_active boolean,
				PRIMARY KEY (id)
			);
			
			ALTER TABLE secret.users
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "hashes",
			CreateStatement: `CREATE TABLE secret.hashes
			(
				id bigserial NOT NULL,
				user_id uuid NOT NULL,
				value bytea NOT NULL,
				PRIMARY KEY (id)
			);
	
			ALTER TABLE secret.hashes
				OWNER to postgres;
				
			ALTER TABLE secret.hashes
				ADD CONSTRAINT hashes_user_id_fkey FOREIGN KEY (user_id)
				REFERENCES secret.users (id) MATCH FULL
				ON UPDATE RESTRICT
				ON DELETE CASCADE;
			CREATE INDEX fki_hashes_user_id_fkey
				ON secret.hashes(user_id);`,
		},
		NamedCreateStatement{
			TableName: "confirmations",
			CreateStatement: `CREATE TABLE secret.confirmations
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				created timestamp with time zone NOT NULL,
				expires timestamp with time zone NOT NULL,
				CONSTRAINT confirmations_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			);
			ALTER TABLE secret.confirmations
				OWNER to postgres;
			CREATE INDEX fki_confirmations_user_id_fkey
				ON secret.confirmations USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "password_resets",
			CreateStatement: `CREATE TABLE secret.password_resets
			(
				user_id uuid,
				token character varying(200) COLLATE pg_catalog."default" NOT NULL,
				created timestamp with time zone NOT NULL,
				expires timestamp with time zone NOT NULL,
				CONSTRAINT password_resets_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;

			ALTER TABLE secret.password_resets
				OWNER to postgres;

			CREATE INDEX fki_password_resets_user_id_fkey
				ON secret.password_resets USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "totp",
			CreateStatement: `CREATE TABLE secret.totp
			(
				user_id uuid NOT NULL,
				secret text COLLATE pg_catalog."default",
				key bytea,
				confirmed boolean,
				CONSTRAINT user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE secret.totp
				OWNER to postgres;
			
			CREATE INDEX fki_user_id_fkey
				ON secret.totp USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
