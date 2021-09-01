pragma solidity ^0.5.1;

contract HRNG {
    function hrand() public view returns (bytes32) {
         return block.random;
    }
}