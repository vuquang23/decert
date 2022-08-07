import image from "assets/500.jpg";
import Center from "components/Center";
import { NavigateFunction, useLocation, useNavigate } from "react-router-dom";

const ErrorPage = ({
  message,
  image,
  attrUrl,
  attrContent,
}: {
  message?: string;
  image: any;
  attrUrl: string;
  attrContent: string;
}) => {
  const navigate = useNavigate();
  return (
    <>
      <Center className="vh-100">
        <div className="h-50 d-flex flex-column align-items-center">
          <img src={image} alt="Internal error" className="h-75 d-block" />
          {message !== undefined && <p className="lead">{message}</p>}
          <button className="btn btn-primary" onClick={() => navigate("/")}>
            Go to homepage
          </button>
        </div>
      </Center>
      <a
        className="link-dark position-fixed bottom-0 end-0 m-2 small opacity-50"
        href={attrUrl}
      >
        {attrContent}
      </a>
    </>
  );
};

const DefaultErrorPage = () => {
  const location = useLocation();
  return (
    <ErrorPage
      message={
        typeof location.state === "string"
          ? location.state
          : "Please try again later!"
      }
      image={image}
      attrUrl="https://www.freepik.com/vectors/server-error"
      attrContent="Server error vector created by storyset - www.freepik.com"
    />
  );
};

const onPromiseRejected = (reason: any, navigate: NavigateFunction) => {
  let state: string | undefined = undefined;
  if (reason instanceof Error) {
    state = reason.message;
  } else if (typeof reason === "string") {
    state = reason;
  }
  return navigate("/error", { state: state });
};

export default DefaultErrorPage;
export { ErrorPage, onPromiseRejected };
