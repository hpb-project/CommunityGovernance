pragma solidity >=0.7.0 <0.9.0;


contract govermentSet {
    struct BlockNumber {
        uint256 number;
        bool valid;
    }

    address private owner;
    // key => value  (key不能重复)
    mapping(string => BlockNumber) map;  // 字典（mapping）类型   (映射类型)

    // event for EVM logging
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    // modifier to check if caller is owner
    modifier isOwner() {
        // If the first argument of 'require' evaluates to 'false', execution terminates and all
        // changes to the state and to Ether balances are reverted.
        // This used to consume all gas in old EVM versions, but not anymore.
        // It is often a good idea to use 'require' to check if functions are called correctly.
        // As a second argument, you can also provide an explanation about what went wrong.
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }

    /**
     * @dev Change owner
     * @param newOwner address of new owner
     */
    function changeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    function getOwner() external view returns (address) {
        return owner;
    }

    function setValue(string memory key, uint256 value) public isOwner {
        map[key]  = BlockNumber(value, true);
    }

    function getValue(string memory key) external view returns (uint256) {
        require(map[key].valid);
        return map[key].number;
    }


}
