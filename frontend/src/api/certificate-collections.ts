import { MetaMask } from "components/MetaMaskProvider";
import { arrayFromSize, DelayedPromise } from "helper";

interface CertificateCollection {
  id: number;
  title: string;
  address: string;
  issuer: string;
  issued: number;
  revoked: number;
}

const newCertCollection = (n: number | string) =>
  typeof n === "number"
    ? {
        id: Math.floor(Math.random() * 100),
        title: `Sinh viên ${n} tốt`,
        address: `0xb${n}a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7`,
        issuer: "",
        issued: 30,
        revoked: 20,
      }
    : {
        id: Math.floor(Math.random() * 100),
        title: n,
        address: `0xb${
          Math.floor(Math.random() * 100) + 31
        }a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7`,
        issuer: "",
        issued: 30,
        revoked: 20,
      };

const mockData: CertificateCollection[] = arrayFromSize(30, (index) =>
  newCertCollection(index)
);

const create = async (metaMask: MetaMask, name: string) => {
  mockData.push(newCertCollection(name));
  await DelayedPromise(0);
  return mockData.at(-1);
};

const readAll = (issuer: string) => DelayedPromise(mockData);

const read = (id: number) =>
  DelayedPromise(
    new Promise<CertificateCollection>((resolve, reject) => {
      const result = mockData.find((collection) => collection.id === id);
      return typeof result !== "undefined" ? resolve(result) : reject();
    })
  );

export type { CertificateCollection };
export { create, readAll, read };
