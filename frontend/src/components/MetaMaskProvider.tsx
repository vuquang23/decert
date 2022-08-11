import { platform } from "const";
import { BigNumber, ethers } from "ethers";
import React, { ReactNode, useContext, useEffect, useState } from "react";
import { RequestParam } from "utils";

interface MetaMask {
  isReady: boolean;
  address: string;
  request: (requestParam: RequestParam) => Promise<any>;
  getBalance: () => Promise<string>;
  connectToMetaMask: () => Promise<void>;
}

const getShortAddress = (address: string) =>
  address.substring(0, 5) + "..." + address.substring(address.length - 4);

class MetaMaskNotFound extends Error {}

const MetaMaskContext = React.createContext<MetaMask>({
  isReady: false,
  address: "",
  request: () => Promise.reject(),
  getBalance: () => Promise.reject(),
  connectToMetaMask: () => Promise.reject(),
});

const MetaMaskProvider = ({ children }: { children: ReactNode }) => {
  const [provider, setProvider] = useState<ethers.providers.Web3Provider>();
  const [address, setAddress] = useState("");
  const [isReady, setIsReady] = useState(false);

  const updateAddress = (address: string) => {
    address !== undefined ? setAddress(address) : setAddress("");
    setIsReady(true);
  };

  useEffect(() => {
    if ((window as any).ethereum) {
      const provider = new ethers.providers.Web3Provider(
        (window as any).ethereum,
        platform
      );
      setProvider(provider);
      provider.listAccounts().then((addresses) => updateAddress(addresses[0]));
    } else {
      setIsReady(true);
    }
  }, []);

  const doWithProvider = async <T,>(
    action: (provider: ethers.providers.Web3Provider) => Promise<T>
  ) => {
    if (provider !== undefined) {
      return await action(provider);
    } else {
      throw new MetaMaskNotFound();
    }
  };

  const context: MetaMask = {
    isReady: isReady,
    address: address,
    request: ({ method, params }) =>
      doWithProvider((provider) => provider.send(method, params)),
    getBalance: async function () {
      return weiToEther(
        await doWithProvider((provider) => provider.getBalance(this.address))
      );
    },
    connectToMetaMask: () =>
      doWithProvider((provider) =>
        provider
          .send("eth_requestAccounts", [])
          .then((addresses) => updateAddress(addresses[0]))
      ),
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
export { getShortAddress, useMetaMask, MetaMaskNotFound };
