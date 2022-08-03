import { ethers } from "ethers";
import React, { ReactNode, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

interface MetaMask {
  address: string;
  getBalance: () => Promise<number>;
  connectToMetaMask: () => void;
}

const getShortAddress = (address: string) =>
  address.substring(0, 5) + "..." + address.substring(address.length - 4);

// TODO: What if metamask is not installed
const provider = new ethers.providers.Web3Provider(
  (window as any).ethereum,
  97
);

const MetaMaskContext = React.createContext<MetaMask>({
  address: "",
  getBalance: () => Promise.resolve(0),
  connectToMetaMask: () => {},
});

const MetaMaskProvider = ({ children }: { children: ReactNode }) => {
  const [address, setAddress] = useState("");
  const navigate = useNavigate();

  const updateAddress = (
    address: string,
    navigateToCollections: boolean = false
  ) => {
    if (address !== undefined) {
      setAddress(address);
      if (navigateToCollections) {
        navigate("/collections");
      }
    } else {
      setAddress("");
    }
  };

  provider
    .send("eth_accounts", [])
    .then((addresses) => updateAddress(addresses[0]));

  const context: MetaMask = {
    address: address,
    getBalance: async function () {
      return (await provider.getBalance(this.address)).toNumber();
    },
    connectToMetaMask: () =>
      provider
        .send("eth_requestAccounts", [])
        .then((addresses) => updateAddress(addresses[0], true)),
  };

  return <MetaMaskContext.Provider value={context} children={children} />;
};

const useMetaMask = () => useContext(MetaMaskContext);

export default MetaMaskProvider;
export type { MetaMask };
export { getShortAddress, useMetaMask };
