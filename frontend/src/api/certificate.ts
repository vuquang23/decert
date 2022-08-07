import { MetaMask } from "components/MetaMaskProvider";
import { arrayFromSize, DelayedPromise, toDDMMYYYYstring } from "helper";
import * as Utils from "utils";

interface Certificate {
  id: number;
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
  issuedAt: number;
  expiredAt: number | "null";
  certImage: string;
  platform: string;
  imgFiles?: FileList;
  revocation?: {
    revokedAt: Date;
    revokeReason: String;
  };
}

const today = new Date(Date.now());

const mockData: Certificate[] = arrayFromSize(30, (index) => ({
  id: Math.floor(Math.random() * 100),
  certTitle: `Sinh viên ${index} tốt`,
  issuer: {
    name: "Nguyen Van A",
    wallet: "0xb43a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7",
    position: "Principal",
  },
  receiver: {
    name: "Nguyen Van A",
    wallet: "0xb43a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7",
    dateOfBirth: toDDMMYYYYstring(today),
  },
  description: "Đã đạt thành tích xuất sắc trong học tập",
  issuedAt: today.getTime(),
  expiredAt: new Date(
    new Date(today).setDate(today.getDate() + Math.floor(Math.random() * 3))
  ).getTime(),
  certImage: "https://picsum.photos/620/877",
  platform: "97",
  revocation:
    Math.random() > 0.7
      ? { revokedAt: today, revokeReason: "Issued by mistake" }
      : undefined,
}));

const isExpired = (cert: Certificate) =>
  cert.expiredAt !== "null" && cert.expiredAt - Date.now() <= 0;

enum VerifyState {
  Unverified,
  Verified,
  Expired,
  Invalid,
}

const verify = (cert: Certificate) =>
  DelayedPromise(
    cert.id % 2 === 0
      ? isExpired(cert)
        ? VerifyState.Expired
        : VerifyState.Verified
      : VerifyState.Invalid
  );

const read = (id: number) =>
  DelayedPromise(
    new Promise<Certificate>((resolve, reject) => {
      const result = mockData.find((cert) => cert.id === id);
      return result !== undefined ? resolve(result) : reject();
    })
  );

const readAll = ({
  receiver,
  collectionId,
}: {
  receiver?: string;
  collectionId?: number;
}) => DelayedPromise(mockData);

const issue = async (metaMask: MetaMask, cert: Certificate) => {
  await Utils.pushToIpfs(cert.imgFiles![0]);
  cert.id = Math.floor(Math.random() * 100 + 100);
  cert.certImage = URL.createObjectURL(cert.imgFiles![0]);
  mockData.push(cert);
  await DelayedPromise(0);
  return cert.id;
};

const revoke = async (metaMask: MetaMask, id: number) => {
  const cert = mockData.find((cert) => cert.id === id);
  if (cert !== undefined) {
    cert.revocation = { revokedAt: today, revokeReason: "Issued by mistake" };
  }
  await DelayedPromise(0);
  return read(id);
};

export type { Certificate };
export { read, readAll, revoke, issue, isExpired, verify, VerifyState };
