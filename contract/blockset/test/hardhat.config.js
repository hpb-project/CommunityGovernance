require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
//task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
//  const accounts = await hre.ethers.getSigners();
//
//  for (const account of accounts) {
//    console.log(account.address);
//  }
//});

// Define mnemonic for accounts.
let mnemonic = process.env.MNEMONIC;
if (!mnemonic) {
	// NOTE: this fallback is for development only!
	// When using other networks, set the secret in .env.
	// DO NOT commit or share your mnemonic with others!
	//mnemonic = 'test test test test test test test test test test test test';
	mnemonic = 'dove vintage will kangaroo shaft enrich craft rail exact mask arrange express';
}

const accounts = { mnemonic };

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
	solidity: "0.5.1",
	networks: {
		hardhat: {
			accounts,
		},
		testnet: {
			url: 'http://114.242.26.15:8006',
			accounts,
		},
		mainnet: {
			url: 'https://hpbnode.com',
			accounts,
		}
	},
};
