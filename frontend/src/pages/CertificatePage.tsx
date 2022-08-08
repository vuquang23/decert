import { Certificate, read, verify, VerifyState } from "api/certificate";
import Address from "components/Address";
import Center from "components/Center";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { dateFromDDMMYYYY } from "helper";
import { onPromiseRejected } from "pages/ErrorPage";
import { NotFoundError } from "pages/NotFoundPage";
import { ReactNode, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

const CertificatePage = () => {
  const { certId } = useParams();
  const navigate = useNavigate();
  const [cert, setCert] = useState<Certificate>();
  const [verifyState, setVerifyState] = useState(VerifyState.Unverified);
  useEffect(() => {
    read(parseInt(certId!))
      .then((cert) => {
        if (cert === undefined || cert.revocation !== undefined) {
          throw new NotFoundError();
        }
        setCert(cert);
        return verify(cert);
      })
      .then((verifyState) => setVerifyState(verifyState))
      .catch((reason) => {
        if (reason instanceof NotFoundError) {
          navigate("/notfound");
        } else {
          onPromiseRejected(reason, navigate);
        }
      });
  }, [certId, navigate]);

  return (
    <>
      <Offcanvas>
        <CertificateContent cert={cert} verifyState={verifyState} />
      </Offcanvas>
      <div className="container-fluid vh-100">
        <div className="row h-100">
          <div className="col-4 shadow p-4 d-none d-lg-block">
            <CertificateContent cert={cert} verifyState={verifyState} />
          </div>
          <Center className="col-12 col-lg-8 align-self-center h-85">
            {cert !== undefined ? (
              <img
                className="rounded shadow mh-100 mw-100"
                src={cert.certImage}
                alt="Certificate"
              />
            ) : (
              <div className="bg-secondary rounded shadow h-100 w-50" />
            )}
          </Center>
        </div>
      </div>
    </>
  );
};

const Offcanvas = ({ children }: { children: ReactNode }) => {
  const navigate = useNavigate();
  return (
    <>
      <div className="position-fixed top-0 start-0 ms-2 mt-2 btn-group shadow-sm d-lg-none">
        <button className="btn btn-light border" onClick={() => navigate(-1)}>
          <i className="bi bi-chevron-left" />
        </button>
        <button
          className="btn btn-light border"
          data-bs-toggle="offcanvas"
          data-bs-target="#offcanvas"
        >
          <i className="bi bi-list" />
        </button>
      </div>
      <div
        className="offcanvas w-100 offcanvas-start"
        tabIndex={-1}
        id="offcanvas"
      >
        <div className="offcanvas-header">
          <button
            type="button"
            className="btn-close ms-auto"
            data-bs-dismiss="offcanvas"
            data-bs-target="#offcanvasResponsive"
          />
        </div>
        <div className="offcanvas-body">{children}</div>
      </div>
    </>
  );
};

const CertificateContent = ({
  cert,
  verifyState,
}: {
  cert?: Certificate;
  verifyState: VerifyState;
}) => {
  const navigate = useNavigate();
  return (
    <>
      <button
        className="btn btn-lg border rounded-pill shadow-sm mb-5 d-none d-lg-block"
        onClick={() => navigate(-1)}
      >
        <i className="bi bi-chevron-left" /> Back
      </button>
      {cert !== undefined ? (
        <>
          <State verifyState={verifyState} />
          <h2>{cert.certTitle}</h2>
          <p className="lead">
            <strong>Issued to: {cert.receiver.name.toLocaleUpperCase()}</strong>
            <br />
            <small>
              <strong>Date of birth:</strong>{" "}
              {dateFromDDMMYYYY(cert.receiver.dateOfBirth).toDateString()}
              <br />
              <strong>Address:</strong>{" "}
              <Address address={cert.receiver.wallet} />
            </small>
            <br />
            <br />
            {cert.description}
          </p>
          <hr />
          <p>
            <strong>Issued by:</strong> {cert.issuer.name.toLocaleUpperCase()}
            <br />
            <strong>Position:</strong> {cert.issuer.position}
            <br />
            <strong>Address:</strong> <Address address={cert.issuer.wallet} />
            <br />
            <br />
            <strong>Issued at:</strong> {new Date(cert.issuedAt).toDateString()}
            {cert.expiredAt !== "null" && (
              <>
                <br />
                <strong>Expired at:</strong>{" "}
                {new Date(cert.expiredAt).toDateString()}
              </>
            )}
          </p>
        </>
      ) : (
        <CertificateContentPlaceholder />
      )}
    </>
  );
};

const State = ({ verifyState }: { verifyState: VerifyState }) => {
  switch (verifyState) {
    case VerifyState.Unverified:
      return <CalloutPlaceholder />;
    case VerifyState.Verified:
      return <Callout color="success" text="Verified" />;
    case VerifyState.Expired:
      return <Callout color="danger" text="Expired" />;
    case VerifyState.Invalid:
      return <Callout color="danger" text="Invalid" />;
    default:
      return <CalloutPlaceholder />;
  }
};

const Callout = ({
  color,
  text,
}: {
  color: "success" | "danger";
  text: string;
}) => {
  const icon =
    color === "success" ? "shield-fill-check" : "exclamation-diamond-fill";
  return (
    <div className={`bd-callout bd-callout-${color} rounded h5`}>
      <i className={`bi bi-${icon} text-${color}`} /> {text}
    </div>
  );
};

const CalloutPlaceholder = () => (
  <div className="bd-callout rounded">
    <span className="placeholder col-4" />
  </div>
);

const CertificateContentPlaceholder = () => (
  <div className="placeholder-glow">
    <CalloutPlaceholder />
    <div className="h2 placeholder col-4" />
    <ParagraphPlaceholder />
  </div>
);

export default CertificatePage;
