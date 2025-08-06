// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract MyNTF is ERC721URIStorage {
    // 计数器初始化
    // 将计数器库的方法（如increment()）附加到_tokenIds变量上
    using Counters for Counters.Counter;
    // 私有状态变量，记录当前已铸造的NFT数量，每次铸造时自增，确保每个NFT的ID唯一
    Counters.Counter private _tokenIds;

    // 初始化NFT集合名称（Chitanda Eru-NFT）和符号（CE-NTF）
    constructor() ERC721("Chitanda Eru-NFT", "CE-NTF") {}

    // ​NFT铸造函数​
    function mintNFT(address recipient, string memory tokenURI)
        public
        returns (uint256)
    {
        //  增加计数器值（如从0→1）
        _tokenIds.increment();
        // 获取新ID（如ID=1）
        uint256 newTokenId = _tokenIds.current();
        // 将NFT铸造给接收者地址
        _mint(recipient, newTokenId);
        // 绑定ID与元数据URI
        _setTokenURI(newTokenId, tokenURI);
        // 返回新NFT的ID
        return newTokenId;
    }
}
