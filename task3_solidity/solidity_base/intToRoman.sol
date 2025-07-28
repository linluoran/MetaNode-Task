// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract RomanToInteger {
    // 罗马字符到数值的映射（Gas 优化：用 bytes1 代替 string）
    function charValue(bytes1 c) internal pure returns (uint16) {
        if (c == "I") return 1;
        if (c == "V") return 5;
        if (c == "X") return 10;
        if (c == "L") return 50;
        if (c == "C") return 100;
        if (c == "D") return 500;
        if (c == "M") return 1000;
        revert("Invalid Roman character");
    }

    // 核心转换函数（一次遍历+左减右加规则）
    function romanToInt(string memory roman) public pure returns (uint256) {
        bytes memory romanBytes = bytes(roman);
        uint256 length = romanBytes.length;
        uint256 total = 0;

        for (uint256 i = 0; i < length; ) {
            uint256 current = charValue(romanBytes[i]);

            // 检查下一个字符是否更大（需减法）
            if (i + 1 < length) {
                uint256 next = charValue(romanBytes[i + 1]);
                if (current < next) {
                    total += (next - current);
                    i += 2; // 跳过已处理的两个字符
                    continue;
                }
            }

            total += current;
            unchecked {
                i++;
            } // Gas 优化：关闭溢出检查
        }
        return total;
    }
}
