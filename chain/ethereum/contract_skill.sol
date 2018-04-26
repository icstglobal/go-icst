pragma solidity ^0.4.21;

contract ConsumeSkill{
    address public publisher;
    address public platform;
    address public consumer;
    // price of the skill service 
    uint32 public price; 
    // ratio to split the fee. 
    uint8 public ratio; 
    // hash of the digital form of service contract, as hex string
    string public hash;
    mapping (address=>uint) balances;
    State public state;
    mapping (address=>bool) completeConfirmed;

    enum State {Init, InProgress, Complete, Aborted }

    event StateChange(address actor, State from, State to);
    
    modifier publisherOnly(){
        require(msg.sender == publisher);
        _;
    }

    modifier consumerOnly(){
        require(msg.sender == consumer);
        _;
    }

    modifier anyOf(address[2] users){
        require(msg.sender == users[0] || msg.sender == users[1]);
        _;
    }

    // a new contract for every producer-consumer pair
    function ConsumeSkill(string pHash, address pPublisher, address pPlatform, address pConsumer, uint32 pPrice, uint8 pRatio) public{
        hash = pHash;
        publisher = pPublisher;
        platform = pPlatform;
        consumer = pConsumer;
        price = pPrice;
        ratio = pRatio;
        state = State.Init;
    }

    // user start consuming the skill service and pay for it
    function start() public payable consumerOnly{
        if (state != State.Init){
            return;
        }
        // don't accept less or more than the pric
        require(msg.value == price);

        //split and transfer the fee
        uint256 toPublisher = msg.value * ratio / 100;
        balances[publisher] += toPublisher;
        balances[platform] += msg.value - toPublisher; 

        emit StateChange(msg.sender, State.Init, State.InProgress);
    }

    function Complete() public anyOf([publisher, consumer]) {
        if (state != State.InProgress){
            return;
        }

        completeConfirmed[msg.sender] = true;
        //need both confirmed by the publisher and consumer
        if (completeConfirmed[publisher] && completeConfirmed[consumer]){
            state = State.Complete;
        }

        emit StateChange(msg.sender, State.InProgress, State.Complete);
    }

    function abort() public publisherOnly {
        if (state != State.Init){
            return;
        }

        state = State.Complete;

        emit StateChange(msg.sender, State.Init, State.Aborted);
    }

    //withdraw each one's own
    function withDraw() public{
        if( state != State.Complete && state != State.Aborted){
            return;
        }

        require(balances[msg.sender] > 0);

        uint amount = balances[msg.sender];
        balances[msg.sender] = 0;
        msg.sender.transfer(amount);
    }

    function () public payable {
        //do nothing
    }
}