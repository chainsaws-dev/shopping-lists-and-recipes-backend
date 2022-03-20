// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "github.com/jackc/pgx/v4/pgxpool"

// PostgreSQLCreateTablesPublic - создаёт таблицы для схемы public (для рецептов и списка покупок)
func PostgreSQLCreateTablesPublic(dbc *pgxpool.Pool) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "Files",
			CreateStatement: `CREATE TABLE public."Files"
			(
				id bigserial NOT NULL,
				filename character varying(255) COLLATE pg_catalog."default",
				filesize bigint,
				filetype character varying(50) COLLATE pg_catalog."default",
				file_id character varying(50) COLLATE pg_catalog."default",
				preview_id character varying(50) COLLATE pg_catalog."default",
				CONSTRAINT "Files_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."Files"
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "Ingredients",
			CreateStatement: `CREATE TABLE public."Ingredients"
			(
				id bigserial NOT NULL,
				name character varying(100) COLLATE pg_catalog."default" NOT NULL,
				CONSTRAINT "Ingredients_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."Ingredients"
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "Recipes",
			CreateStatement: `CREATE TABLE public."Recipes"
			(
				id bigserial NOT NULL,
				name character varying(100) COLLATE pg_catalog."default",
				description text COLLATE pg_catalog."default",
				image_id bigint,
				CONSTRAINT "Recipes_pkey" PRIMARY KEY (id),
				CONSTRAINT "Recipes_image_id_fkey" FOREIGN KEY (image_id)
					REFERENCES public."Files" (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."Recipes"
				OWNER to postgres;
			
			CREATE INDEX "fki_Recipes_image_id_fkey"
				ON public."Recipes" USING btree
				(image_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "RecipesIngredients",
			CreateStatement: `CREATE TABLE public."RecipesIngredients"
			(
				recipe_id bigint NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL,
				CONSTRAINT "RecipesIngredients_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
					REFERENCES public."Ingredients" (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE SET NULL,
				CONSTRAINT "RecipesIngredients_recipe_id_fkey" FOREIGN KEY (recipe_id)
					REFERENCES public."Recipes" (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE SET NULL
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."RecipesIngredients"
				OWNER to postgres;
			
			CREATE INDEX "fki_RecipesIngredients_ingredient_id_fkey"
				ON public."RecipesIngredients" USING btree
				(ingredient_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "ShoppingList",
			CreateStatement: `CREATE TABLE public."ShoppingList"
			(
				id bigserial NOT NULL,
				ingredient_id bigint NOT NULL,
				quantity bigint NOT NULL,
				CONSTRAINT "ShoppingList_pkey" PRIMARY KEY (id),
				CONSTRAINT "ShoppingList_ingredient_id_fkey" FOREIGN KEY (ingredient_id)
					REFERENCES public."Ingredients" (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE SET NULL
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE public."ShoppingList"
				OWNER to postgres;
			
			CREATE INDEX "fki_ShoppingList_ingredient_id_fkey"
				ON public."ShoppingList" USING btree
				(ingredient_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
