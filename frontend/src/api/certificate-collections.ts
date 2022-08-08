import { GET, POST } from "api/api";
import { MetaMask } from "components/MetaMaskProvider";
import * as utils from "utils";

interface CertificateCollection {
  id: number;
  collectionName: string;
  collectionAddress: string;
  issuer: string;
  totalIssued: number;
  totalRevoked: number;
}

const create = (metaMask: MetaMask, title: string) =>
  utils
    .createNewCollectionTx(metaMask.address, title)
    .then((requestParam) => metaMask.request(requestParam))
    .then((txHash: string) =>
      POST("collections", { txHash: txHash, platform: "" })
    );

const readAll = (
  issuer: string,
  limit: number = 0,
  offset: number = 0,
  name?: string
) =>
  GET(
    "collections",
    new URLSearchParams({
      issuer: issuer,
      limit: limit.toString(),
      offset: offset.toString(),
      ...(name !== undefined && name.length > 0 ? { name: name } : {}),
    })
  ).then((data) => data as CertificateCollection[]);

const read = (issuer: string, id: number) =>
  readAll(issuer).then((data) =>
    data.find((collection) => collection.id === id)
  );

export type { CertificateCollection };
export { create, readAll, read };
