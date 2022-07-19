import { ReactComponent as Logo } from "assets/global.svg";
import { Link } from "react-router-dom";
import { useMetaMask } from "./MetaMaskProvider";

const NavBar = ({ transparent }: { transparent?: boolean }) => {
  const metaMask = useMetaMask();
  const isConnected = metaMask.account.length > 0;
  const background =
    typeof transparent === "undefined" || transparent === false
      ? "bg-dark"
      : "";

  return (
    <nav className={`navbar navbar-dark ${background} z-index-2 py-4`}>
      <div className="container">
        <Link to="/" className="navbar-brand fw-bold">
          <Logo className="d-inline-block align-text-bottom" /> DECERT
        </Link>
        <ul className="navbar-nav">
          <li className="nav-item">
            <Link
              to="/collections/1"
              className={`nav-link fw-semibold ${isConnected ? "" : "d-none"}`}
            >
              Certificates
            </Link>
          </li>
        </ul>
        <button
          className={`btn btn-outline-light ${isConnected ? "d-none" : ""}`}
          type="button"
          onClick={() => metaMask.connectToMetaMask()}
        >
          Connect to MetaMask
        </button>
      </div>
    </nav>
  );
};

export default NavBar;
