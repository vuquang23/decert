import { ReactComponent as Logo } from "assets/global.svg";
import { useNavigate } from "react-router-dom";

const NavBar = (props: { isConnected?: boolean; transparent?: boolean }) => {
  const navigate = useNavigate();
  return (
    <nav
      className={`navbar navbar-dark ${
        typeof props.transparent === "undefined" || props.transparent === false
          ? "bg-dark"
          : ""
      } navbar-expand-md z-index-2 py-4`}
    >
      <div className="container">
        <a className="navbar-brand fw-bold" href="/">
          <Logo className="d-inline-block align-text-bottom" /> DECERT
        </a>
        <button
          className="btn btn-outline-light ms-auto"
          type="button"
          onClick={() => navigate(props.isConnected ? "/" : "/issue/1")}
        >
          {props.isConnected
            ? "Disconnect from MetaMask"
            : "Connect to MetaMask"}
        </button>
      </div>
    </nav>
  );
};

export default NavBar;
