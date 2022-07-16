import NavBar from "../components/NavBar";
import { useParams, Link } from "react-router-dom";

const itemsPerPage = 5;

const certificateCollections = Array.from(Array(30).keys(), (_, index) => ({
  certificateName: `Sinh viên ${index} tốt`,
  quantity: 30,
}));

const numOfPages = Math.ceil(certificateCollections.length / itemsPerPage);

const IssuePage = () => (
  <>
    <NavBar isConnected />
    <MainContent />
  </>
);

const MainContent = () => {
  const { page } = useParams();
  return (
    <div className="container mt-5">
      <h1 className="display-4 mb-4">Hello, VNU</h1>
      <Table page={parseInt(page!)} />
    </div>
  );
};

const Table = (props: { page: number }) => {
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
          (props.page - 1) * itemsPerPage,
          (props.page - 1) * itemsPerPage + itemsPerPage
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
      <Pagination page={props.page} />
    </>
  );
};

const Row = (props: {
  firstCol: string;
  secondCol: string;
  thirdCol: string;
  certificateName: string;
  quantity: number;
}) => (
  <div className="card my-1">
    <div className="card-body row align-items-center d-none d-md-flex">
      <div className={props.firstCol}>{props.certificateName}</div>
      <div className={props.secondCol}>{props.quantity}</div>
      <div className={props.thirdCol}>
        <Actions />
      </div>
    </div>
    <div className="card-body d-block d-md-none">
      <h3 className="fw-bold pt-3">{props.certificateName}</h3>
      <div className="row align-items-center row justify-content-between">
        <div className="col-6 pb-2 pb-sm-0">
          <strong>Issued:</strong> {props.quantity}
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

const Pagination = (props: { page: number }) => (
  <ul className="pagination mt-4 justify-content-center">
    {props.page > 1 ? (
      <li className="page-item">
        <Link to={`/issue/${props.page - 1}`} className="page-link">
          Previous
        </Link>
      </li>
    ) : null}
    {Array.from(Array(numOfPages + 1).keys())
      .slice(Math.max(1, props.page - 1), Math.max(1, props.page - 1) + 3)
      .map((value) => (
        <li
          key={value}
          className={`page-item ${value === props.page ? "active" : ""}`}
        >
          <Link to={`/issue/${value}`} className="page-link">
            {value}
          </Link>
        </li>
      ))}
    {props.page < numOfPages ? (
      <li className="page-item">
        <Link to={`/issue/${props.page + 1}`} className="page-link">
          Next
        </Link>
      </li>
    ) : null}
  </ul>
);

export default IssuePage;
