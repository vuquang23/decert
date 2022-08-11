import NavBar from "components/NavBar";
import { Outlet } from "react-router-dom";

const PageLayout = () => (
  <>
    <NavBar />
    <div className="container mt-5">
      <Outlet />
    </div>
  </>
);

export default PageLayout;
