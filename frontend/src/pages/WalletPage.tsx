import { Certificate, isExpired, readAll } from "api/certificate";
import Empty from "components/Empty";
import HeaderSearch, { Inputs, searchByTitle } from "components/HeaderSearch";
import { useMetaMask } from "components/MetaMaskProvider";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { arrayFromSize } from "helper";
import { onPromiseRejected } from "pages/ErrorPage";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const WalletPage = () => {
  const navigate = useNavigate();
  const [certs, setCerts] = useState<Certificate[]>();
  const [searchInputs, setSearchInputs] = useState<Inputs>({
    searchQuery: "",
    filter: "Valid",
  });
  const { address } = useMetaMask();
  useEffect(() => {
    // TODO: Change to receiver
    readAll({ collectionId: 2 })
      .then((value) =>
        setCerts(value.filter((cert) => cert.revocation === undefined))
      )
      .catch((reason) => onPromiseRejected(reason, navigate));
  }, [address, navigate]);

  return (
    <>
      <HeaderSearch
        title="Your Certificates"
        placeholder="Certificate name..."
        filters={["All", "Valid", "Expired"]}
        defaultFilter="Valid"
        onSearchSubmit={(inputs) => setSearchInputs(inputs)}
      />
      <div className="row gy-5 mb-5">
        {certs !== undefined ? (
          certs.length === 0 ? (
            <Empty />
          ) : (
            searchByTitle(
              certs.map((cert) => ({ ...cert, title: cert.certTitle })),
              searchInputs.searchQuery
            )
              .filter((cert) => {
                switch (searchInputs.filter) {
                  case "Valid":
                    return !isExpired(cert);
                  case "Expired":
                    return isExpired(cert);
                  default:
                    return true;
                }
              })
              .map((cert, index) => <CertificateCard key={index} cert={cert} />)
          )
        ) : (
          arrayFromSize(8, (index) => (
            <CertificateCardPlaceholder key={index} />
          ))
        )}
      </div>
    </>
  );
};

const CertificateCard = ({ cert }: { cert: Certificate }) => {
  const navigate = useNavigate();
  return (
    <div className="col-12 col-sm-6 col-md-4 col-lg-3">
      <div
        className={`card hover-shadow clickable ${
          isExpired(cert) ? "text-bg-secondary" : ""
        }`}
        onClick={() => navigate(`/certificate/${cert.id}`)}
      >
        <img
          src={cert.certImage}
          className="card-img-top crop"
          alt="Certificate"
        />
        <div className="card-body">
          <h5 className="card-title">{cert.certTitle}</h5>
          <p className="card-text">
            <strong>Issued at:</strong> {new Date(cert.issuedAt).toDateString()}
            {cert.expiredAt !== "null" && (
              <>
                <br />
                <strong>Expired at:</strong>{" "}
                {new Date(cert.expiredAt).toDateString()}
              </>
            )}
          </p>
          <p className="card-text line-clamp">{cert.description}</p>
        </div>
      </div>
    </div>
  );
};

const CertificateCardPlaceholder = () => (
  <div className="col-12 col-sm-6 col-md-4 col-lg-3">
    <div className="card placeholder-glow">
      <div className="card-img-top placeholder" />
      <div className="card-body">
        <div className="h5 card-title placeholder col-4" />
        <div className="card-text">
          <ParagraphPlaceholder />
        </div>
      </div>
    </div>
  </div>
);

export default WalletPage;
