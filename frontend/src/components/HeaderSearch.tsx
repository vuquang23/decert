import { ReactNode } from "react";
import { SubmitHandler, useForm } from "react-hook-form";

interface Inputs {
  searchQuery: string;
  filter: string;
}

/**
 * To enable search, {@link placeholder} must be defined.
 *
 * To enable button, {@link buttonOnClick} must be defined. {@link buttonIconName}
 * must be a valid {@link https://icons.getbootstrap.com/ Bootstrap icon} (e.g.
 * "plus-lg").
 *
 * To enable search filter, {@link filters} must be defined. First element will
 * be default value if {@link defaultFilter} is not specified.
 */
const HeaderSearch = ({
  title,
  placeholder,
  filters,
  defaultFilter,
  onSearchSubmit,
  buttonText,
  buttonIconName,
  buttonOnClick,
}: {
  title: string | ReactNode;
  placeholder?: string;
  filters?: string[];
  defaultFilter?: string;
  onSearchSubmit: (inputs: Inputs) => void;
  buttonText?: string;
  buttonIconName?: string;
  buttonOnClick?: () => void;
}) => (
  <div className="row align-items-center mb-5 gx-1">
    <div className="col-12 col-md-8 col-lg-7">
      {typeof title === "string" ? (
        <h1 className="display-4">{title}</h1>
      ) : (
        title
      )}
    </div>
    <div className="col-12 col-md-4 col-lg-5">
      {buttonOnClick === undefined ? (
        <SearchForm
          placeholder={placeholder}
          filters={filters}
          defaultFilter={defaultFilter}
          onSearchSubmit={onSearchSubmit}
        />
      ) : (
        <div className="row gx-1">
          <div className="col-10 col-xl-8">
            <SearchForm
              placeholder={placeholder}
              filters={filters}
              defaultFilter={defaultFilter}
              onSearchSubmit={onSearchSubmit}
            />
          </div>
          <div className="col-2 col-xl-4">
            <button
              className="btn btn-success w-100 px-0"
              type="button"
              onClick={buttonOnClick}
            >
              <i className={`bi bi-${buttonIconName} d-inline d-xl-none`} />
              <span className="d-none d-xl-inline">{buttonText}</span>
            </button>
          </div>
        </div>
      )}
    </div>
  </div>
);

const SearchForm = ({
  placeholder,
  filters,
  defaultFilter,
  onSearchSubmit,
}: {
  placeholder?: string;
  filters?: string[];
  defaultFilter?: string;
  onSearchSubmit: (inputs: Inputs) => void;
}) => {
  const { register, handleSubmit } = useForm<Inputs>({
    defaultValues: { filter: defaultFilter },
  });
  const submitHandler: SubmitHandler<Inputs> = (inputs) =>
    onSearchSubmit(inputs);

  return (
    <form
      className="input-group"
      onSubmit={handleSubmit(submitHandler)}
      onBlur={handleSubmit(submitHandler)}
    >
      {placeholder && (
        <>
          <button className="btn btn-outline-secondary" type="submit">
            <i className="bi bi-search d-inline d-lg-none" />
            <span className="d-none d-lg-inline">Search</span>
          </button>
          <input
            type="text"
            className="form-control flex-grow-2"
            placeholder={placeholder}
            {...register("searchQuery")}
          />
        </>
      )}
      {Array.isArray(filters) && (
        <select
          className="form-select"
          {...register("filter", { onChange: handleSubmit(submitHandler) })}
        >
          {filters.map((option, index) => (
            <option key={index} value={option}>
              {option}
            </option>
          ))}
        </select>
      )}
    </form>
  );
};

const searchByTitle = <T extends { title: String }>(
  array: T[],
  searchQuery: string
) =>
  array.filter(
    (value) =>
      searchQuery.length === 0 ||
      value.title
        .trim()
        .toLocaleLowerCase()
        .includes(searchQuery.trim().toLocaleLowerCase())
  );

export default HeaderSearch;
export type { Inputs };
export { searchByTitle };
