import React, { ReactElement, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import Web3 from "web3";

interface MetaMask {
  account: string;
  accountShort: string;
  isIssuer: boolean;
  connectToMetaMask: () => void;
}

const getShortAccount = (account: string) =>
  account.substring(0, 5) + "..." + account.substring(account.length - 4);

const web3 = new Web3(Web3.givenProvider);

const MetaMaskContext = React.createContext<MetaMask>({
  account: "",
  get accountShort() {
    return getShortAccount(this.account);
  },
  isIssuer: false,
  connectToMetaMask: function () {
    web3.eth.requestAccounts((_error, account) => (this.account = account[0]));
  },
});

const MetaMaskProvider = ({ children }: { children: ReactElement }) => {
  const [account, setAccount] = useState("");
  const [isIssuer, setIsIssuer] = useState(false);
  const navigate = useNavigate();

  const updateAccount = (
    account: string,
    navigateToCollections: boolean = false
  ) => {
    if (account !== undefined) {
      setAccount(account);
      setIsIssuer(true);
      if (navigateToCollections) {
        navigate("/collections");
      }
    } else {
      setAccount("");
    }
  };

  web3.eth.getAccounts((_error, account) => updateAccount(account[0]));

  const context: MetaMask = {
    account: account,
    get accountShort() {
      return getShortAccount(this.account);
    },
    isIssuer: isIssuer,
    connectToMetaMask: () =>
      web3.eth.requestAccounts((_error, account) =>
        updateAccount(account[0], true)
      ),
  };

  return <MetaMaskContext.Provider value={context} children={children} />;
};

const useMetaMask = () => useContext(MetaMaskContext);

export default MetaMaskProvider;
export { getShortAccount, useMetaMask };
