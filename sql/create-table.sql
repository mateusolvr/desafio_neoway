CREATE TABLE IF NOT EXISTS public.analise_compra_usuario (
	cpf varchar(25) NOT NULL,
	private bool NOT NULL,
	incompleto bool NOT NULL,
	data_ultima_compra date NULL,
	ticket_medio float8 NULL,
	ticket_ultima_compra float8 NULL,
	loja_mais_frequente varchar(25) NULL,
	loja_ultima_compra varchar(25) NULL,
	inserido_em timestamp NOT NULL,
	cpf_valido bool NOT NULL,
	cnpj_mais_frequente_valido bool NOT NULL,
	cnpj_ultima_compra_valido bool NOT NULL
);
