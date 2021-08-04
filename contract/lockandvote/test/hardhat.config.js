require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    const balance = await network.provider.send("eth_getBalance", [account.address])
    console.log(account.address, balance);
  }
});

// Define mnemonic for accounts.
let mnemonic = process.env.MNEMONIC;
if (!mnemonic) {
        // NOTE: this fallback is for development only!
        // When using other networks, set the secret in .env.
        // DO NOT commit or share your mnemonic with others!
        mnemonic = 'test test test test test test test test test test test test';        
}
const accounts = { mnemonic };

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  networks: {
    hardhat: {
      accounts:{
        accountsBalance:"800000000000000000000000"
      }
    }
  },
  solidity: "0.5.1",
};
