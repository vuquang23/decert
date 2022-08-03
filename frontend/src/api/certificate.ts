import { MetaMask } from "components/MetaMaskProvider";
import { arrayFromSize, DelayedPromise } from "helper";

interface Issuer {
  name: string;
  address: string;
  position: string;
}

interface Receiver {
  name: string;
  address: string;
  dateOfBirth: Date;
}

interface Revocation {
  revokedAt: Date;
  revokeReason: String;
}

interface Certificate {
  id: number;
  title: string;
  issuer: Issuer;
  receiver: Receiver;
  description: string;
  issuedAt: Date;
  expiredAt: Date;
  imgUrl: string;
  revocation?: Revocation;
}

const today = new Date(Date.now());

const mockData: Certificate[] = arrayFromSize(30, (index) => ({
  id: Math.floor(Math.random() * 100),
  title: `Sinh viên ${index} tốt`,
  issuer: {
    name: "Nguyen Van A",
    address: "0xb43a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7",
    position: "Principal",
  },
  receiver: {
    name: "Nguyen Van A",
    address: "0xb43a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7",
    dateOfBirth: today,
  },
  description: "Đã đạt thành tích xuất sắc trong học tập",
  issuedAt: today,
  expiredAt: new Date(
    new Date(today).setDate(today.getDate() + Math.floor(Math.random() * 3))
  ),
  imgUrl: "https://picsum.photos/620/877",
  revocation:
    Math.random() > 0.7
      ? { revokedAt: today, revokeReason: "Issued by mistake" }
      : undefined,
}));

const isExpired = (cert: Certificate) =>
  cert.expiredAt.getTime() - Date.now() <= 0;

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
      return typeof result !== "undefined" ? resolve(result) : reject();
    })
  );

const readAll = ({
  receiver,
  collectionId,
}: {
  receiver?: string;
  collectionId?: number;
}) => DelayedPromise(mockData);

const revoke = async (metaMask: MetaMask, id: number) => {
  const cert = mockData.find((cert) => cert.id === id);
  if (typeof cert !== "undefined") {
    cert.revocation = { revokedAt: today, revokeReason: "Issued by mistake" };
  }
  return read(id);
};

export type { Certificate };
export { read, readAll, revoke, isExpired, verify, VerifyState };
