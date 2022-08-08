import { GET, POST } from "api/api";
import { CertificateCollection } from "api/certificate-collections";
import { MetaMask } from "components/MetaMaskProvider";
import { platform } from "const";
import { arrayFromSize, DelayedPromise, toDDMMYYYYstring } from "helper";
import * as utils from "utils";
import { CertData } from "utils";

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

interface APICertificate {
  certData: CertData;
  id: number;
  expiredAt: string;
  revokedAt: string;
  revokedReason: string;
}

const CertificateToCertData = (cert: Certificate): utils.CertData => ({
  certTitle: cert.certTitle,
  issuer: cert.issuer,
  receiver: cert.receiver,
  description: cert.description,
  issuedAt: cert.issuedAt.toString(),
  expiredAt: cert.expiredAt.toString(),
  certImage: cert.certImage,
  platform: cert.platform,
});

const APICertificateToCertificate = (value: APICertificate): Certificate => ({
  id: value.id,
  ...value.certData,
  issuedAt: parseInt(value.certData.issuedAt),
  expiredAt: value.expiredAt !== "null" ? parseInt(value.expiredAt) : "null",
  ...(value.revokedAt !== "null"
    ? {
        revocation: {
          revokedAt: new Date(parseInt(value.revokedAt)),
          revokeReason: value.revokedReason,
        },
      }
    : {}),
});

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
  GET(`certs/${id}`).then((data) =>
    APICertificateToCertificate(data as APICertificate)
  );

const readAll = ({
  receiver,
  collectionId,
  limit,
  offset,
}: {
  receiver?: string;
  collectionId?: number;
  limit?: number;
  offset?: number;
}) =>
  GET(
    "certs",
    new URLSearchParams({
      limit: (limit ?? 0).toString(),
      offset: (offset ?? 0).toString(),
      ...(collectionId !== undefined
        ? { collectionId: collectionId.toString() }
        : {}),
    })
  ).then((data) =>
    (data as APICertificate[]).map((item) => APICertificateToCertificate(item))
  );

const issue = async (
  metaMask: MetaMask,
  collection: CertificateCollection,
  cert: Certificate
) => {
  cert.platform = platform.toString();
  // TODO: uncomment
  /* const imgUrl = await utils.pushToIpfs(cert.imgFiles![0]);
  if (imgUrl instanceof Error) {
    throw imgUrl;
  }
  cert.certImage = imgUrl; */
  cert.certImage = "https://picsum.photos/620/877";

  const certData = CertificateToCertData(cert);
  const certHash = utils.hashCert(certData);

  const txHash = await metaMask.request(
    await utils.createNewCertTx(collection.collectionAddress, {
      issuer: metaMask.address,
      recipient: cert.receiver.wallet,
      certHash: certHash,
      link: cert.certImage,
      issuedAt: cert.issuedAt,
    })
  );
  return POST("certs", {
    collectionId: collection.id,
    certData: certData,
    txHash: txHash,
    platform: "",
  });
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
