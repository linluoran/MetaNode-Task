// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;


/*
上传图片到 https://app.pinata.cloud
获取链接 http://网关/ipfs/图片CID 写入到 metadata.json
上传 metadata.json 文件到 https://app.pinata.cloud
获取链接  http://网关/ipfs/json CID

测试步骤 通过 Remix web 部署到钱包的测试网络上
调用 mintNFT 铸造, recipient 为自己的钱包地址  tokenURI 为 metadata.json 的链接地址
铸造成功后就可以在 钱包内看到 NTF
opensea 测试网已经下线了 无法查看
只能通过 Etherscan 测试网 进行查看 https://sepolia.etherscan.io/ 无法从浏览器直接查看 NTF 图片等信息

approve 授权另一个地址（操作员）管理你的NFT
mintNFT 铸造新NFT
safeTransfer 安全转账（2种形式）
    形式1：safeTransferFrom(from, to, tokenId)
    形式2：safeTransferFrom(from, to, tokenId, data)
setApproval 批量授权/取消授权
transferFrom 直接NFT转账 from(发送方), to(接收方), tokenId 不检查接收方能否处理NFT（比safeTransfer风险高）
balanceOf 查询地址持有的NFT数量
isApprovedForAll 检查批量授权状态
name 返回NFT集合名称
ownerOf 查询NFT所有者
supportsInterface(interfaceId) 检查合约是否支持某接口（如ERC721）
symbol 返回NFT代币符号
tokenURI 获取NFT元数据链接
*/
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.3/contracts/utils/Counters.sol";

contract MyNTF is ERC721URIStorage {
    // 计数器初始化
    // 将计数器库的方法（如increment()）附加到_tokenIds变量上
    using Counters for Counters.Counter;
    // 私有状态变量，记录当前已铸造的NFT数量，每次铸造时自增，确保每个NFT的ID唯一
    Counters.Counter private _tokenIds;

    // 初始化NFT集合名称（Chitanda Eru-NFT）和符号（CE-NTF）
    constructor() ERC721("Chitanda Eru-NFT", "CE-NTF") {}

    // NFT铸造函数
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
