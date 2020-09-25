pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

library safeMath {
    function safeMul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function safeDiv(uint a, uint b) internal pure returns (uint) {
        assert(b > 0);
        uint c = a / b;
        assert(a == b * c + a % b);
        return c;
    }

    function safeSub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    function safeAdd(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c>=a && c>=b);
        return c;
    }
}

contract Owner {
    address public owner;
    bool public finished = false;
    uint starting;
    uint deadline;      // 候选人激活截止期 合约部署以后3天
    address to;  // 指定人
    string public name;
    string public introduction;
    uint fare; //选票面值
    uint minPar; //最低起投
    uint maxPar; //最高可投

    constructor() public{
        owner = msg.sender;
    }

    modifier onlyOwner {
        require (msg.sender == owner);
        _;
    }

    function finishVote() public {
        require(msg.sender == to,"to address err");
        finished = true;
    }

    modifier OnlyFinishVote {
        require(finished == true);
        _;
    }

    modifier onlyUnfinishVote {
        require(finished == false);
        _;
    }
}

contract CouncilElections is Owner{
    using safeMath for uint;

    struct VoteInfo {
        string name;
        uint vote;
        string source;
        // bool activated;
    }
    VoteInfo[] public Nominees;//候选人

    mapping(string => uint) addrLocation;
    mapping(string => mapping(address => uint)) Votes;

    address[] public Councils;
    address[] public NotCouncils;

    event LogNewNominees(string nominee,string from,uint num);
    event LogActivateNominee(address nominee,uint num);
    event LogVoteNominee(address voter,string council,uint votes);
    event LogNewCouncil(address council,string source);
    event LogTransfer(address voter,uint value);

    constructor(string memory _name,uint _begin,uint _end,string memory _info,address _to,uint _fare,uint _minPar,uint _maxPar,string[] memory nominees,string memory source)payable public {
        name = _name;
        starting = _begin;
        deadline = _end;
        introduction = _info;
        to=_to;
        fare = _fare;
        minPar = _minPar;
        maxPar = _maxPar;
        for(uint i=0;i<nominees.length;i++){
            VoteInfo memory info = VoteInfo(nominees[i],0,source);
            Nominees.push(info);
            addrLocation[nominees[i]] = Nominees.length;
            emit LogNewNominees(nominees[i],source,addrLocation[nominees[i]]-1);
        }
    }

    //查询候选人的信息，是否激活，得票数，来源
    function GetNomineeInfo(string memory name)public view returns(uint vote,string memory source, uint num){
        require(addrLocation[name]>0,"not name");
        VoteInfo memory info =Nominees[addrLocation[name]-1];
        return (info.vote,info.source,addrLocation[name]-1);

    }


    function StartVote(uint num) public payable onlyUnfinishVote {
        require(now >= starting && now < deadline,"out of time");
        VoteNominee(num);

    }

    function VoteNominee(uint num) internal {
        VoteInfo memory info = Nominees[num];
        //uint value = msg.value.safeDiv(10 ** 18);
        uint shan = msg.value.safeDiv(fare);
        require(shan >= minPar && shan<= maxPar-Votes[info.name][msg.sender],"out of range");
        require(Votes[info.name][msg.sender] + shan <= maxPar,"more then maxPar");

        info.vote += shan;
        Nominees[num] = info;
        Votes[info.name][msg.sender] += shan; //投向某候选项最多128票，允许多选
        emit LogVoteNominee(msg.sender,info.name,shan);
    }

    //有owner发起， 对未当选理事的投票人退sipc，需要外部记录投票的event，批量发起退款
    function safeWithdrawal(address payable voter,uint num) onlyOwner OnlyFinishVote public{
        uint allSipc=0;
        allSipc=num.safeMul(fare);
        voter.transfer(allSipc);
        emit LogTransfer(voter,allSipc);
    }

    //向基金会转账
    function transferToFoundation(address payable foundation) onlyOwner OnlyFinishVote public{
        foundation.transfer(address(this).balance);
    }

    //查询投票人对候选人的投票数
    function GetVote(string memory nominee,address voter) public view returns(uint){
        return Votes[nominee][voter];
    }

    function withdraw() OnlyFinishVote onlyOwner public {
        msg.sender.transfer(address(this).balance);
    }
}