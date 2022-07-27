import { Popover } from "bootstrap";
import { getShortAccount } from "components/MetaMaskProvider";
import NavBar from "components/NavBar";
import { Row, Table } from "components/Table";
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
        <CollectionsTable
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

const Header = ({ onSubmit }: { onSubmit: (searchQuery: string) => void }) => {
  interface Inputs {
    searchQuery: string;
  }

  const { register, handleSubmit } = useForm<Inputs>();
  const submitHandler: SubmitHandler<Inputs> = ({ searchQuery }) =>
    onSubmit(searchQuery);

  return (
    <div className="row align-items-center mb-5 gx-1">
      <div className="col-12 col-md-8 col-lg-7">
        <h1 className="display-4">Certificate Collections</h1>
      </div>
      <div className="col-12 col-md-4 col-lg-5">
        <div className="row gx-1">
          <div className="col-10 col-xl-8">
            <form
              className="input-group"
              onSubmit={handleSubmit(submitHandler)}
            >
              <button className="btn btn-outline-secondary" type="submit">
                <i className="bi bi-search d-inline d-lg-none" />
                <span className="d-none d-lg-inline">Search</span>
              </button>
              <input
                type="text"
                className="form-control"
                placeholder="Collection name..."
                {...register("searchQuery")}
              />
            </form>
          </div>
          <div className="col-2 col-xl-4">
            <button className="btn btn-success w-100 px-0" type="button">
              <i className="bi bi-plus-lg d-inline d-xl-none" />
              <span className="d-none d-xl-inline">New collection</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

const CollectionsTable = ({
  certCollections,
  page,
  setPage,
}: {
  certCollections: CertificateCollection[];
  page: number;
  setPage: (page: number) => void;
}) => {
  const columnsClassName = ["col-5", "col-2", "col-2", "col-3 d-flex"];
  return (
    <Table
      columnHeaders={["Certificate name", "Issued", "Revoked", ""]}
      columnsClassName={columnsClassName}
      itemsPerPage={5}
      page={page}
      setPage={setPage}
      rows={certCollections.map((item, index) => (
        <Collection
          key={index}
          columnsClassName={columnsClassName}
          certCollection={item}
        />
      ))}
    />
  );
};

const Collection = ({
  columnsClassName,
  certCollection,
}: {
  columnsClassName: string[];
  certCollection: CertificateCollection;
}) => (
  <Row
    columnsClassName={columnsClassName}
    columnsValue={[
      certCollection.certificateName,
      certCollection.issued.toString(),
      certCollection.revoked.toString(),
      <Actions certCollection={certCollection} />,
    ]}
    compactContent={
      <>
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
      </>
    }
  />
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
    <button type="button" className="btn btn-outline-dark">
      <i className="bi bi-list" />
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

export default CollectionsPage;
