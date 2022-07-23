import { Popover } from "bootstrap";
import { getShortAccount } from "components/MetaMaskProvider";
import NavBar from "components/NavBar";
import { LegacyRef, useEffect, useRef, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";

interface CertificateCollection {
  certificateName: string;
  address: string;
  issued: number;
  revoked: number;
}

const certCollections: CertificateCollection[] = Array.from(
  Array(30).keys(),
  (_, index) => ({
    certificateName: `Sinh viên ${index} tốt`,
    address: `0xb${index}a904b0E45Cd99Ef4D9C9C6cb11f293bD77cB7`,
    issued: 30,
    revoked: 20,
  })
);

const CollectionsPage = () => {
  const [searchQuery, setSearchQuery] = useState("");
  const [page, setPage] = useState(1);
  return (
    <>
      <NavBar />
      <div className="container mt-5">
        <Header
          onSubmit={(searchQuery) => {
            setSearchQuery(searchQuery);
            setPage(1);
          }}
        />
        <Table
          certCollections={certCollections.filter(
            (e) =>
              searchQuery.length === 0 ||
              e.certificateName.toLowerCase().trim().includes(searchQuery)
          )}
          page={page}
          setPage={(page) => setPage(page)}
        />
      </div>
    </>
  );
};

interface Inputs {
  searchQuery: string;
}

const Header = ({ onSubmit }: { onSubmit: (searchQuery: string) => void }) => {
  const { register, handleSubmit } = useForm<Inputs>();
  const submitHandler: SubmitHandler<Inputs> = ({ searchQuery }) =>
    onSubmit(searchQuery);

  return (
    <div className="row align-items-center mb-5">
      <div className="col-12 col-md-8">
        <h1 className="display-4">Certificate Collections</h1>
      </div>
      <div className="col-12 col-md-4">
        <form className="input-group" onSubmit={handleSubmit(submitHandler)}>
          <input
            type="text"
            className="form-control"
            placeholder="Collection name..."
            {...register("searchQuery")}
          />
          <button className="btn btn-success fw-semibold" type="submit">
            Search
          </button>
        </form>
      </div>
    </div>
  );
};

const itemsPerPage = 5;

const Table = ({
  certCollections,
  page,
  setPage,
}: {
  certCollections: CertificateCollection[];
  page: number;
  setPage: (page: number) => void;
}) => {
  const numOfPages = Math.ceil(certCollections.length / itemsPerPage);
  const columnsClassName = ["col-6", "col-2", "col-2", "col-2 d-flex"];
  return (
    <>
      <div className="card border-0">
        <div className="card-body row fw-bold align-items-center d-none d-md-flex">
          <div className={columnsClassName[0]}>Certificate name</div>
          <div className={columnsClassName[1]}>Issued</div>
          <div className={columnsClassName[2]}>Revoked</div>
          <div className={columnsClassName[3]} />
        </div>
      </div>
      {certCollections
        .slice(
          (page - 1) * itemsPerPage,
          (page - 1) * itemsPerPage + itemsPerPage
        )
        .map((item, index) => (
          <Row
            key={index}
            columnsClassName={columnsClassName}
            certCollection={item}
          />
        ))}
      <Pagination page={page} numOfPages={numOfPages} setPage={setPage} />
    </>
  );
};

const Row = ({
  columnsClassName,
  certCollection,
}: {
  columnsClassName: string[];
  certCollection: CertificateCollection;
}) => (
  <div className="card my-1">
    <div className="card-body row align-items-center d-none d-md-flex">
      <div className={columnsClassName[0]}>
        {certCollection.certificateName}
      </div>
      <div className={columnsClassName[1]}>{certCollection.issued}</div>
      <div className={columnsClassName[2]}>{certCollection.revoked}</div>
      <div className={columnsClassName[3]}>
        <Actions certCollection={certCollection} />
      </div>
    </div>
    <div className="card-body d-block d-md-none">
      <h3 className="fw-bold pt-3">{certCollection.certificateName}</h3>
      <div className="row align-items-center row justify-content-between">
        <div className="col-6 col-sm-3 pb-2 pb-sm-0">
          <strong>Issued:</strong> {certCollection.issued}
        </div>
        <div className="col-6 col-sm-3 pb-2 pb-sm-0">
          <strong>Revoked:</strong> {certCollection.revoked}
        </div>
        <div className="col-6 d-flex">
          <Actions certCollection={certCollection} />
        </div>
      </div>
    </div>
  </div>
);

const Actions = ({
  certCollection,
}: {
  certCollection: CertificateCollection;
}) => (
  <div className="btn-group ms-sm-auto">
    <button type="button" className="btn btn-outline-dark">
      <i className="bi bi-plus-lg" />
    </button>
    <ContractAddressButton certCollection={certCollection} />
  </div>
);

const ContractAddressButton = ({
  certCollection,
}: {
  certCollection: CertificateCollection;
}) => {
  const popoverRef =
    useRef<HTMLButtonElement>() as LegacyRef<HTMLButtonElement>;

  useEffect(() => {
    new Popover(
      (popoverRef as React.MutableRefObject<HTMLButtonElement>).current,
      {
        html: true,
        sanitize: false,
        container: "body",
        trigger: "focus",
      }
    );
  });

  return (
    <button
      type="button"
      className="btn btn-outline-dark"
      data-bs-toggle="popover"
      data-bs-title="Contract address"
      data-bs-content={`
          <div class="d-flex align-items-center">
            <div class="font-monospace me-2">
              ${getShortAccount(certCollection.address)}
            </div>
            <button
              type="button"
              class="btn btn-light"
              onclick="navigator.clipboard.writeText('${
                certCollection.address
              }')"
              >
              <svg xmlns="http://www.w3.org/2000/svg" width="14 " height="14" fill="currentColor" class="bi bi-clipboard align-baseline" viewBox="0 0 16 16">
                <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>
                <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>
              </svg>
            </button>
          </div>
        `}
      ref={popoverRef}
    >
      <i className="bi bi-wallet2" />
    </button>
  );
};

const Pagination = ({
  page,
  numOfPages,
  setPage,
}: {
  page: number;
  numOfPages: number;
  setPage: (page: number) => void;
}) => (
  <ul className="pagination mt-4 justify-content-center">
    {page > 1 ? (
      <li className="page-item">
        <button className="page-link" onClick={() => setPage(page - 1)}>
          Previous
        </button>
      </li>
    ) : null}
    {Array.from(Array(numOfPages + 1).keys())
      .slice(Math.max(1, page - 1), Math.max(1, page - 1) + 3)
      .map((value) => (
        <li
          key={value}
          className={`page-item ${value === page ? "active" : ""}`}
        >
          <button className="page-link" onClick={() => setPage(value)}>
            {value}
          </button>
        </li>
      ))}
    {page < numOfPages ? (
      <li className="page-item">
        <button className="page-link" onClick={() => setPage(page + 1)}>
          Next
        </button>
      </li>
    ) : null}
  </ul>
);

export default CollectionsPage;
