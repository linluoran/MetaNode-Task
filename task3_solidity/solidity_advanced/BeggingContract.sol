// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 导入 OpenZeppelin
// Ownable: 所有权, PullPayment: 安全提款, ReentrancyGuard: 防止重入攻击,  EnumerableSet: 高效管理集合
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/access/Ownable.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/security/PullPayment.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/security/ReentrancyGuard.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/utils/structs/EnumerableSet.sol";

contract BeggingContract is Ownable, PullPayment, ReentrancyGuard {
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
        uint256 startHour,
        uint256 startMinute,
        uint256 endHour,
        uint256 endMinute
    );

    // 错误
    // 参数错误
    error InvalidParam();
    // 当前时间段禁止捐赠
    error DonationDisabledCurrently();
    // 捐赠金额需要大于 0
    error DonationGtZero();
    // 提取金额失败
    error WithdrawFailed();
    // 禁止绕过限制直接捐款
    error DirectTransfersDisabled();

    // 记录用户捐赠金额
    mapping(address => uint256) public donationAmount;

    // 所有捐赠者集合
    EnumerableSet.AddressSet private allDonors;

    // 前 3 名捐赠者
    Donor[3] public topDonors;

    // 每日捐赠时间窗口 (UTC时间)
    uint256 public startTime; // 记录每天开始分钟数 12:30 = 12 * 60  + 30
    uint256 public endTime; // 记录每天结束分钟数
    bool public timeRestrictionEnabled; // 是否启用时间限制
    uint256 public timezoneOffset = 8 hours; // 例如北京时间 UTC+8

    // 初始化用户所有权
    constructor() {
        // 设置合约部署者为所有者
        transferOwnership(msg.sender);
    }

    // 设置时间限制
    function setTimeLimit(
        uint256 startHour,
        uint256 startMinute,
        uint256 endHour,
        uint256 endMinute
    ) external onlyOwner nonReentrant returns (bool success) {
        // 校验时间格式
        if (startHour > 23 || startMinute > 59) revert InvalidParam();
        if (endHour > 23 || endMinute > 59) revert InvalidParam();

        // 校验结束时间 > 开始时间
        uint256 startTotalMinutes = startHour * 60 + startMinute;
        uint256 endTotalMinutes = endHour * 60 + endMinute;
        if (endTotalMinutes <= startTotalMinutes) revert InvalidParam();

        startTime = startTotalMinutes;
        endTime = endTotalMinutes;

        // 启用时间限制
        timeRestrictionEnabled = true;
        emit DonationWindowUpdated(startHour, startMinute, endHour, endMinute);
        return true;
    }

    // 关闭时间限制
    function disableTimeRestriction()
    external
    onlyOwner
    returns (bool success)
    {
        timeRestrictionEnabled = false;
        success = true;
    }

    // 计算现在过去多少分钟了
    function _getCurrentMinutes() private view returns (uint256) {
        return ((block.timestamp + timezoneOffset) % 1 days) / 60;
    }

    // 判断当前时间是否在开始结束时间内
    function isDonationWindowOpen() public view returns (bool) {
        uint256 currentMinutes = _getCurrentMinutes();
        return currentMinutes >= startTime && currentMinutes <= endTime;
    }

    // 捐赠函数 同时会检查是否在捐赠时间内
    function donate() external payable nonReentrant returns (bool success) {
        if (msg.value <= 0) revert DonationGtZero();

        if (timeRestrictionEnabled) {
            if (!isDonationWindowOpen()) {
                // 不在捐赠时间内 记录捐赠金额为 0
                emit Donation(msg.sender, 0);
                revert DonationDisabledCurrently();
            }
        }

        // 更新捐赠记录
        donationAmount[msg.sender] += msg.value;

        // 记录所有捐赠人
        allDonors.add(msg.sender);

        // 更新排行榜
        _updateTopDonors(msg.sender, donationAmount[msg.sender]);

        emit Donation(msg.sender, msg.value);
        return true;
    }

    // 获取捐赠者地址列表
    function getDonors(uint256 startIndex, uint256 count)
    external
    view
    returns (address[] memory)
    {
        uint256 total = allDonors.length();
        if (startIndex >= total) {
            return new address[](0);
        }

        uint256 endIndex = startIndex + count;
        if (endIndex > total) {
            endIndex = total;
        }

        address[] memory result = new address[](endIndex - startIndex);
        for (uint256 i = startIndex; i < endIndex; i++) {
            result[i - startIndex] = allDonors.at(i);
        }

        return result;
    }

    // 获取捐赠总人数
    function donorCount() external view returns (uint256) {
        return allDonors.length();
    }

    // 获取排行榜
    function getTopDonors() external view returns (Donor[3] memory) {
        return topDonors;
    }

    // 获取任意用户捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donationAmount[donor];
    }

    // 提取所有捐赠金额
    function withdraw() external onlyOwner nonReentrant {
        // 查询合约内的金额
        uint256 balance = address(this).balance;
        if (balance <= 0) revert DonationGtZero();

        // 安全转账给所有者
        Address.sendValue(payable(owner()), balance);
        emit Withdrawal(owner(), balance);
    }

    // -------------------------    内部函数    -------------------------
    // 更新排行榜
    function _updateTopDonors(address _donor, uint256 _amount) private {
        // 检查是否已存在
        for (uint256 i = 0; i < 3; i++) {
            if (topDonors[i].donorAddress == _donor) {
                // 用户已存在 更新金额排序后推出
                topDonors[i].amount = _amount;
                _sortDonors();
                return;
            }
        }

        // 是否超过最低金额
        if (_amount > topDonors[2].amount) {
            topDonors[2] = Donor({donorAddress: _donor, amount: _amount});
            _sortDonors();
        }
    }

    // 捐赠榜排序
    // 改为插入排序 (更适合小数据集)
    function _sortDonors() private {
        for (uint256 i = 1; i < 3; i++) {
            Donor memory temp = topDonors[i];
            uint256 j = i;
            while (j > 0 && temp.amount > topDonors[j - 1].amount) {
                topDonors[j] = topDonors[j - 1];
                j--;
            }
            topDonors[j] = temp;
        }
    }

    receive() external payable {
        // 禁止绕过限制捐款
        revert DirectTransfersDisabled();
    }
}
