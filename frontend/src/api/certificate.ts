import { arrayFromSize, DelayedPromise } from "helper";

interface Certificate {
  id: number;
  title: string;
  receiver: string;
  description: string;
  issuedAt: Date;
  expiredAt: Date;
  imgUrl: string;
  revokeAt?: Date;
  revokeReason?: String;
}

const today = new Date(Date.now());

const mockData: Certificate[] = arrayFromSize(30, (index) => ({
  id: Math.floor(Math.random() * 100),
  title: `Sinh viên ${index} tốt`,
  receiver: "0xb43a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7",
  description: "Đã đạt thành tích xuất sắc trong học tập",
  issuedAt: today,
  expiredAt: new Date(
    new Date(today).setDate(today.getDate() + Math.floor(Math.random() * 3))
  ),
  imgUrl: "https://picsum.photos/620/877",
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

const readAll = (receiver: string) => DelayedPromise(mockData);

export type { Certificate };
export { read, readAll, isExpired, verify, VerifyState };
