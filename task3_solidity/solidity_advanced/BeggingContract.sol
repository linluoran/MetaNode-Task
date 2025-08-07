// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 导入 OpenZeppelin
// Ownable: 所有权, PullPayment: 安全提款,  EnumerableSet: 高效管理集合
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/PullPayment.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract BeggingContract is Ownable, PullPayment {
    // 使用 EnumerableSet
    using EnumerableSet for EnumerableSet.AddressSet;

    // 捐赠者结构体
    struct Donor {
        address donorAddress;
        uint256 amount;
    }

    // 事件: 捐款 提款 限制捐款时间
    event Donation(address indexed donor, uint256 amount);
    event Withdrawal(address indexed owner, uint256 amount);
    event DonationWindowUpdated(
        uint8 startHour,
        uint8 startMinute,
        uint8 endHour,
        uint8 endMinute
    );

    // 错误
    // 参数错误
    error InvalidParam();

    // 记录用户捐赠金额
    mapping(address => uint256) public donationAmount;

    // 所有捐赠者集合
    EnumerableSet.AddressSet private allDonors;

    // 前 3 名捐赠者
    Donor[3] public topDonors;

    // 每日捐赠时间窗口 (UTC时间)
    uint8 public startHour; // 开始小时 (0-23)
    uint8 public startMinute; // 开始分钟 (0-59)
    uint8 public endHour; // 结束小时 (0-23)
    uint8 public endMinute; // 结束分钟 (0-59)
    bool public timeRestrictionEnabled;

    // 初始化用户所有权
    constructor() {
        // 设置合约部署者为所有者
        transferOwnership(msg.sender);
    }

    // 设置时间限制
    function setTimeLimit(
        uint8 _startHour,
        uint8 _startMinute,
        uint8 _endHour,
        uint8 _endMinute
    ) public onlyOwner returns (bool success) {
        // 校验时间格式
        if (_startHour > 23 || _startMinute > 59) revert InvalidParam();
        if (_endHour > 23 || _endMinute > 59) revert InvalidParam();

        // 校验结束时间 > 开始时间
        uint256 startTotalMinutes = _startHour * 60 + _startMinute;
        uint256 endTotalMinutes = _endHour * 60 + _endMinute;
        if (endTotalMinutes <= startTotalMinutes) revert InvalidParam();

        startHour = _startHour;
        startMinute = _startMinute;
        endHour = _endHour;
        endMinute = _endMinute;
        success = true;
    }

    // function donate payable () public returns(bool success){

    // }
    receive() external payable {}
}
