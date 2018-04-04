pragma solidity ^0.4.21;

contract ConsumeContent {
    address public publisher;
    address public platform;
    // price of the content
    uint public price; 
    // ratio to split the fee. 
    uint8 public ratio; 
    // total times of conent being consumed
    uint public count;
    mapping (address=>uint) deposits;
    
    function ConsumeContent(address pPublisher, address pPlatform, uint pPrice, uint8 pRatio) public{
        publisher = pPublisher;
        platform = pPlatform;
        price = pPrice;
        ratio = pRatio;
    }

    // user consume the content and pay for it
    function consume() public payable {
        // don't accept less or more than the price
        require(msg.value == price);

        count++;
        //split and transfer the fee
        uint256 toPublisher = msg.value * ratio / 100;
        deposits[publisher] += toPublisher;
        deposits[platform] += msg.value - toPublisher; 
    }

    //withdraw each one's own
    function withDraw() public{
        require(deposits[msg.sender] > 0);

        uint amount = deposits[msg.sender];
        deposits[msg.sender] = 0;
        msg.sender.transfer(amount);
    }
}