import { useMetaMask } from "components/MetaMaskProvider";
import CollectionsPage from "pages/CollectionsPage";
import HomePage from "pages/HomePage";
import PageLayout from "pages/PageLayout";
import { Navigate, Route, Routes } from "react-router-dom";

const App = () => (
  <Routes>
    <Route path="/" element={<HomePage />} />
    <Route
      element={
        <RequiredMetaMask>
          <PageLayout />
        </RequiredMetaMask>
      }
    >
      <Route path="/collections" element={<CollectionsPage />} />
    </Route>
  </Routes>
);

const RequiredMetaMask = ({ children }: { children: JSX.Element }) => {
  const metaMask = useMetaMask();
  return metaMask.account.length === 0 ? <Navigate to="/" /> : children;
};

export default App;
