import image from "assets/metamask.jpg";
import { ErrorPage } from "pages/ErrorPage";
import { Navigate } from "react-router-dom";

const InstallMetaMaskPage = () =>
  (window as any).ethereum !== undefined ? (
    <Navigate to="/" />
  ) : (
    <ErrorPage
      image={image}
      message="Please install MetaMask and reload this page"
      customButton={
        <a
          className="btn btn-warning"
          href="https://metamask.io/"
          target="_blank"
          rel="noreferrer noopener"
        >
          Install MetaMask
        </a>
      }
    />
  );

export default InstallMetaMaskPage;
