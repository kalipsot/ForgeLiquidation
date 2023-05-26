// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";


import "https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPool.sol";
import "https://github.com/aave/aave-v3-core/blob/master/contracts/flashloan/base/FlashLoanSimpleReceiverBase.sol";



contract AaveFlashloan is FlashLoanSimpleReceiverBase {

    using SafeERC20 for IERC20;
    address public owner ;
    constructor(IPoolAddressesProvider provider)
        FlashLoanSimpleReceiverBase(provider)
    {
        owner = msg.sender;
    }

    modifier IsAutorised {
      require(msg.sender == owner || msg.sender == address(POOL));
      _;
   }

    function aaveFlashloan(address loanToken, uint256 loanAmount, bytes calldata params) IsAutorised external {
        IPool(address(POOL)).flashLoanSimple(
            address(this),
            loanToken,
            loanAmount,
            params,
            0
        );
    }

    function executeOperation(
        address asset,
        uint256 amount,
        uint256 premium,
        address, // initiator
        bytes calldata _params
    ) IsAutorised public override returns (bool) {
        require(
            amount <= IERC20(asset).balanceOf(address(this)),
            "Invalid balance for the contract"
        );
        address[] memory tos ;
        bytes[] memory data ;

        (tos, data) = abi.decode(_params, (address[], bytes[]));

        swap(tos, data);

        // pay back the loan amount and the premium (flashloan fee)
        uint256 amountToReturn = amount + premium;
        uint256 tosend = IERC20(asset).balanceOf(address(this)) - amountToReturn;
        IERC20(asset).transfer(owner,tosend);
        require(
            IERC20(asset).balanceOf(address(this)) >= amountToReturn,
            "Not enough amount to return loan"
        );

        approveToken(asset, address(POOL), amountToReturn);

        return true;
    }

    function approveToken(
        address token,
        address to,
        uint256 amountIn
    ) internal {
        require(IERC20(token).approve(to, amountIn), "approve failed");
    }


      function swap(address[] memory tos, bytes[] memory data) internal  {
        require(tos.length > 0 && tos.length == data.length, "Invalid input");

        for(uint256 i; i < tos.length; i++) {
        (bool success,bytes memory returndata) = tos[i].call{value: address(this).balance, gas: gasleft()}(data[i]);
        require(success, string(returndata));
        }
    }

    function getTime() external view returns (uint256) {
    return block.timestamp + 3 hours ;
     }
}