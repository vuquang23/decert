import { ReactComponent as Logo } from "assets/global.svg";
import { useState } from "react";
import { Link, NavLink } from "react-router-dom";
import { getShortAddress, useMetaMask } from "./MetaMaskProvider";

const NavBar = ({ transparent }: { transparent?: boolean }) => {
  const isConnected = useMetaMask().address.length > 0;
  const background =
    transparent === undefined || transparent === false
      ? "bg-dark"
      : isConnected
      ? "bg-dark-sm"
      : "";

  return (
    <nav
      className={`navbar navbar-expand-sm navbar-dark ${background} z-index-2 py-4`}
    >
      <div className="container">
        <Link to="/" className="navbar-brand fw-bold">
          <Logo className="d-inline-block align-text-bottom" /> DECERT
        </Link>
        {isConnected ? <ConnectedNav /> : <NotConnectedNav />}
      </div>
    </nav>
  );
};

const ConnectedNav = () => (
  <>
    <button
      className="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target="#navbar-collapse"
    >
      <span className="navbar-toggler-icon"></span>
    </button>
    <div className="collapse navbar-collapse" id="navbar-collapse">
      <ul className="navbar-nav ms-auto">
        <li className="nav-item">
          <NavLink
            to="/collections"
            className={`nav-link fw-semibold ${({
              isActive,
            }: {
              isActive: boolean;
            }) => (isActive ? "active" : "")}`}
          >
            Collections
          </NavLink>
        </li>
        <li className="nav-item">
          <NavLink
            to="/wallet"
            className={`nav-link fw-semibold ${({
              isActive,
            }: {
              isActive: boolean;
            }) => (isActive ? "active" : "")}`}
          >
            Wallet
          </NavLink>
        </li>
      </ul>
      <hr className="border-light d-block d-sm-none" />
      <AccountInfo />
    </div>
  </>
);

const AccountInfo = () => {
  const [balance, setBalance] = useState("0");
  const metaMask = useMetaMask();

  // TODO: handle big value
  metaMask.getBalance().then((value) => setBalance(value));

  return (
    <div className="input-group w-auto ms-sm-2">
      <span className="input-group-text text-light bg-transparent font-monospace">
        {getShortAddress(metaMask.address)}
      </span>
      <span className="input-group-text">{balance}</span>
    </div>
  );
};

const NotConnectedNav = () => {
  const metaMask = useMetaMask();
  return (
    <button
      className="btn btn-outline-light ms-auto"
      type="button"
      onClick={() => metaMask.connectToMetaMask()}
    >
      Connect to MetaMask
    </button>
  );
};

export default NavBar;
