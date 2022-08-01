import { MetaMask } from "components/MetaMaskProvider";

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

const mockData: CertificateCollection[] = Array.from(
  Array(30).keys(),
  (_, index) => newCertCollection(index)
);

const create = async (metaMask: MetaMask, name: string) => {
  mockData.push(newCertCollection(name));
  return mockData.at(-1);
};

const readAll = (issuer: string) => Promise.resolve(mockData);

export type { CertificateCollection };
export { create, readAll };
