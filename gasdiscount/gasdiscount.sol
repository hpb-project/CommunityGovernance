pragma solidity ^0.5.1;
pragma experimental ABIEncoderV2;
contract gasdiscount {
    address  superadmin;
    address[]  admins;
    mapping(address => bool) mapadmin;
    uint256 threshold;
    struct Proposal {
	    address dapp;
    	uint256 rate;
     	uint256 max_gas;
    	bool voted;
    }

    address[] proposallist; //所有的proposal
    mapping(address => Proposal) proposals; //key为dapp地址
    mapping(address => address[]) votes;//key为dapp ，value为投票地址
    
    function deleteproposal(address dappaddr) public {
        require(superadmin == msg.sender,"Not superadmin");//只有superadmin能删除
        require(proposals[dappaddr].voted == false,"Proposal Voted");
        delete proposals[dappaddr];
        delete votes[dappaddr];
        for (uint256 i = 0; i < proposallist.length; i++){
            if (dappaddr == proposallist[i]){
                proposallist[i] = proposallist[proposallist.length-1];
                delete proposallist[proposallist.length-1];
            }
        }
    }
    
    constructor() public{
        superadmin = msg.sender;
        threshold = 1;
    }
    
    function setThreshold(uint256 acount) public{
        require(superadmin == msg.sender,"Not superadmin");
        if (acount > admins.length){
            threshold = admins.length;
        }else if (acount < 1){
            threshold = 1;
        }else{
            threshold = acount;
        }
    }
    
    function changeSuperAdmin(address spadmin) public{
        require(superadmin == msg.sender,"Not superadmin");
        superadmin = spadmin;
    }
    
    function addAdmin(address admin)public{
        require(superadmin == msg.sender,"Not superamdin");
        if (!mapadmin[admin]){
            mapadmin[admin] = true;
            admins.push(admin);
        }
    }
    function getAdmins() public view returns (address[] memory){
        return admins;
    }
    
    function deleteAdmin(address admin) public{
        require(superadmin == msg.sender,"Not superamdin");
        if (mapadmin[admin]){
            for (uint256 i = 0; i< admins.length;i++){
                if (admins[i] == admin){
                    admins[i] = admins[admins.length-1];
                    delete admins[admins.length-1];
                    delete mapadmin[admin];
                }
            }
        }
        if (threshold > admins.length){
            threshold = admins.length;
        }
    }
    
    function addProposal(address dappaddr,uint256 rate,uint256 max_gas) public{
        require(mapadmin[msg.sender],"Not admin");
        require(proposals[dappaddr].max_gas == 0,"Not new Proposal");
        require(max_gas > 0 && rate > 0 && rate < 100,"check param");
        
        Proposal memory newproposal;
        newproposal.dapp = dappaddr;
        newproposal.rate = rate;
        newproposal.max_gas = max_gas;
        newproposal.voted = false;
        if (threshold == 1){
            newproposal.voted = true;
        }
        proposals[dappaddr] = newproposal;
        votes[dappaddr].push(msg.sender);
        proposallist.push(dappaddr);
    }
    
    function voteProposal(address dappaddr) public{
        require(mapadmin[msg.sender],"Not admin");
        require(proposals[dappaddr].max_gas > 0,"Not Proposal");
        for (uint256 i = 0; i < votes[dappaddr].length; i++){
            if (votes[dappaddr][i] == msg.sender){
                return;
            }
        }
        votes[dappaddr].push(msg.sender);
        if (votes[dappaddr].length >= threshold){
            proposals[dappaddr].voted = true;
        }
    }
    
    function listproposal(uint256 args) public view returns(Proposal[] memory){
        uint256 amount = proposallist.length;
        if (args == 0){
            args = amount;
        } else if (args > amount){
            args = amount;
        }
        Proposal[] memory ps = new Proposal[](args);
        for (uint256 i = 0; i < args; i++){
            ps[i] = proposals[proposallist[proposallist.length-1-i]];
        }
        return ps;
    }
    
    function getratebyaddr(address dappaddr) public view returns(bool,uint256,uint256){
        Proposal memory proposal = proposals[dappaddr];
        return (proposal.voted,proposal.rate,proposal.max_gas);
    }

}
