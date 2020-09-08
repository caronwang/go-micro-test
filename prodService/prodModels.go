package prodService

import "strconv"

type ProdModel struct {

	ProdId int
	ProdName string
}

func NewProd(id int,name string) *ProdModel{
	return &ProdModel{
		id,name,
	}
}

func NewProdList(n int) []*ProdModel{
	ret := make([]*ProdModel,0)
	for i:=0;i<n;i++{
		ret= append(ret,NewProd(100+i,"prod_"+strconv.Itoa(100+i)))
	}

	return ret
}
