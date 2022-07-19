import React, { ReactElement, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import Web3 from "web3";

interface MetaMask {
  account: string;
  isIssuer: boolean;
  connectToMetaMask: () => void;
}

const web3 = new Web3(Web3.givenProvider);

const MetaMaskContext = React.createContext<MetaMask>({
  account: "",
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
        navigate("/collections/1");
      }
    } else {
      setAccount("");
    }
  };

  web3.eth.getAccounts((_error, account) => updateAccount(account[0]));

  const context: MetaMask = {
    account: account,
    isIssuer: isIssuer,
    connectToMetaMask: () =>
      web3.eth.requestAccounts((_error, account) =>
        updateAccount(account[0], true)
      ),
  };

  return <MetaMaskContext.Provider value={context} children={children} />;
};

export const useMetaMask = () => useContext(MetaMaskContext);

export default MetaMaskProvider;
