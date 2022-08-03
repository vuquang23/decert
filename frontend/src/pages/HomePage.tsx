import Center from "components/Center";
import NavBar from "components/NavBar";
import { formValidationClassName } from "helper";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

const HomePage = () => (
  <div className="vh-100">
    <NavBar transparent />
    <MainContent />
    <ImageAttribution />
  </div>
);

interface Inputs {
  searchQuery: string;
}

const MainContent = () => {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();
  const submitHandler: SubmitHandler<Inputs> = ({ searchQuery }) =>
    navigate(`certificate/${searchQuery.match(/\d+$/g)?.at(0)}`);

  return (
    <Center className="h-100 w-100 fixed-top z-index-1 homepage-bg">
      <div className="container">
        <h1 className="display-2 text-light text-center mb-4">
          Decert Verifier
        </h1>
        <form
          className="input-group input-group-lg"
          onSubmit={handleSubmit(submitHandler)}
        >
          <input
            type="text"
            className={`form-control ${formValidationClassName(
              errors.searchQuery
            )}`}
            placeholder="Certificate URL or Certificate ID"
            {...register("searchQuery", {
              required: true,
              pattern: /^((http:\/\/localhost:3000\/certificate\/\d+)|(\d+))$/g,
            })}
          />
          <button className="btn btn-success" type="submit">
            <span className="mx-3 fw-bold">Verify</span>
          </button>
        </form>
      </div>
    </Center>
  );
};

const ImageAttribution = () => {
  const authorUrl =
    "https://unsplash.com/@hjrc33?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText";
  const imageUrl =
    "https://unsplash.com/s/photos/business?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText";
  const iconUrl = "https://www.flaticon.com/free-icons/global";

  return (
    <p className="position-fixed bottom-0 end-0 me-2 z-index-2 text-light small opacity-50">
      Photo by{" "}
      <a className="link-light" href={authorUrl}>
        Héctor J. Rivas
      </a>{" "}
      on{" "}
      <a className="link-light" href={imageUrl}>
        Unsplash
      </a>{" "}
      &bull;{" "}
      <a className="link-light" href={iconUrl}>
        Global icons created by Freepik - Flaticon
      </a>
    </p>
  );
};

export default HomePage;
