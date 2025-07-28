// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract IntegerToRoman {
    // 预定义罗马数字值及符号（从大到小排序）
    uint16[] private _values = [
        1000,
        900,
        500,
        400,
        100,
        90,
        50,
        40,
        10,
        9,
        5,
        4,
        1
    ];
    string[] private _symbols = [
        "M",
        "CM",
        "D",
        "CD",
        "C",
        "XC",
        "L",
        "XL",
        "X",
        "IX",
        "V",
        "IV",
        "I"
    ];

    // 核心转换函数（Gas 优化版）
    function toRoman(uint256 num) public view returns (string memory) {
        require(num >= 1 && num <= 3999, "Invalid input: 1-3999");
        bytes memory roman;
        uint256 remaining = num;

        // 贪心算法：从最大单位开始匹配
        for (uint256 i = 0; i < _values.length; ) {
            while (remaining >= _values[i]) {
                // 动态扩展字节数组存储罗马符号
                bytes memory symbol = bytes(_symbols[i]);
                roman = bytes.concat(roman, symbol);
                remaining -= _values[i];
            }
            // unchecked 减少循环 Gas 消耗
            unchecked {
                i++;
            }
        }
        return string(roman);
    }
}
