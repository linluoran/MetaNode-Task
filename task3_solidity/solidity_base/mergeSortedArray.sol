// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract mergeSortedArray {
    function merge(int256[] memory arr1, int256[] memory arr2)
        external
        pure
        returns (int256[] memory)
    {
        uint256 m = arr1.length;
        uint256 n = arr2.length;
        int256[] memory merged = new int256[](m + n); // 创建新数组
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        // 正向双指针合并
        while (i < m && j < n) {
            merged[k++] = (arr1[i] < arr2[j]) ? arr1[i++] : arr2[j++];
        }
        // 处理剩余元素
        while (i < m) merged[k++] = arr1[i++];
        while (j < n) merged[k++] = arr2[j++];
        return merged;
    }
}
