pragma solidity ^0.5.1;
pragma experimental ABIEncoderV2;

contract gasdiscountcache{
    address superadmin;
    address[] gasaddrs;
    constructor() public{
        superadmin = msg.sender;
    }
    function addgasaddr(address gasaddr) public{
        require(superadmin == msg.sender,"not admin");
        gasaddrs.push(gasaddr);
    }
    function getaddrs() public returns(address[] memory){
        return gasaddrs;
    }
    function changeSuperAdmin(address spadmin) public{
        require(superadmin == msg.sender,"Not superadmin");
        superadmin = spadmin;
    }
}
