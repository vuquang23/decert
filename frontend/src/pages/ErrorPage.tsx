import { ReactComponent as Image } from "assets/500.svg";
import Center from "components/Center";
import { ReactNode } from "react";
import { NavigateFunction, useLocation, useNavigate } from "react-router-dom";

const ErrorPage = ({
  message,
  svg,
  customButton,
  attrUrl,
  attrContent,
}: {
  message?: string;
  svg: ReactNode;
  customButton?: ReactNode;
  attrUrl?: string;
  attrContent?: string;
}) => {
  const navigate = useNavigate();
  return (
    <>
      <Center className="vh-100">
        <div className="h-50 d-flex flex-column align-items-center">
          <div className="h-75">{svg}</div>
          {message !== undefined && (
            <p className="lead text-center">{message}</p>
          )}
          {customButton ?? (
            <button className="btn btn-primary" onClick={() => navigate("/")}>
              Go to homepage
            </button>
          )}
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
      svg={<Image className="h-100" />}
      attrUrl="https://storyset.com/internet"
      attrContent="Internet illustrations by Storyset"
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
