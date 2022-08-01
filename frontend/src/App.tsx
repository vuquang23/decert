import { useMetaMask } from "components/MetaMaskProvider";
import CertificatePage from "pages/CertificatePage";
import CollectionsPage from "pages/CollectionsPage";
import HomePage from "pages/HomePage";
import NotFoundPage from "pages/NotFoundPage";
import PageLayout from "pages/PageLayout";
import WalletPage from "pages/WalletPage";
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
      <Route path="/wallet" element={<WalletPage />} />
    </Route>
    <Route path="/certificate/:certId" element={<CertificatePage />} />
    <Route path="*" element={<NotFoundPage />} />
  </Routes>
);

const RequiredMetaMask = ({ children }: { children: JSX.Element }) =>
  useMetaMask().account.length === 0 ? <Navigate to="/" /> : children;

export default App;
