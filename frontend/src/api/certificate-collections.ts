interface CertificateCollection {
  certificateName: string;
  address: string;
  issued: number;
  revoked: number;
  createdDate: Date;
}

const newCertCollection = (n: number | string) =>
  typeof n === "number"
    ? {
        certificateName: `Sinh viên ${n} tốt`,
        address: `0xb${n}a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7`,
        issued: 30,
        revoked: 20,
        createdDate: new Date(Date.now()),
      }
    : {
        certificateName: n,
        address: `0xb${
          Math.floor(Math.random() * 100) + 31
        }a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7`,
        issued: 30,
        revoked: 20,
        createdDate: new Date(Date.now()),
      };

const mockData: CertificateCollection[] = Array.from(
  Array(30).keys(),
  (_, index) => newCertCollection(index)
);

const create = async (name: string) => {
  mockData.push(newCertCollection(name));
  console.log(mockData[-1]);
  return mockData.at(-1);
};

const readAll = () => Promise.resolve(mockData);

export type { CertificateCollection };
export { create, readAll };
