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
    string public name;
    string public introduction;
    
    constructor() public{
        owner = msg.sender;
    }

    modifier onlyOwner {
        require (msg.sender == owner);
        _;
    }

    function finishVote() onlyOwner public {
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
        uint id;
        uint vote;
        string source;
    }
    VoteInfo[] public Nominees;//候选项
    
    constructor(string memory _name,uint _begin,uint _end,string memory _info) public {
        name = _name;
        starting = _begin;
        deadline = _end;
        introduction = _info;
    }

    //First 手动按序输入候选人名单，依次为B.VIP,SIPC.VIP,Normal
    function AddVoteInfo(string[] memory nominees)onlyOwner public {
        uint j = Nominees.length;
        for (uint i=0;i<nominees.length;i++){
            VoteInfo memory info = VoteInfo(j+i,0,nominees[i]); 
            Nominees.push(info);
        }
    }

    //只由owner提供链上投票
    function SendVote(uint _id) onlyUnfinishVote onlyOwner public {
        require(now > starting && now < deadline,"out of time");
        require(_id >= 0 && _id <= Nominees.length,"uncorrect number");
        VoteNominee(_id);
    }
    
    function VoteNominee(uint num) internal {
        VoteInfo memory info = Nominees[num];
        info.vote += 1;
        Nominees[num] = info;
    }
} 