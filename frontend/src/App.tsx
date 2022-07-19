import { useMetaMask } from "components/MetaMaskProvider";
import CollectionsPage from "pages/CollectionsPage";
import HomePage from "pages/HomePage";
import { Navigate, Route, Routes } from "react-router-dom";

// TODO:

const App = () => (
  <Routes>
    <Route path="/" element={<HomePage />} />
    <Route
      path="/collections/:page"
      element={
        <RequiredMetaMask>
          <CollectionsPage />
        </RequiredMetaMask>
      }
    />
  </Routes>
);

const RequiredMetaMask = ({ children }: { children: JSX.Element }) => {
  const metaMask = useMetaMask();
  if (metaMask.account.length === 0) {
    return <Navigate to="/" />;
  }
  return children;
};

export default App;
