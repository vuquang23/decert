require('dotenv').config();
require("@nomiclabs/hardhat-waffle");
require('@nomiclabs/hardhat-etherscan');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.9",
  networks: {
    bscTest: {
      url: process.env.BSC_TEST_RPC,
      accounts: [
        process.env.BSC_TEST_ACCOUNT
      ]
    }
  },
  etherscan: {
    apiKey: process.env.BSCSCAN_API_KEY
  },
};
