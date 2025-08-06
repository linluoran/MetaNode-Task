// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 参考 https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol

/*
测试步骤, 
1. 部署 MyERC20. owner 为部署账户, 初始化产生的币也在这个账户里
2. 复制部署好的 MyERC20 合约地址, 部署 TestERC20 合约.
3. 注意此时 TestERC20 合约有一个所有者地址就是部署人的地址. 你直接调用 TestERC20 msg.sender 就是这个调用人的地址. 当前就是部署人的地址.
4. 但是你通过 TestERC20 调用 MyERC20, 就会导致 MyERC20 中 msg.sender 为 TestERC20 合约地址.
5. 所以默认使用 TestERC20 合约地址作为被授权用户. 通过 MyERC20 进行授权给 TestERC20 合约账户额度, 再调用代扣额度, 随便给个 to 钱包地址即可.
6. 测试直接扣款 和 代币增发也是一样 随便给个 to 钱包地址即可.
*/

contract MyERC20 {
    // 代币元数据
    string public name; // 代币名称
    string public symbol; // 代币符号 如: MTK
    uint8 public decimals; // 代币小数位数
    uint256 public totalSupply; // 代币总供应量

    // 合约所有者地址
    address public owner;

    // 余额映射: 用户 => 余额
    mapping(address => uint256) private _balances;

    // 授权映射 类似亲密付: 用户 => (被授权用户 => 授权额度)
    mapping(address => mapping(address => uint256)) private _allowances;

    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    // 授权事件
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    // 错误定义
    error PermissionDenied(); // 权限不足
    error InvalidAddress(); // 非法地址
    error InsufficientBalance(); // 余额不足

    // 权限修饰器 仅所有者可以调用
    modifier onlyOwner() {
        if (msg.sender != owner) revert PermissionDenied();
        _;
    }

    // 构造函数
    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply
    ) {
        owner = msg.sender;
        name = _name;
        symbol = _symbol;
        decimals = _decimals;

        // 给自己铸币
        _mint(msg.sender, _initialSupply * (10**uint256(decimals)));
    }

    // 查询余额
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    // 授权额度
    function approve(address spender, uint256 amount) public returns (bool) {
        return _approve(spender, amount);
    }

    // 转账功能
    function transfer(address to, uint256 amount) public returns (bool) {
        return _transfer(to, amount);
    }

    // 代扣转账
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        return _transferFrom(from, to, amount);
    }

    // 额外增发代币
    function mint(address to, uint256 amount) public onlyOwner returns (bool) {
        return _mint(to, amount * (10**uint256(decimals)));
    }

    // -------------------------    内部函数    -------------------------

    // 转账函数
    function _transfer(address to, uint256 amount)
        private
        returns (bool success)
    {
        address from = msg.sender;
        if (from == to) revert InvalidAddress();
        if (from == address(0)) revert InvalidAddress();
        if (to == address(0)) revert InvalidAddress();

        uint256 fromAmount = _balances[from];
        if (fromAmount < amount) revert InsufficientBalance();

        // 这里不涉及公链转账 且不涉及总币数加减
        unchecked {
            _balances[from] -= amount;
            _balances[to] += amount;
        }

        success = true;
        emit Transfer(from, to, amount);
    }

    // 授权函数
    function _approve(address spender, uint256 amount)
        private
        returns (bool success)
    {
        address ownerTmp = msg.sender;
        if (ownerTmp == address(0)) revert InvalidAddress();
        if (spender == address(0)) revert InvalidAddress();

        // 授权时不检查用户额度
        _allowances[ownerTmp][spender] = amount;
        success = true;

        emit Approval(ownerTmp, spender, amount);
    }

    // 代扣转账
    function _transferFrom(
        address from,
        address to,
        uint256 amount
    ) private returns (bool success) {
        address spender = msg.sender;
        if (from == address(0)) revert InvalidAddress();
        if (to == address(0)) revert InvalidAddress();

        // 检查授权额度够不够
        uint256 allowance = _allowances[from][spender];
        if (allowance < amount) revert InsufficientBalance();

        // 检查from余额够不够
        uint256 balance = _balances[from];
        if (balance < amount) revert InsufficientBalance();

        _allowances[from][spender] -= amount;
        unchecked {
            _balances[from] -= amount;
            _balances[to] += amount;
        }
        success = true;
        emit Transfer(from, to, amount);
    }

    function _mint(address to, uint256 amount)
        internal
        onlyOwner
        returns (bool success)
    {
        address ownerTmp = msg.sender;
        if (ownerTmp == address(0)) revert InvalidAddress();

        unchecked {
            totalSupply += amount;
            _balances[to] += amount;
        }

        success = true;
        emit Transfer(address(0), to, amount);
    }
}

// 被授权用户
contract TestERC20 {
    MyERC20 public myERC20;

    constructor(address myAddress) {
        myERC20 = MyERC20(address(myAddress));
    }

    function getUserAddress() public view returns (address) {
        // 以合约地址作为用户地址
        return address(this);
    }

    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        return myERC20.transferFrom(from, to, amount);
    }
}
