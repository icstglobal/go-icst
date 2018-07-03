DROP SCHEMA IF EXISTS `icst`;
CREATE SCHEMA IF NOT EXISTS `icst` CHARACTER SET utf8;

DROP TABLE IF EXISTS `icst`.`Account`;
CREATE TABLE `icst`.`Account` (
  `accountID`   char(36)    NOT NULL COMMENT '账户ID',
  `walletID`    char(32)    COMMENT '钱包ID',
  `chainType`   INT         NOT NULL COMMENT '链的类型',
  `pubKey`      varchar(100)   NOT NULL COMMENT '链的类型',
  --  `balance`     varchar(50) COMMENT '余额',
  `createTime`  timestamp   NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`accountID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `icst`.`Transaction`;
CREATE TABLE `icst`.`Transaction` (
  `txHash`     varchar(100)  	NOT NULL COMMENT '交易hash',
  `accountID`  char(36) 	    COMMENT '账户ID',
  `txType`     TINYINT  		NOT NULL COMMENT '交易类型',
  `blockHeight`varchar(50)  	NOT NULL COMMENT '区块高度',
  `timestamp`	timestamp  		NOT NULL COMMENT '区块编号',
  `from`		varchar(50)  	NOT NULL COMMENT 'from地址',
  `to`			varchar(50)  			 COMMENT 'to地址',
  `gasPrice`	varchar(50)  	NOT NULL COMMENT 'gas价格',
  `gasUsed`		varchar(50)  	NOT NULL COMMENT 'gas用量',
  `gasLimited`	varchar(50)  	NOT NULL COMMENT 'gas上限',
  `hash`		varchar(100)  	NOT NULL COMMENT '交易hash值',
  `value`		varchar(50)  	NOT NULL COMMENT '交易eth值',
  `nonce`		varchar(50)  	NOT NULL COMMENT '随机值',
  `cxAddr`		varchar(50)  			 COMMENT '合约地址',

  PRIMARY KEY (`txHash`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `icst`.`ContentContract`;
CREATE TABLE `icst`.`ContentContract` (
  `cxAddr`     	varchar(50)  	NOT NULL COMMENT '合约地址',
  `accountID`   char(36) 	    COMMENT '账户ID',
  `txHash`		varchar(100)  	NOT NULL COMMENT '交易hash值',
  `publisher`	varchar(50)  	NOT NULL COMMENT '创作者',
  `platform`	varchar(50)  	NOT NULL COMMENT '平台',
  `price`		varchar(50)  	NOT NULL COMMENT '内容价格',
  `ratio`		INT  			NOT NULL COMMENT '分成',
  `count`		varchar(50)  	NOT NULL COMMENT '支付次数',
  PRIMARY KEY (`cxAddr`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `icst`.`ContentContractDeposit`;
CREATE TABLE `icst`.`ContentContractDeposit` (
  `id` 			INT 			NOT NULL AUTO_INCREMENT,
  `cxAddr`     	varchar(50)  	NOT NULL COMMENT '合约地址',
  `addr`		varchar(100)  	NOT NULL COMMENT '用户addr',
  `value`		varchar(50)  	NOT NULL COMMENT '用户余额',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
