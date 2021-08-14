pragma solidity ^0.5.1;


contract blockSet {
    struct BlockNumber {
        string  keywords;
        uint256 blockNumber;
        uint256 threshold;
        bool valid;
    }

    address private owner; // owner is always an admin.
    uint256 threshold;
    mapping(string => BlockNumber) blockmap; // keywords ==> blocknumber
    mapping(string => BlockNumber) history;
    address[]  admins;
    string[]   proposalList;
    mapping(address => bool) mapAdmin;
    mapping(string => address[]) votes;

    // event for EVM logging
    event OwnerSet(address indexed oldOwner, address indexed newOwner);
    event AddAdmin(address indexed user, address indexed newAdmin);
    event DelAdmin(address indexed user, address indexed admin);
    event PassedThreshold(string indexed keywords);

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
    modifier isAdmin() {
        require(mapAdmin[msg.sender] || (msg.sender == owner), "Caller is not owner and admin");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() public {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        threshold = 1;
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

    function addAdmin(address admin) public isOwner {
        if (!mapAdmin[admin]){
            mapAdmin[admin] = true;
            admins.push(admin);
            emit AddAdmin(msg.sender, admin);
        }
    }

    function getAdmins() public view returns (address[] memory){
        return admins;
    }

    function getThreshold() public view returns (uint256){
        return threshold;
    }

    function getProposalThreshold(string memory key) public view returns (uint256) {
        require(blockmap[key].threshold > 0, "proposol not exist");
        return blockmap[key].threshold;
    }

    function deleteAdmin(address admin) public isOwner {
        if (mapAdmin[admin]){
            for (uint256 i = 0; i< admins.length;i++){
                if (admins[i] == admin){
                    admins[i] = admins[admins.length-1];
                    delete admins[admins.length-1];
                    delete mapAdmin[admin];
                    admins.length--;
                    emit DelAdmin(msg.sender, admin);
                    break;
                }
            }
            if (threshold > admins.length){
                threshold = admins.length;
                if (threshold == 0) {
                    threshold = 1; // owner is always an admin.
                }
            }
        }
    }

    function setThreshold(uint256 newThreshold) public isOwner {
        require(newThreshold > 0,"threshold == 0");
        if (newThreshold > admins.length){
            threshold = admins.length;
        }else{
            threshold = newThreshold;
        }
        if (threshold == 0) {
            threshold = 1;
        }
    }

    function getValue(string calldata key) external view returns (uint256) {
        require(blockmap[key].valid || history[key].valid, "not valid");

        if (blockmap[key].valid) {
            return blockmap[key].blockNumber;
        } else {
            return history[key].blockNumber;
        }
    }

    function addProposal(string memory key,uint256 number) public isAdmin {
        require(blockmap[key].threshold == 0,"Not new Proposal");

        BlockNumber memory newproposal;
        newproposal.keywords = key;
        newproposal.threshold = threshold;
        newproposal.blockNumber = number;
        votes[key].push(msg.sender);
        if (votes[key].length >= newproposal.threshold) {
            newproposal.valid = true;
        } else {
            newproposal.valid = false;
        }

        blockmap[key] = newproposal;
    }

    function resetProposal(string memory key,uint256 number) public isAdmin {
        if (blockmap[key].valid) {
            // backup to history.
            history[key] = blockmap[key];
        }

        delete blockmap[key];
        delete votes[key];
        addProposal(key, number);
    }

    function voteProposal(string memory key) public isAdmin {
        require(blockmap[key].threshold > 0, "Not found proposal");
        for (uint256 i = 0; i < votes[key].length; i++){
            if (votes[key][i] == msg.sender){
                return;
            }
        }
        votes[key].push(msg.sender);
        if (votes[key].length >= blockmap[key].threshold){
            blockmap[key].valid = true;
        }
    }
    function getVoter(string memory key) public returns (address [] memory) {
        return votes[key];
    }
}
