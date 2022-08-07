import { BigNumber, ethers } from "ethers";
import React, { ReactNode, useContext, useEffect, useState } from "react";
import { RequestParam } from "utils";

interface MetaMask {
  isReady: boolean;
  address: string;
  request: (requestParam: RequestParam) => Promise<any>;
  getBalance: () => Promise<string>;
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
  isReady: false,
  address: "",
  request: () => Promise.reject(),
  getBalance: () => Promise.reject(),
  connectToMetaMask: () => {},
});

const MetaMaskProvider = ({ children }: { children: ReactNode }) => {
  const [address, setAddress] = useState("");
  const [isReady, setIsReady] = useState(false);

  const updateAddress = (address: string) => {
    address !== undefined ? setAddress(address) : setAddress("");
    setIsReady(true);
  };

  useEffect(() => {
    provider.listAccounts().then((addresses) => updateAddress(addresses[0]));
  }, []);

  const context: MetaMask = {
    isReady: isReady,
    address: address,
    request: ({ method, params }) => provider.send(method, params),
    getBalance: async function () {
      return weiToEther(await provider.getBalance(this.address));
    },
    connectToMetaMask: () =>
      provider
        .send("eth_requestAccounts", [])
        .then((addresses) => updateAddress(addresses[0])),
  };

  return <MetaMaskContext.Provider value={context} children={children} />;
};

const weiToEther = (wei: BigNumber) => {
  const divResult = wei.div(BigNumber.from("10000000000000000")).toString();
  if (divResult.length === 2) {
    return `0.${divResult}`;
  }
  return `${divResult.substring(0, divResult.length - 2)}.${divResult.substring(
    divResult.length - 2
  )}`;
};

const useMetaMask = () => useContext(MetaMaskContext);

export default MetaMaskProvider;
export type { MetaMask };
export { getShortAddress, useMetaMask };
