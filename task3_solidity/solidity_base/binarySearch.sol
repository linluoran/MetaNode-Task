// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract binarySearch {
    function search(uint256[] memory arr, uint256 target)
        external
        pure
        returns (uint256)
    {
        // 空数组直接返回无效索引
        if (arr.length == 0) {
            return type(uint256).max; // 返回最大值表示未找到
        }

        uint256 left = 0;
        uint256 right = arr.length - 1;

        // 循环直到搜索范围缩小至0
        while (left <= right) {
            // 防溢出计算中间位置
            uint256 mid = left + (right - left) / 2;

            if (arr[mid] == target) {
                return mid; // 找到目标
            } else if (arr[mid] < target) {
                left = mid + 1; // 目标在右半区 [7](@ref)
            } else {
                // 避免下溢：当mid=0时right=0-1会下溢
                if (mid == 0) break;
                right = mid - 1; // 目标在左半区
            }
        }
        return type(uint256).max; // 未找到
    }
}
