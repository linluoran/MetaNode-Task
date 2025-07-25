// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract  Voting{

    // 候选人结构体
    struct Candidate{
        uint id;
        string name;
        uint voteCount;
    }

    // 候选人索引 0,1,2 => 候选人  
    mapping(uint voteIndex => Candidate) public idToCandidate;

    // 是否投过票 不允许重复投票
    mapping(addres voter => bool is voted) public voterToVoted
    
    // 投票函数
    function vote (){}

}



