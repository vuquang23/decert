import Empty from "components/Empty";
import { getRandomCol } from "components/ParagraphPlaceholder";
import { onPromiseRejected } from "pages/ErrorPage";
import { useEffect, useState } from "react";
import { NavigateFunction } from "react-router-dom";

const Table = ({
  columnHeaders,
  columnsClassName,
  rows,
  page,
  itemsPerPage,
  setPage,
}: {
  columnHeaders: string[];
  columnsClassName: string[];
  rows: JSX.Element[];
  page: number;
  itemsPerPage: number;
  setPage: (page: number) => void;
}) =>
  rows.length === 0 ? (
    <Empty />
  ) : (
    <>
      <div className="card border-0">
        <div className="card-body row fw-bold align-items-center d-none d-md-flex">
          {columnHeaders.map((header, index) => (
            <div key={index} className={columnsClassName[index]}>
              {header}
            </div>
          ))}
        </div>
      </div>
      {rows.slice(
        (page - 1) * itemsPerPage,
        (page - 1) * itemsPerPage + itemsPerPage
      )}
      <Pagination
        page={page}
        numOfPages={Math.ceil(rows.length / itemsPerPage)}
        setPage={setPage}
      />
    </>
  );

const Row = ({
  columnsClassName,
  columnsValue,
  compactContent,
}: {
  columnsClassName: string[];
  columnsValue: (JSX.Element | string)[];
  compactContent: JSX.Element;
}) => (
  <div className="card my-1">
    <div className="card-body row align-items-center d-none d-md-flex">
      {columnsValue.map((item, index) => (
        <div key={index} className={columnsClassName[index]} children={item} />
      ))}
    </div>
    <div className="card-body d-block d-md-none" children={compactContent} />
  </div>
);

const RowPlaceholder = ({
  columnsClassName,
  compactContent,
}: {
  columnsClassName: string[];
  compactContent: JSX.Element;
}) => (
  <div className="card my-1 placeholder-glow">
    <div className="card-body row align-items-center d-none d-md-flex">
      {columnsClassName.map((_, index) => (
        <div key={index} className={columnsClassName[index]}>
          <span className={`placeholder col-${getRandomCol()}`} />
        </div>
      ))}
    </div>
    <div className="card-body d-block d-md-none" children={compactContent} />
  </div>
);

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
    {page > 1 && (
      <li className="page-item">
        <button className="page-link" onClick={() => setPage(page - 1)}>
          Previous
        </button>
      </li>
    )}
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
    {page < numOfPages && (
      <li className="page-item">
        <button className="page-link" onClick={() => setPage(page + 1)}>
          Next
        </button>
      </li>
    )}
  </ul>
);

/**
 * {@link fetch} function should be useCallback()'s output. If {@link fetch} is
 * changed, page will be set to 1 and array will be fetched again.
 */
const useTableState = <T,>(
  itemsPerPage: number,
  fetch: (itemsPerPage: number, offset: number) => Promise<T[]>,
  navigate: NavigateFunction
) => {
  const [array, setArray] = useState<T[]>();
  const [page, setPage] = useState(1);

  useEffect(() => {
    setPage(1);
    setArray(undefined);
    fetch(itemsPerPage * 4, 0)
      .then((value) => setArray(value))
      .catch((reason) => onPromiseRejected(reason, navigate));
  }, [fetch, itemsPerPage, navigate]);

  return {
    array: array,
    page: page,
    setPage: (page: number) =>
      setPage((prevPage) => {
        if (page > prevPage) {
          fetch(itemsPerPage, array !== undefined ? array.length : 0)
            .then((value) => {
              if (value.length > 0) {
                setArray((prevValue) => [...(prevValue ?? []), ...value]);
              }
            })
            .catch((reason) => onPromiseRejected(reason, navigate));
        }
        return page;
      }),
  };
};

export { Table, Row, RowPlaceholder, Pagination, useTableState };
