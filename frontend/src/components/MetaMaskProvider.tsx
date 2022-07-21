import { ethers } from "ethers";
import React, { ReactElement, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

interface MetaMask {
  account: string;
  accountShort: string;
  getBalance: () => Promise<number>;
  connectToMetaMask: () => void;
}

const getShortAccount = (account: string) =>
  account.substring(0, 5) + "..." + account.substring(account.length - 4);

// TODO: What if metamask is not installed
const provider = new ethers.providers.Web3Provider(
  (window as any).ethereum,
  97
);

const MetaMaskContext = React.createContext<MetaMask>({
  account: "",
  accountShort: "",
  getBalance: () => Promise.resolve(0),
  connectToMetaMask: () => {},
});

const MetaMaskProvider = ({ children }: { children: ReactElement }) => {
  const [account, setAccount] = useState("");
  const navigate = useNavigate();

  const updateAccount = (
    account: string,
    navigateToCollections: boolean = false
  ) => {
    if (account !== undefined) {
      setAccount(account);
      if (navigateToCollections) {
        navigate("/collections");
      }
    } else {
      setAccount("");
    }
  };

  provider
    .send("eth_accounts", [])
    .then((accounts) => updateAccount(accounts[0]));

  const context: MetaMask = {
    account: account,
    get accountShort() {
      return getShortAccount(this.account);
    },
    getBalance: async function () {
      return (await provider.getBalance(this.account)).toNumber();
    },
    connectToMetaMask: () =>
      provider
        .send("eth_requestAccounts", [])
        .then((accounts) => updateAccount(accounts[0], true)),
  };

  return <MetaMaskContext.Provider value={context} children={children} />;
};

const useMetaMask = () => useContext(MetaMaskContext);

export default MetaMaskProvider;
export { getShortAccount, useMetaMask };
