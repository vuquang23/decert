import { Certificate, isExpired, readAll } from "api/certificate";
import HeaderSearch, { Inputs, searchByTitle } from "components/HeaderSearch";
import { useMetaMask } from "components/MetaMaskProvider";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const WalletPage = () => {
  const [certs, setCerts] = useState<Certificate[]>([]);
  const [searchInputs, setSearchInputs] = useState<Inputs>({
    searchQuery: "",
    filter: "Valid",
  });
  const { account } = useMetaMask();
  useEffect(() => {
    readAll(account).then((value) => setCerts(value));
  });

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
        {searchByTitle(certs, searchInputs.searchQuery)
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
          .map((cert, index) => (
            <CertificateCard key={index} cert={cert} />
          ))}
      </div>
    </>
  );
};

const CertificateCard = ({ cert }: { cert: Certificate }) => {
  const navigate = useNavigate();
  return (
    <div className="col-12 col-sm-6 col-md-4 col-lg-3">
      <div
        className={`card hover-shadow ${
          isExpired(cert) ? "text-bg-secondary" : ""
        }`}
        onClick={() => navigate(`/certificate/${cert.id}`)}
      >
        <img
          src={cert.imgUrl}
          className="card-img-top crop"
          alt="Certificate"
        />
        <div className="card-body">
          <h5 className="card-title">{cert.title}</h5>
          <p className="card-text">
            <strong>Issued at:</strong> {cert.issuedAt.toDateString()}
            <br />
            <strong>Expired at:</strong> {cert.expiredAt.toDateString()}
          </p>
          <p className="card-text">{cert.description}</p>
        </div>
      </div>
    </div>
  );
};

export default WalletPage;
