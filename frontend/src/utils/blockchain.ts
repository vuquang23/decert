import { ethers } from "ethers";
import factoryAbi from "./abi/DecertFactory.json";
import certAbi from "./abi/Decert.json";
import { createHash } from "crypto";

export interface RequestParam {
  method: string;
  params: [object];
}

export async function createNewCollectionTx(
  from: string,
  collectionName: string
): Promise<RequestParam> {
  const factoryAddress = "0x84463fe95b4bf11becff67f8894b7ce1464aaa4c";
  const iface = new ethers.utils.Interface(factoryAbi);
  const data = iface.encodeFunctionData("CreateNewDecertCollection", [
    from,
    collectionName,
    "CERT",
  ]);

  const tx = {
    nonce: "0x00",
    to: factoryAddress,
    from: from,
    value: "0x00",
    chainId: 97, // binance testnet
    data: data,
  };
  const provider = ethers.getDefaultProvider("https://data-seed-prebsc-1-s1.binance.org:8545");
  const gasLimit = await provider.estimateGas(tx);
  const gasPrice = await provider.getGasPrice();

  const transactionParameters = {
    nonce: tx.nonce,
    gasPrice: gasPrice.toHexString(),
    gas: gasLimit.toHexString(),
    to: tx.to,
    from: tx.from,
    value: tx.value,
    data: tx.data,
    chainId: "97",
  };

  return {
    method: "eth_sendTransaction",
    params: [transactionParameters],
  };
}

export interface Certificate {
  issuer: string;
  recipient: string;
  certHash: string; // should begin with 0x..
  link: string;
  issuedAt: number;
}

export async function createNewCertTx(
  collectionAddress: string,
  cert: Certificate
): Promise<RequestParam> {
  const iface = new ethers.utils.Interface(certAbi);
  const data = iface.encodeFunctionData("batchMint", [[cert]]);
  const tx = {
    nonce: "0x00",
    to: collectionAddress,
    from: cert.issuer,
    value: "0x00",
    chainId: 97, // binance testnet
    data: data,
  };
  const provider = ethers.getDefaultProvider("https://data-seed-prebsc-1-s1.binance.org:8545");
  const gasLimit = await provider.estimateGas(tx);
  const gasPrice = await provider.getGasPrice();
  const transactionParameters = {
    nonce: tx.nonce,
    gasPrice: gasPrice.toHexString(),
    gas: gasLimit.toHexString(),
    to: tx.to,
    from: tx.from,
    value: tx.value,
    data: tx.data,
    chainId: "97",
  };

  return {
    method: "eth_sendTransaction",
    params: [transactionParameters],
  };
}

export async function revokeCertTx(
  collectionAddress: string,
  tokenId: number,
  reason: string,
  from: string
): Promise<RequestParam> {
  const iface = new ethers.utils.Interface(certAbi);
  const data = iface.encodeFunctionData("revokeCertificate", [[tokenId], reason, Date.now()]);
  const tx = {
    nonce: "0x00",
    to: collectionAddress,
    from: from,
    value: "0x00",
    chainId: 97, // binance testnet
    data: data,
  };
  const provider = ethers.getDefaultProvider("https://data-seed-prebsc-1-s1.binance.org:8545");
  const gasLimit = await provider.estimateGas(tx);
  const gasPrice = await provider.getGasPrice();

  const transactionParameters = {
    nonce: tx.nonce,
    gasPrice: gasPrice.toHexString(),
    gas: gasLimit.toHexString(),
    to: tx.to,
    from: tx.from,
    value: tx.value,
    data: tx.data,
    chainId: "97",
  };

  return {
    method: "eth_sendTransaction",
    params: [transactionParameters],
  };
}

export interface CertData {
  certTitle: string;
  issuer: {
    name: string;
    wallet: string;
    position: string;
  };
  receiver: {
    name: string;
    wallet: string;
    dateOfBirth: string;
  };
  description: string;
  issuedAt: string; // string-number
  expiredAt: string; // string-number or "null"
  certImage: string;
  platform: string; // default "97"
}

export interface Proof {
  certAddress: string;
  certId: number;
}

export async function verifyCert(certData: CertData, proof: Proof): Promise<boolean> {
  const certBuffer = Buffer.from(JSON.stringify(certData));
  const hash = createHash("sha256");
  hash.update(certBuffer);
  const certHash = hash.digest("hex");
  const provider = ethers.getDefaultProvider("https://data-seed-prebsc-1-s1.binance.org:8545");
  const contract = new ethers.Contract(proof.certAddress, certAbi, provider);
  const response = await contract.certData(proof.certId);
  return `0x${certHash}` === response["certHash"];
}
