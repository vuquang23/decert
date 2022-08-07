import { Certificate, isExpired, readAll, revoke } from "api/certificate";
import { CertificateCollection, read } from "api/certificate-collections";
import Address from "components/Address";
import { BootstrapSwalDanger } from "components/BootstrapSwal";
import HeaderSearch, { Inputs } from "components/HeaderSearch";
import { useMetaMask } from "components/MetaMaskProvider";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { Row, RowPlaceholder, Table } from "components/Table";
import { arrayFromSize } from "helper";
import { onPromiseRejected } from "pages/ErrorPage";
import { NotFoundError } from "pages/NotFoundPage";
import { createContext, useContext, useEffect, useState } from "react";
import { SubmitHandler } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import Swal from "sweetalert2";

enum CertFilter {
  All = "All",
  Valid = "Valid",
  Expired = "Expired",
  Revoked = "Revoked",
}

const ChangeCounterContext = createContext(() => {});

const CollectionPage = () => {
  const navigate = useNavigate();
  const metaMask = useMetaMask();
  const { collectionId } = useParams();
  const [page, setPage] = useState(1);
  const [filter, setFilter] = useState(CertFilter.Valid);
  const [collection, setCollection] = useState<CertificateCollection>();
  const [certs, setCerts] = useState<Certificate[]>();
  const [changeCounter, setChangeCounter] = useState(0);
  useEffect(() => {
    setCerts(undefined);
    read(metaMask.address, parseInt(collectionId!))
      .then((collection) => {
        if (collection === undefined) {
          throw new NotFoundError();
        }
        setCollection(collection);
        return readAll({ collectionId: collection.id });
      })
      .then((certs) => setCerts(certs))
      .catch((reason) => {
        if (reason instanceof NotFoundError) {
          navigate("/notfound");
        }
        onPromiseRejected(reason, navigate);
      });
  }, [changeCounter, collectionId, metaMask.address, navigate]);

  return (
    <ChangeCounterContext.Provider
      value={() => setChangeCounter((prev) => prev + 1)}
    >
      <Header
        collection={collection}
        onFilterSubmit={(filter) => setFilter(filter)}
        defaultFilter={filter}
      />
      <CertsTable
        certs={certs}
        filter={filter}
        page={page}
        setPage={(page) => setPage(page)}
      />
    </ChangeCounterContext.Provider>
  );
};

const Header = ({
  collection,
  onFilterSubmit,
  defaultFilter,
}: {
  collection?: CertificateCollection;
  onFilterSubmit: (filter: CertFilter) => void;
  defaultFilter: CertFilter;
}) => {
  const navigate = useNavigate();
  const submitHandler: SubmitHandler<Inputs> = ({ filter }) => {
    if (collection !== undefined) {
      onFilterSubmit(CertFilter[filter as keyof typeof CertFilter]);
    }
  };

  return (
    <HeaderSearch
      title={
        <h1 className="display-4 placeholder-glow">
          {collection !== undefined ? (
            collection.collectionName
          ) : (
            <span className="placeholder col-5" />
          )}
        </h1>
      }
      onSearchSubmit={submitHandler}
      filters={[
        CertFilter.All,
        CertFilter.Valid,
        CertFilter.Expired,
        CertFilter.Revoked,
      ]}
      defaultFilter={defaultFilter}
      buttonText="Issue new"
      buttonIconName="plus-lg"
      buttonOnClick={() =>
        navigate("/certificate/new", {
          state: collection,
        })
      }
    />
  );
};

const CertsTable = ({
  certs,
  filter,
  page,
  setPage,
}: {
  certs?: Certificate[];
  filter: CertFilter;
  page: number;
  setPage: (page: number) => void;
}) => {
  const columnsClassName = ["col-3", "col-2", "col-2", "col-4", "col-1 d-flex"];
  return (
    <Table
      columnHeaders={["Receiver", "Issued at", "Expired at", "Description", ""]}
      columnsClassName={columnsClassName}
      itemsPerPage={5}
      page={page}
      setPage={setPage}
      rows={
        certs !== undefined
          ? certs
              .filter((cert) => doFilter(cert, filter))
              .map((cert, index) => (
                <Cert
                  key={index}
                  columnsClassName={columnsClassName}
                  cert={cert}
                />
              ))
          : arrayFromSize(5, (index) => (
              <RowPlaceholder
                key={index}
                columnsClassName={columnsClassName}
                compactContent={<ParagraphPlaceholder />}
              />
            ))
      }
    />
  );
};

const doFilter = (cert: Certificate, filter: CertFilter) => {
  switch (filter) {
    case CertFilter.Valid:
      return cert.revocation === undefined && !isExpired(cert);
    case CertFilter.Expired:
      return cert.revocation === undefined && isExpired(cert);
    case CertFilter.Revoked:
      return cert.revocation !== undefined;
    default:
      return true;
  }
};

const Cert = ({
  cert,
  columnsClassName,
}: {
  columnsClassName: string[];
  cert: Certificate;
}) => (
  <Row
    columnsClassName={columnsClassName}
    columnsValue={[
      <Receiver cert={cert} />,
      new Date(cert.issuedAt).toDateString(),
      <div className={isExpired(cert) ? "text-danger fw-bold" : ""}>
        {cert.expiredAt !== "null"
          ? new Date(cert.expiredAt).toDateString()
          : "never"}
      </div>,
      <div className="text-truncate">{cert.description}</div>,
      cert.revocation === undefined ? <RevokeButton cert={cert} /> : <></>,
    ]}
    compactContent={
      <>
        <div className="row">
          <div className="col-10">
            <strong>Receiver:</strong> <Receiver cert={cert} />
          </div>
          {cert.revocation === undefined && (
            <div className="col-2">
              <RevokeButton cert={cert} />
            </div>
          )}
        </div>
        <p>
          <br />
          <strong>Issued at:</strong> {new Date(cert.issuedAt).toDateString()}
          <br />
          <strong>Expired at:</strong>{" "}
          <span className={isExpired(cert) ? "text-danger fw-bold" : ""}>
            {cert.expiredAt !== "null"
              ? new Date(cert.expiredAt).toDateString()
              : "never"}
          </span>
        </p>
        <div className="text-truncate">{cert.description}</div>
      </>
    }
  />
);

const RevokeButton = ({ cert }: { cert: Certificate }) => {
  const updateChangeCounter = useContext(ChangeCounterContext);
  const metaMask = useMetaMask();
  return (
    <button
      className="btn btn-outline-danger ms-auto"
      onClick={() =>
        BootstrapSwalDanger.fire({
          icon: "warning",
          title: "Do you want to revoke this certificate?",
          showConfirmButton: true,
          showCancelButton: true,
          showLoaderOnConfirm: true,
          preConfirm: () => revoke(metaMask, cert.id),
          allowOutsideClick: () => !Swal.isLoading(),
        }).then((result) => {
          if (result.isConfirmed) {
            updateChangeCounter();
          }
        })
      }
    >
      <i className="bi bi-trash3" />
    </button>
  );
};

const Receiver = ({ cert }: { cert: Certificate }) => (
  <Address
    address={cert.receiver.wallet}
    customTextClassName={
      cert.revocation !== undefined
        ? "text-dark text-decoration-line-through"
        : ""
    }
  />
);

export default CollectionPage;
