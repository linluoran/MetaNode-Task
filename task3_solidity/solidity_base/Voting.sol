// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract  Voting{

    // 候选人结构体
    struct Candidate{
        uint id;
        string name;
        uint voteCount;
    }

    // 候选人index => 候选人  
    mapping(uint voteIndex => Candidate) public idToCandidate;

    // 

}


