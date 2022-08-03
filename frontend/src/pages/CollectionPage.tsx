import { Certificate, isExpired, readAll, revoke } from "api/certificate";
import { CertificateCollection, read } from "api/certificate-collections";
import Address from "components/Address";
import { BootstrapSwalDanger } from "components/BootstrapSwal";
import { useMetaMask } from "components/MetaMaskProvider";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { Row, RowPlaceholder, Table } from "components/Table";
import { arrayFromSize } from "helper";
import { createContext, useContext, useEffect, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import Swal from "sweetalert2";

enum CertFilter {
  All,
  Valid,
  Expired,
  Revoked,
}

const ChangeCounterContext = createContext(() => {});

const CollectionPage = () => {
  const navigate = useNavigate();
  const { collectionId } = useParams();
  const [page, setPage] = useState(1);
  const [filter, setFilter] = useState(1);
  const [collection, setCollection] = useState<CertificateCollection>();
  const [certs, setCerts] = useState<Certificate[]>();
  const [changeCounter, setChangeCounter] = useState(0);
  useEffect(() => {
    setCerts(undefined);
    read(parseInt(collectionId!))
      .then((collection) => {
        setCollection(collection);
        return readAll({ collectionId: collection.id });
      })
      .then((certs) => setCerts(certs))
      .catch(() => navigate("/notfound"));
  }, [changeCounter, collectionId, navigate]);

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

interface Inputs {
  filter: CertFilter;
}

const Header = ({
  collection,
  onFilterSubmit,
  defaultFilter,
}: {
  collection?: CertificateCollection;
  onFilterSubmit: (filter: CertFilter) => void;
  defaultFilter: CertFilter;
}) => {
  const { register, handleSubmit } = useForm<Inputs>({
    defaultValues: { filter: defaultFilter },
  });
  const submitHandler: SubmitHandler<Inputs> = ({ filter }) => {
    if (typeof collection !== "undefined") {
      onFilterSubmit(filter);
    }
  };

  return (
    <div className="row align-items-center mb-5 gx-1">
      <div className="col-9">
        <h1 className="display-4 placeholder-glow">
          {typeof collection !== "undefined" ? (
            collection.title
          ) : (
            <span className="placeholder col-5" />
          )}
        </h1>
      </div>
      <form className="col-3" onBlur={handleSubmit(submitHandler)}>
        <select className="form-select" {...register("filter")}>
          <option value={CertFilter.All}>All</option>
          <option value={CertFilter.Valid}>Valid</option>
          <option value={CertFilter.Expired}>Expired</option>
          <option value={CertFilter.Revoked}>Revoked</option>
        </select>
      </form>
    </div>
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
        typeof certs !== "undefined"
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
  // Somehow "===" doesn't work here, as well as switch case

  // eslint-disable-next-line eqeqeq
  if (filter == CertFilter.Valid) {
    return typeof cert.revocation === "undefined" && !isExpired(cert);
  }
  // eslint-disable-next-line eqeqeq
  if (filter == CertFilter.Expired) {
    return typeof cert.revocation === "undefined" && isExpired(cert);
  }
  // eslint-disable-next-line eqeqeq
  if (filter == CertFilter.Revoked) {
    return typeof cert.revocation !== "undefined";
  }
  return true;
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
      cert.issuedAt.toDateString(),
      <div className={isExpired(cert) ? "text-danger fw-bold" : ""}>
        {cert.expiredAt.toDateString()}
      </div>,
      <div className="text-truncate">{cert.description}</div>,
      typeof cert.revocation === "undefined" ? (
        <RevokeButton cert={cert} />
      ) : (
        <></>
      ),
    ]}
    compactContent={
      <>
        <div className="row">
          <div className="col-10">
            <strong>Receiver:</strong> <Receiver cert={cert} />
          </div>
          {typeof cert.revocation === "undefined" && (
            <div className="col-2">
              <RevokeButton cert={cert} />
            </div>
          )}
        </div>
        <p>
          <br />
          <strong>Issued at:</strong> {cert.issuedAt.toDateString()}
          <br />
          <strong>Expired at:</strong>{" "}
          <span className={isExpired(cert) ? "text-danger fw-bold" : ""}>
            {cert.expiredAt.toDateString()}
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
          title: "Do you want to revoke this certificate?",
          showConfirmButton: true,
          showCancelButton: true,
          showLoaderOnConfirm: true,
          backdrop: true,
          icon: "warning",
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
    address={cert.receiver.address}
    customTextClassName={
      typeof cert.revocation !== "undefined"
        ? "text-dark text-decoration-line-through"
        : ""
    }
  />
);

export default CollectionPage;
