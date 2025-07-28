// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Voting {
    // 初始化候选人
    constructor() {
        _addCandidate("HelloA");
        _addCandidate("HelloB");
        _addCandidate("HelloC");
    }

    // 候选人数量
    uint32 public candidateCount;

    // 候选人结构体
    struct Candidate {
        // 使用uint32 最大支持42亿
        uint32 id;
        string name;
        uint32 voteCount;
    }

    // 候选人索引 0,1,2 => 候选人
    mapping(uint256 => Candidate) public idToCandidate;

    // 是否投过票 不允许重复投票
    mapping(address => bool) public voterToVoted;

    // 不允许重复添加候选人
    mapping(string => bool) public candidateExists;

    // 自定义重复投票错误
    error AlreadyVoted(); // 已经投过票了
    error InvalidCandidate(); // 候选人不存在

    // 定义投票事件
    event Voted(address indexed voter, uint256 candidateId);

    // 添加候选人
    function _addCandidate(string memory _name) private {
        if (candidateExists[_name]) {
            revert InvalidCandidate();
        }

        candidateCount++;
        candidateExists[_name] = true;
        idToCandidate[candidateCount] = Candidate(candidateCount, _name, 0);
    }

    // 投票函数
    function vote(address voter, uint256 voteIndex) external {
        if (voterToVoted[voter]) {
            revert AlreadyVoted();
        }

        if (voteIndex < 0 || voteIndex > candidateCount) {
            revert InvalidCandidate();
        }

        voterToVoted[voter] = true;

        // 使用unchecked避免溢出检查
        unchecked {
            idToCandidate[voteIndex].voteCount++;
        }

        emit Voted(voter, voteIndex);
    }

    // 获取某个候选人投票数
    function getVotes(uint256 voteIndex)
        public
        view
        returns (uint256 voteCount)
    {
        if (voteIndex < 0 || voteIndex > candidateCount) {
            revert InvalidCandidate();
        }

        voteCount = idToCandidate[voteIndex].voteCount;
    }

    // 重置所有候选人投票
    function resetVotes() public {
        for (uint256 i = 0; i <= candidateCount; i++) {
            idToCandidate[i].voteCount = 0;
        }
    }
}
