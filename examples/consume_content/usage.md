To deploy a new contract, run without `--addrHex` parameter,
```
go run main.go --ethURL http://[IP]:8545 --ownerKeyFile [path_to_key_file] --ownerKeyPwd [passworkd] --platformKeyFile [path_to_key_file] --platformKeyPwd [password] --consumerKeyFile [path_to_key_file] --consumerKeyPwd [password]
```


To to call `count` of the contract, which needs **no** transaction. `--addrHex` is the address of a contract already deoplyed and confirmed on chian,

```
go run main.go --ethURL http://[IP]:8545 --addrHex [address_of_contract_deployed] --ownerKeyFile [path_to_key_file] --ownerKeyPwd [passworkd] --platformKeyFile [path_to_key_file] --platformKeyPwd [password] --consumerKeyFile [path_to_key_file] --consumerKeyPwd [password]
```


To to call `consume` of the contract, which needs consumer sign a transaction. `--addrHex` is the address of a contract already deoplyed and confirmed on chian. `--doConsume` means run the consuming logic,

```
go run main.go --ethURL http://[IP]:8545 --doconsume true --addrHex [address_of_contract_deployed] --ownerKeyFile [path_to_key_file] --ownerKeyPwd [passworkd] --platformKeyFile [path_to_key_file] --platformKeyPwd [password] --consumerKeyFile [path_to_key_file] --consumerKeyPwd [password]
```