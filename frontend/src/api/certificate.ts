import { GET, POST, PUT } from "api/api";
import {
  CertificateCollection,
  read as readCollection,
} from "api/certificate-collections";
import { upload } from "api/imgur";
import { MetaMask } from "components/MetaMaskProvider";
import { platform } from "const";
import * as utils from "utils";
import { CertData } from "utils";

//==============================================================================
// Type
//==============================================================================

interface Certificate extends Omit<CertData, "issuedAt" | "expiredAt"> {
  id: number;
  certNftId: number;
  collectionId: number;
  collectionAddress: string;
  issuedAt: number;
  expiredAt: number | "null";
  imgFiles?: FileList;
  revocation?: {
    revokedAt: Date;
    revokeReason: String;
  };
}

interface APICertificate {
  certData: CertData;
  id: number;
  certNftId: number;
  collectionId: number;
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

const APICertificateToCertificate = async (
  value: APICertificate
): Promise<Certificate> => ({
  id: value.id,
  certNftId: value.certNftId,
  collectionId: value.collectionId,
  collectionAddress: (await readCollection(value.collectionId))
    .collectionAddress,
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

enum VerifyState {
  Unverified,
  Verified,
  Expired,
  Invalid,
}

//==============================================================================
// Read
//==============================================================================

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
      ...(receiver !== undefined ? { receiverAddress: receiver } : {}),
      ...(collectionId !== undefined
        ? { collectionId: collectionId.toString() }
        : {}),
    })
  )
    .then((data) =>
      (data as APICertificate[]).map(async (item) =>
        APICertificateToCertificate(item)
      )
    )
    .then((promises) => Promise.all(promises));

//==============================================================================
// Issue
//==============================================================================

const issue = async (
  metaMask: MetaMask,
  collection: CertificateCollection,
  cert: Certificate
) => {
  cert.platform = platform.toString();
  cert.certImage = await upload(cert.imgFiles![0]);

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

//==============================================================================
// Verify
//==============================================================================

const verify = async (cert: Certificate) => {
  const isVerified = await utils.verifyCert(CertificateToCertData(cert), {
    certAddress: cert.collectionAddress,
    certId: cert.certNftId,
  });

  return isVerified
    ? isExpired(cert)
      ? VerifyState.Expired
      : VerifyState.Verified
    : VerifyState.Invalid;
};

const isExpired = (cert: Certificate) =>
  cert.expiredAt !== "null" && cert.expiredAt - Date.now() <= 0;

//==============================================================================
// Revoke
//==============================================================================

const revoke = async (
  metaMask: MetaMask,
  cert: Certificate,
  reason: string
) => {
  const collectionAddress = (await readCollection(cert.collectionId))
    .collectionAddress;

  const txHash = await metaMask.request(
    await utils.revokeCertTx(
      collectionAddress,
      cert.certNftId,
      reason,
      metaMask.address
    )
  );

  return PUT(`certs/${cert.id}`, { txHash: txHash, platform: "" });
};

export type { Certificate };
export { issue, read, readAll, verify, isExpired, VerifyState, revoke };
