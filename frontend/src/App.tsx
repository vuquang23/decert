import { useMetaMask } from "components/MetaMaskProvider";
import CertificatePage from "pages/CertificatePage";
import CollectionPage from "pages/CollectionPage";
import CollectionsPage from "pages/CollectionsPage";
import DefaultErrorPage from "pages/ErrorPage";
import HomePage from "pages/HomePage";
import InstallMetaMaskPage from "pages/InstallMetaMaskPage";
import NewCertificatePage from "pages/NewCertificatePage";
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
      <Route path="/collections">
        <Route index element={<CollectionsPage />} />
        <Route path=":collectionId" element={<CollectionPage />} />
      </Route>
      <Route path="/certificate/new" element={<NewCertificatePage />} />
      <Route path="/wallet" element={<WalletPage />} />
    </Route>
    <Route path="/certificate/:certId" element={<CertificatePage />} />
    <Route path="/error" element={<DefaultErrorPage />} />
    <Route path="/install-metamask" element={<InstallMetaMaskPage />} />
    <Route path="*" element={<NotFoundPage />} />
  </Routes>
);

const RequiredMetaMask = ({ children }: { children: JSX.Element }) => {
  const metaMask = useMetaMask();
  return metaMask.isReady ? (
    metaMask.address.length === 0 ? (
      <Navigate to="/" replace />
    ) : (
      children
    )
  ) : (
    <div />
  );
};

export default App;
