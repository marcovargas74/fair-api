package entity

import (
	"time"
)

//TODO Criar a Entidade Feira em Varias outras como Endereço subPrefeituras
//Fair Entity data
type Fair struct {
	ID           ID
	Name         string
	District     string
	Region5      string
	Neighborhood string

	/*
		//Endereço TODO criar entidade
		Logradouro string
		Numero     int
		Referencia string
		Longitude  string
		Latitude   string

		//Subprefeitura TODO criar entidade
		Setor            string
		Area             string
		CodDistrito      string
		CodSubPrefeitura string
		SubPrefeitura    string
		Regiao8          string
		NumRegistro      string*/
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewFair create a new Fair
func NewFair(name string, district string, region5 string, neighborhood string) (*Fair, error) {
	f := &Fair{
		ID:           NewID(),
		Name:         name,
		District:     district,
		Region5:      region5,
		Neighborhood: neighborhood,
		CreatedAt:    time.Now(),
	}
	err := f.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return f, nil
}

//Validate validate book
func (f *Fair) Validate() error {
	if f.Name == "" || f.District == "" || f.Region5 == "" || f.Neighborhood == "" {
		return ErrInvalidEntity
	}
	return nil
}

/*
	Name         string 30
	District     string 18
	Region5      string 6
	Neighborhood string 20

Nome da variável,Descrição da variável,Fonte,Tipo,Tamanho

ID,Identificador,Número de identificação do estabelecimento georreferenciado por SMDU/Deinfo,SMDU/Deinfo,TEXTO,08
NOME_FEIRA,Nome da feira livre,Denominação da feira livre atribuída pela Supervisão de Abastecimento,SMSP/AB,TEXTO,30


DEINFO_AB_FEIRASLIVRES_2014,LONG,Longitude,"Longitude da localização do estabelecimento no território do Município, conforme MDC",SMDU/Deinfo,TEXTO,10

DEINFO_AB_FEIRASLIVRES_2014,LAT,Latitude,"Latitude da localização do estabelecimento no território do Município, conforme MDC",SMDU/Deinfo,TEXTO,10
DEINFO_AB_FEIRASLIVRES_2014,SETCENS,Setor censitário,Setor censitário conforme IBGE,IBGE,TEXTO,15

DEINFO_AB_FEIRASLIVRES_2014,AREAP,Área de ponderação,Área de ponderação (agrupamento de setores censitários) conforme IBGE 2010,IBGE,TEXTO,13
DEINFO_AB_FEIRASLIVRES_2014,CODDIST,Código do distrito conforme IBGE,Código do Distrito Municipal conforme IBGE,IBGE,TEXTO,09

DEINFO_AB_FEIRASLIVRES_2014,DISTRITO,Distrito municipal,Nome do Distrito Municipal,SMDU/Deinfo,TEXTO,18
DEINFO_AB_FEIRASLIVRES_2014,CODSUBPREF,Código da Subprefeitura,Código de cada uma das 31 Subprefeituras (2003 a 2012),SMDU/Deinfo,TEXTO,02

DEINFO_AB_FEIRASLIVRES_2014,SUBPREF,Subprefeitura,Nome da Subprefeitura (31 de 2003 até 2012),SMDU/Deinfo,TEXTO,25
DEINFO_AB_FEIRASLIVRES_2014,REGIAO5,Região conforme divisão do Município em 5 áreas,Região conforme divisão do Município em 5 áreas,SMDU/Deinfo,TEXTO,06

DEINFO_AB_FEIRASLIVRES_2014,REGIAO8,Região conforme divisão do Município em 8 áreas,Região conforme divisão do Município em 8 áreas,SMDU/Deinfo,TEXTO,07
DEINFO_AB_FEIRASLIVRES_2014,NOME_FEIRA,Nome da feira livre,Denominação da feira livre atribuída pela Supervisão de Abastecimento,SMSP/AB,TEXTO,30

DEINFO_AB_FEIRASLIVRES_2014,REGISTRO,Registro da feira livre,Número do registro da feira livre na PMSP,SMSP/AB,TEXTO,6
DEINFO_AB_FEIRASLIVRES_2014,LOGRADOURO,Logradouro,Nome do logradouro onde se localiza a feira livre,SMSP/AB,TEXTO,34
DEINFO_AB_FEIRASLIVRES_2014,NUMERO,Número,Um número do logradouro onde se localiza a feira livre,SMSP/AB,TEXTO,5
DEINFO_AB_FEIRASLIVRES_2014,BAIRRO,Bairro,Bairro de localização da feira livre,SMSP/AB,TEXTO,20
DEINFO_AB_FEIRASLIVRES_2014,REFERENCIA,Ponto de referência,Ponto de referência da localização da feira livre,SMSP/AB,TEXTO,24
*/
