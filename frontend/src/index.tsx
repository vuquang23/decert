import App from "App";
import "bootstrap";
import "bootstrap-icons/font/bootstrap-icons.css";
import "bootstrap/dist/css/bootstrap.min.css";
import MetaMaskProvider from "components/MetaMaskProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import "styles/styles.css";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <MetaMaskProvider>
        <App />
      </MetaMaskProvider>
    </BrowserRouter>
  </React.StrictMode>
);
