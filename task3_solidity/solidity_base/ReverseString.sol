// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract StringReverser {
    // 识别字符字节长度（支持中文）
    function _getCharSize(bytes1 b) private pure returns (uint256) {
        // 检查首位比特：0 表示单字节，11 表示多字节首字节
        if (uint8(b) & 0x80 == 0) {
            return 1; // ASCII 字符
        } else if (uint8(b) & 0xE0 == 0xC0) {
            return 2; // 2 字节字符（如部分符号）
        } else {
            return 3; // 3 字节字符（如中文）[4,7](@ref)
        }
    }

    function stringReverser(string memory _inputStr)
        public
        pure
        returns (string memory newStr)
    {
        bytes memory inputBytes = bytes(_inputStr);
        uint256 length = inputBytes.length;
        if (length == 0) return _inputStr;

        bytes memory reversedBytes = new bytes(length);

        uint256 i = 0;
        uint256 j = length;

        while (i < length) {
            uint256 charSize = _getCharSize(inputBytes[i]);
            unchecked {
                j -= charSize;
            }
            for (uint256 k = 0; k < charSize; k++) {
                reversedBytes[j + k] = inputBytes[i + k];
            }
            unchecked {
                i += charSize;
            }
        }
        return string(reversedBytes);
    }
}
