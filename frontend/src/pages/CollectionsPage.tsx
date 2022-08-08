import {
  CertificateCollection,
  create,
  readAll,
} from "api/certificate-collections";
import { Popover } from "bootstrap";
import { BootstrapSwal } from "components/BootstrapSwal";
import HeaderSearch, { searchByTitle } from "components/HeaderSearch";
import {
  getShortAddress,
  MetaMask,
  useMetaMask,
} from "components/MetaMaskProvider";
import ParagraphPlaceholder from "components/ParagraphPlaceholder";
import { Row, RowPlaceholder, Table, useTableState } from "components/Table";
import { arrayFromSize, userRejectTransaction } from "helper";
import { onPromiseRejected } from "pages/ErrorPage";
import { LegacyRef, useCallback, useEffect, useRef, useState } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";

const itemsPerPage = 5;

const CollectionsPage = () => {
  const metaMask = useMetaMask();
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState("");

  const fetchCollections = useCallback(
    (itemsPerPage: number, offset: number) =>
      readAll(metaMask.address, itemsPerPage, offset, searchQuery),
    [metaMask.address, searchQuery]
  );

  const {
    array: certCollections,
    page,
    setPage,
  } = useTableState(itemsPerPage, fetchCollections, navigate);

  return (
    <>
      <HeaderSearch
        title="Certificate Collections"
        placeholder="Collection name..."
        onSearchSubmit={({ searchQuery }) => {
          setSearchQuery(searchQuery);
          setPage(1);
        }}
        buttonText="New collection"
        buttonIconName="plus-lg"
        buttonOnClick={
          certCollections !== undefined
            ? () => createCollectionModal(metaMask, navigate)
            : () => {}
        }
      />
      <CollectionsTable
        certCollections={
          certCollections !== undefined
            ? searchByTitle(
                certCollections.map((item) => ({
                  ...item,
                  title: item.collectionName,
                })),
                searchQuery
              )
            : undefined
        }
        page={page}
        setPage={(page) => setPage(page)}
      />
    </>
  );
};

const createCollectionModal = (
  metaMask: MetaMask,
  navigate: NavigateFunction
) =>
  BootstrapSwal.fire({
    title: "Enter your new collection name",
    input: "text",
    confirmButtonText: "Create",
    showLoaderOnConfirm: true,

    inputValidator: (result) =>
      result.length === 0 ? "Collection name cannot be empty" : null,

    preConfirm: (collectionName) => create(metaMask, collectionName),
    allowOutsideClick: () => !Swal.isLoading(),
  })
    .then((result) => {
      if (result.isConfirmed) {
        BootstrapSwal.fire({
          icon: "success",
          title: "Collection created!",
          showConfirmButton: false,
          showCloseButton: true,
        });
      }
    })
    .catch((reason) => {
      if (userRejectTransaction(reason)) {
        BootstrapSwal.fire({
          icon: "error",
          title: "You rejected the transaction.",
        });
      } else {
        onPromiseRejected(reason, navigate);
      }
    });

//#region Collections Table

const CollectionsTable = ({
  certCollections,
  page,
  setPage,
}: {
  certCollections?: CertificateCollection[];
  page: number;
  setPage: (page: number) => void;
}) => {
  const columnsClassName = ["col-5", "col-2", "col-2", "col-3 d-flex"];
  return (
    <Table
      columnHeaders={["Certificate name", "Issued", "Revoked", ""]}
      columnsClassName={columnsClassName}
      itemsPerPage={itemsPerPage}
      page={page}
      setPage={setPage}
      rows={
        certCollections !== undefined
          ? certCollections.map((item, index) => (
              <Collection
                key={index}
                columnsClassName={columnsClassName}
                certCollection={item}
              />
            ))
          : arrayFromSize(itemsPerPage, (index) => (
              <CollectionPlaceholder
                key={index}
                columnsClassName={columnsClassName}
              />
            ))
      }
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
      certCollection.collectionName,
      certCollection.totalIssued.toString(),
      certCollection.totalRevoked.toString(),
      <Actions certCollection={certCollection} />,
    ]}
    compactContent={
      <>
        <h3 className="fw-bold pt-3">{certCollection.collectionName}</h3>
        <div className="row align-items-center row justify-content-between">
          <div className="col-6 col-sm-3 pb-2 pb-sm-0">
            <strong>Issued:</strong> {certCollection.totalIssued}
          </div>
          <div className="col-6 col-sm-3 pb-2 pb-sm-0">
            <strong>Revoked:</strong> {certCollection.totalRevoked}
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
}) => {
  const navigate = useNavigate();
  return (
    <div className="btn-group ms-sm-auto">
      <button
        type="button"
        className="btn btn-outline-dark"
        onClick={() =>
          navigate("/certificate/new", {
            state: certCollection,
          })
        }
      >
        <i className="bi bi-plus-lg" />
      </button>
      <button
        type="button"
        className="btn btn-outline-dark"
        onClick={() => navigate(certCollection.id.toString())}
      >
        <i className="bi bi-list" />
      </button>
      <ContractAddressButton certCollection={certCollection} />
    </div>
  );
};

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
  }, []);

  return (
    <button
      type="button"
      className="btn btn-outline-dark"
      data-bs-toggle="popover"
      data-bs-title="Contract address"
      data-bs-content={contractAddressPopoverHTML(
        certCollection.collectionAddress
      )}
      ref={popoverRef}
    >
      <i className="bi bi-wallet2" />
    </button>
  );
};

const contractAddressPopoverHTML = (address: string) =>
  `
    <div class="d-flex align-items-center">
      <div class="font-monospace me-2">
        ${getShortAddress(address)}
      </div>
      <button
        type="button"
        class="btn btn-light"
        onclick="navigator.clipboard.writeText('${address}')"
        >
        <svg xmlns="http://www.w3.org/2000/svg" width="14 " height="14" fill="currentColor" class="bi bi-clipboard align-baseline" viewBox="0 0 16 16">
          <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"/>
          <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"/>
        </svg>
      </button>
    </div>
  `;

const CollectionPlaceholder = ({
  columnsClassName,
}: {
  columnsClassName: string[];
}) => (
  <RowPlaceholder
    columnsClassName={columnsClassName}
    compactContent={
      <>
        <div className="h3 placeholder col-4 pt-3" />
        <ParagraphPlaceholder />
      </>
    }
  />
);

//#endregion

export default CollectionsPage;
