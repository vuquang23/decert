import NavBar from "components/NavBar";
import { Link, useParams } from "react-router-dom";

const itemsPerPage = 5;

const certificateCollections = Array.from(Array(30).keys(), (_, index) => ({
  certificateName: `Sinh viên ${index} tốt`,
  quantity: 30,
}));

const numOfPages = Math.ceil(certificateCollections.length / itemsPerPage);

const CollectionsPage = () => {
  const { page } = useParams();
  return (
    <>
      <NavBar />
      <div className="container mt-5">
        <Table page={parseInt(page!)} />
      </div>
    </>
  );
};

const Table = ({ page }: { page: number }) => {
  const firstCol = "col-7";
  const secondCol = "col-2";
  const thirdCol = "col-3 d-flex";
  return (
    <>
      <div className="card border-0">
        <div className="card-body row fw-bold align-items-center d-none d-md-flex">
          <div className={firstCol}>Certificate name</div>
          <div className={secondCol}>Issued</div>
          <div className={thirdCol} />
        </div>
      </div>
      {certificateCollections
        .slice(
          (page - 1) * itemsPerPage,
          (page - 1) * itemsPerPage + itemsPerPage
        )
        .map((item, index) => (
          <Row
            key={index}
            firstCol={firstCol}
            secondCol={secondCol}
            thirdCol={thirdCol}
            certificateName={item.certificateName}
            quantity={item.quantity}
          />
        ))}
      <Pagination page={page} />
    </>
  );
};

const Row = ({
  firstCol,
  secondCol,
  thirdCol,
  certificateName,
  quantity,
}: {
  firstCol: string;
  secondCol: string;
  thirdCol: string;
  certificateName: string;
  quantity: number;
}) => (
  <div className="card my-1">
    <div className="card-body row align-items-center d-none d-md-flex">
      <div className={firstCol}>{certificateName}</div>
      <div className={secondCol}>{quantity}</div>
      <div className={thirdCol}>
        <Actions />
      </div>
    </div>
    <div className="card-body d-block d-md-none">
      <h3 className="fw-bold pt-3">{certificateName}</h3>
      <div className="row align-items-center row justify-content-between">
        <div className="col-6 pb-2 pb-sm-0">
          <strong>Issued:</strong> {quantity}
        </div>
        <div className="col-7 col-sm-6 d-flex">
          <Actions />
        </div>
      </div>
    </div>
  </div>
);

const Actions = () => (
  <div className="btn-group ms-sm-auto">
    <button type="button" className="btn btn-outline-dark">
      <i className="bi bi-plus-lg" />
    </button>
    <button type="button" className="btn btn-outline-dark">
      <i className="bi bi-list-ul" />
    </button>
  </div>
);

const Pagination = ({ page }: { page: number }) => (
  <ul className="pagination mt-4 justify-content-center">
    {page > 1 ? (
      <li className="page-item">
        <Link to={`/collections/${page - 1}`} className="page-link">
          Previous
        </Link>
      </li>
    ) : null}
    {Array.from(Array(numOfPages + 1).keys())
      .slice(Math.max(1, page - 1), Math.max(1, page - 1) + 3)
      .map((value) => (
        <li
          key={value}
          className={`page-item ${value === page ? "active" : ""}`}
        >
          <Link to={`/collections/${value}`} className="page-link">
            {value}
          </Link>
        </li>
      ))}
    {page < numOfPages ? (
      <li className="page-item">
        <Link to={`/collections/${page + 1}`} className="page-link">
          Next
        </Link>
      </li>
    ) : null}
  </ul>
);

export default CollectionsPage;
