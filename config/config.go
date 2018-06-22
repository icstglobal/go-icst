package config
var Conf conf

type conf struct{
	// connect chain url 
	ChainUrl string

	// chain type: Eos or Eth
	ChainType int
}


func InitConfig(chainUrl string, chainType int){
	Conf = conf{
		ChainUrl: chainUrl,
		ChainType: chainType,
	}
}
