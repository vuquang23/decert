import { Certificate, isExpired, readAll } from "api/certificate";
import { CertificateCollection, read } from "api/certificate-collections";
import { getShortAccount } from "components/MetaMaskProvider";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { Row, RowPlaceholder, Table } from "components/Table";
import { arrayFromSize } from "helper";
import { useEffect, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";

enum CertFilter {
  All,
  Valid,
  Expired,
  Revoked,
}

const CollectionPage = () => {
  const navigate = useNavigate();
  const { collectionId } = useParams();
  const [page, setPage] = useState(1);
  const [filter, setFilter] = useState(1);
  const [collection, setCollection] = useState<CertificateCollection>();
  const [certs, setCerts] = useState<Certificate[]>();
  useEffect(() => {
    read(parseInt(collectionId!))
      .then((collection) => {
        setCollection(collection);
        return readAll({ collectionId: collection.id });
      })
      .then((certs) => setCerts(certs))
      .catch(() => navigate("/notfound"));
  });

  return (
    <>
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
    </>
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
  const columnsClassName = ["col-3", "col-2", "col-2", "col-5"];
  return (
    <Table
      columnHeaders={["Receiver", "Issued at", "Expired at", "Description"]}
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
    return typeof cert.revokedAt === "undefined" && !isExpired(cert);
  }
  // eslint-disable-next-line eqeqeq
  if (filter == CertFilter.Expired) {
    return typeof cert.revokedAt === "undefined" && isExpired(cert);
  }
  // eslint-disable-next-line eqeqeq
  if (filter == CertFilter.Revoked) {
    return typeof cert.revokedAt !== "undefined";
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
    ]}
    compactContent={
      <>
        <p>
          <strong>Receiver:</strong> <Receiver cert={cert} />
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

const Receiver = ({ cert }: { cert: Certificate }) => (
  <>
    <code
      className={
        typeof cert.revokedAt !== "undefined"
          ? "text-dark text-decoration-line-through"
          : ""
      }
    >
      {getShortAccount(cert.receiver)}
    </code>{" "}
    <button
      type="button"
      className="btn btn-light"
      onClick={() => navigator.clipboard.writeText(cert.receiver)}
    >
      <i className="bi bi-clipboard align-baseline" />
    </button>
  </>
);

export default CollectionPage;
